using System;
using System.Collections.Generic;
using System.Linq;
using Antlr4.Runtime;

namespace OpenFga.Language
{
    /// <summary>
    /// Simplified OpenFGA DSL Listener that focuses on core ANTLR functionality
    /// without complex OpenFGA model dependencies
    /// </summary>
    public class OpenFgaDslListenerSimple : OpenFGAParserBaseListener
    {
        private const string RELATION_DEFINITION_OPERATOR_OR = "or";
        private const string RELATION_DEFINITION_OPERATOR_AND = "and";
        private const string RELATION_DEFINITION_OPERATOR_BUT_NOT = "but not";

        private readonly OpenFGAParser parser;
        private bool isModularModel = false;
        private Dictionary<string, object> typeDefExtensions = new Dictionary<string, object>();
        private Stack<StackRelationSimple> rewriteStack = null;

        // Simple data structures to hold parsed information
        public Dictionary<string, object> ParsedModel { get; private set; } = new Dictionary<string, object>();
        public List<object> TypeDefinitions { get; private set; } = new List<object>();
        public Dictionary<string, object> Conditions { get; private set; } = new Dictionary<string, object>();

        public OpenFgaDslListenerSimple(OpenFGAParser parser)
        {
            this.parser = parser;
        }

        public override void EnterMain(OpenFGAParser.MainContext context)
        {
            ParsedModel["conditions"] = new Dictionary<string, object>();
            base.EnterMain(context);
        }

        public override void ExitModelHeader(OpenFGAParser.ModelHeaderContext context)
        {
            if (context.SCHEMA_VERSION() != null)
            {
                ParsedModel["schema_version"] = context.SCHEMA_VERSION().GetText();
            }
            base.ExitModelHeader(context);
        }

        public override void ExitModuleHeader(OpenFGAParser.ModuleHeaderContext context)
        {
            this.isModularModel = true;
            base.ExitModuleHeader(context);
        }

        public override void EnterTypeDefs(OpenFGAParser.TypeDefsContext context)
        {
            this.TypeDefinitions = new List<object>();
            base.EnterTypeDefs(context);
        }

        public override void EnterTypeDef(OpenFGAParser.TypeDefContext context)
        {
            if (context.typeName == null)
            {
                return;
            }

            if (context.EXTEND() != null && !this.isModularModel)
            {
                parser.NotifyErrorListeners(context.typeName.Start, "extend can only be used in a modular model", null);
            }

            var typeDef = new Dictionary<string, object>
            {
                ["type"] = context.typeName.GetText(),
                ["relations"] = new Dictionary<string, object>(),
                ["metadata"] = new Dictionary<string, object>
                {
                    ["relations"] = new Dictionary<string, object>()
                }
            };

            // Store current type definition for processing
            ParsedModel["current_type_def"] = typeDef;

            base.EnterTypeDef(context);
        }

        public override void EnterConditions(OpenFGAParser.ConditionsContext context)
        {
            Conditions = new Dictionary<string, object>();
            base.EnterConditions(context);
        }

        public override void EnterCondition(OpenFGAParser.ConditionContext context)
        {
            if (context.conditionName() == null)
            {
                return;
            }

            var conditionName = context.conditionName().GetText();
            if (Conditions.ContainsKey(conditionName))
            {
                var message = string.Format("condition '{0}' is already defined in the model", conditionName);
                parser.NotifyErrorListeners(context.conditionName().Start, message, null);
            }

            var condition = new Dictionary<string, object>
            {
                ["name"] = conditionName,
                ["expression"] = "",
                ["parameters"] = new Dictionary<string, object>()
            };

            ParsedModel["current_condition"] = condition;

            base.EnterCondition(context);
        }

        public override void ExitConditionExpression(OpenFGAParser.ConditionExpressionContext context)
        {
            if (ParsedModel.ContainsKey("current_condition"))
            {
                var condition = (Dictionary<string, object>)ParsedModel["current_condition"];
                condition["expression"] = context.GetText().Trim();
            }
            base.ExitConditionExpression(context);
        }

        public override void ExitCondition(OpenFGAParser.ConditionContext context)
        {
            if (ParsedModel.ContainsKey("current_condition"))
            {
                var condition = (Dictionary<string, object>)ParsedModel["current_condition"];
                Conditions[(string)condition["name"]] = condition;
                ParsedModel.Remove("current_condition");
            }
            base.ExitCondition(context);
        }

        public override void ExitTypeDef(OpenFGAParser.TypeDefContext context)
        {
            if (!ParsedModel.ContainsKey("current_type_def"))
            {
                return;
            }

            var currentTypeDef = (Dictionary<string, object>)ParsedModel["current_type_def"];

            if (currentTypeDef.ContainsKey("metadata"))
            {
                var metadata = (Dictionary<string, object>)currentTypeDef["metadata"];
                if (metadata.ContainsKey("relations"))
                {
                    var relations = (Dictionary<string, object>)metadata["relations"];
                    if (!relations.Any())
                    {
                        currentTypeDef.Remove("metadata");
                    }
                }
            }

            TypeDefinitions.Add(currentTypeDef);

            if (context.EXTEND() != null && this.isModularModel)
            {
                var typeName = (string)currentTypeDef["type"];
                if (typeDefExtensions.ContainsKey(typeName))
                {
                    parser.NotifyErrorListeners(
                            context.typeName.Start,
                            string.Format("'{0}' is already extended in file.", typeName),
                            null);
                }
                else
                {
                    typeDefExtensions[typeName] = currentTypeDef;
                }
            }

            ParsedModel.Remove("current_type_def");
            base.ExitTypeDef(context);
        }

        public override void EnterRelationDeclaration(OpenFGAParser.RelationDeclarationContext context)
        {
            var relation = new Dictionary<string, object>
            {
                ["rewrites"] = new List<object>(),
                ["operator"] = null,
                ["type_info"] = new Dictionary<string, object>
                {
                    ["directly_related_user_types"] = new List<object>()
                }
            };

            ParsedModel["current_relation"] = relation;
            rewriteStack = new Stack<StackRelationSimple>();

            base.EnterRelationDeclaration(context);
        }

        public override void ExitRelationDeclaration(OpenFGAParser.RelationDeclarationContext context)
        {
            if (context.relationName() == null || !ParsedModel.ContainsKey("current_relation"))
            {
                return;
            }

            var relationName = context.relationName().GetText();
            var currentRelation = (Dictionary<string, object>)ParsedModel["current_relation"];
            var currentTypeDef = (Dictionary<string, object>)ParsedModel["current_type_def"];

            if (currentTypeDef != null && currentTypeDef.ContainsKey("relations"))
            {
                var relations = (Dictionary<string, object>)currentTypeDef["relations"];
                if (relations.ContainsKey(relationName))
                {
                    var message = string.Format("'{0}' is already defined in '{1}'", relationName, currentTypeDef["type"]);
                    parser.NotifyErrorListeners(context.relationName().Start, message, null);
                }

                // Store the relation definition
                relations[relationName] = currentRelation;
            }

            ParsedModel.Remove("current_relation");
            base.ExitRelationDeclaration(context);
        }

        public override void EnterRelationDefDirectAssignment(OpenFGAParser.RelationDefDirectAssignmentContext context)
        {
            if (ParsedModel.ContainsKey("current_relation"))
            {
                var currentRelation = (Dictionary<string, object>)ParsedModel["current_relation"];
                currentRelation["type_info"] = new Dictionary<string, object>
                {
                    ["directly_related_user_types"] = new List<object>()
                };
            }
            base.EnterRelationDefDirectAssignment(context);
        }

        public override void ExitRelationDefDirectAssignment(OpenFGAParser.RelationDefDirectAssignmentContext context)
        {
            if (ParsedModel.ContainsKey("current_relation"))
            {
                var currentRelation = (Dictionary<string, object>)ParsedModel["current_relation"];
                var rewrites = (List<object>)currentRelation["rewrites"];
                rewrites.Add(new Dictionary<string, object> { ["type"] = "this" });
            }
            base.ExitRelationDefDirectAssignment(context);
        }

        public override void EnterRelationDefPartials(OpenFGAParser.RelationDefPartialsContext context)
        {
            if (ParsedModel.ContainsKey("current_relation"))
            {
                var currentRelation = (Dictionary<string, object>)ParsedModel["current_relation"];
                
                if (context.OR() != null && context.OR().Length > 0)
                {
                    currentRelation["operator"] = RELATION_DEFINITION_OPERATOR_OR;
                }
                else if (context.AND() != null && context.AND().Length > 0)
                {
                    currentRelation["operator"] = RELATION_DEFINITION_OPERATOR_AND;
                }
                else if (context.BUT_NOT() != null)
                {
                    currentRelation["operator"] = RELATION_DEFINITION_OPERATOR_BUT_NOT;
                }
            }
            base.EnterRelationDefPartials(context);
        }
    }

    /// <summary>
    /// Simple helper class for relation stack operations
    /// </summary>
    public class StackRelationSimple
    {
        public List<object> Rewrites { get; set; }
        public string Operator { get; set; }

        public StackRelationSimple(List<object> rewrites, string @operator)
        {
            Rewrites = rewrites;
            Operator = @operator;
        }
    }
}
