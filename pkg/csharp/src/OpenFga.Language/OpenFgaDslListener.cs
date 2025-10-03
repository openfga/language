using System;
using System.Collections.Generic;
using System.Linq;
using Antlr4.Runtime;
using OpenFga.Sdk.Model;

namespace OpenFga.Language
{
    /// <summary>
    /// OpenFGA DSL Listener that uses OpenFGA SDK model classes
    /// This is the main version that works with the actual SDK
    /// </summary>
    public class OpenFgaDslListener : OpenFGAParserBaseListener
    {
        private const string RELATION_DEFINITION_OPERATOR_OR = "or";
        private const string RELATION_DEFINITION_OPERATOR_AND = "and";
        private const string RELATION_DEFINITION_OPERATOR_BUT_NOT = "but not";

        private readonly AuthorizationModel authorizationModel = new AuthorizationModel();
        private readonly OpenFGAParser parser;
        private TypeDefinition? currentTypeDef = null;
        private Relation? currentRelation = null;
        private Condition? currentCondition = null;
        private bool isModularModel = false;
        private Dictionary<string, TypeDefinition> typeDefExtensions = new Dictionary<string, TypeDefinition>();

        private Stack<StackRelation>? rewriteStack = null;

        public OpenFgaDslListener(OpenFGAParser parser)
        {
            this.parser = parser;
        }

        public AuthorizationModel GetAuthorizationModel()
        {
            return authorizationModel;
        }

        private Userset? ParseExpression(List<Userset> rewrites, string @operator)
        {
            if (rewrites == null || !rewrites.Any())
            {
                return null;
            }

            if (rewrites.Count == 1)
            {
                return rewrites[0];
            }

            Userset? relationDef = null;
            switch (@operator)
            {
                case RELATION_DEFINITION_OPERATOR_OR:
                    relationDef = new Userset
                    {
                        Union = new Usersets
                        {
                            Child = rewrites
                        }
                    };
                    break;
                case RELATION_DEFINITION_OPERATOR_AND:
                    relationDef = new Userset
                    {
                        Intersection = new Usersets
                        {
                            Child = rewrites
                        }
                    };
                    break;
                case RELATION_DEFINITION_OPERATOR_BUT_NOT:
                    relationDef = new Userset
                    {
                        Difference = new Difference
                        {
                            Base = rewrites[0],
                            Subtract = rewrites[1]
                        }
                    };
                    break;
            }

            return relationDef;
        }

        public override void EnterMain(OpenFGAParser.MainContext context)
        {
            authorizationModel.Conditions = new Dictionary<string, Condition>();
            base.EnterMain(context);
        }

        public override void ExitModelHeader(OpenFGAParser.ModelHeaderContext context)
        {
            if (context.SCHEMA_VERSION() != null)
            {
                authorizationModel.SchemaVersion = context.SCHEMA_VERSION().GetText();
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
            this.authorizationModel.TypeDefinitions = new List<TypeDefinition>();
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

            currentTypeDef = new TypeDefinition
            {
                Type = context.typeName.GetText(),
                Relations = new Dictionary<string, Userset>(),
                Metadata = new Metadata
                {
                    Relations = new Dictionary<string, RelationMetadata>()
                }
            };

            base.EnterTypeDef(context);
        }

        public override void EnterConditions(OpenFGAParser.ConditionsContext context)
        {
            authorizationModel.Conditions = new Dictionary<string, Condition>();
            base.EnterConditions(context);
        }

        public override void EnterCondition(OpenFGAParser.ConditionContext context)
        {
            if (context.conditionName() == null)
            {
                return;
            }

            var conditionName = context.conditionName().GetText();
            if (authorizationModel.Conditions.ContainsKey(conditionName))
            {
                var message = string.Format("condition '{0}' is already defined in the model", conditionName);
                parser.NotifyErrorListeners(context.conditionName().Start, message, null);
            }

            currentCondition = new Condition
            {
                Name = conditionName,
                Expression = "",
                Parameters = new Dictionary<string, ConditionParamTypeRef>()
            };

            base.EnterCondition(context);
        }

        public override void ExitConditionParameter(OpenFGAParser.ConditionParameterContext context)
        {
            if (context.parameterName() == null || context.parameterType() == null)
            {
                return;
            }

            var parameterName = context.parameterName().GetText();
            if (currentCondition.Parameters.ContainsKey(parameterName))
            {
                var message = string.Format(
                    "parameter '{0}' is already defined in the condition '{1}'",
                    parameterName, currentCondition.Name);
                parser.NotifyErrorListeners(context.parameterName().Start, message, null);
            }

            var paramContainer = context.parameterType().CONDITION_PARAM_CONTAINER();
            var conditionParamTypeRef = new PartialConditionParamTypeRef();
            var typeName = context.parameterType().GetText();
            if (paramContainer != null)
            {
                typeName = paramContainer.GetText();
                conditionParamTypeRef.TypeName = ParseTypeName(paramContainer.GetText());
                if (context.parameterType().CONDITION_PARAM_TYPE() != null)
                {
                    var genericTypeName =
                        ParseTypeName(context.parameterType().CONDITION_PARAM_TYPE().GetText());
                    if (genericTypeName != (TypeName)13)
                    {
                        conditionParamTypeRef.GenericTypes =
                            [
                                new ConditionParamTypeRef()
                                {
                                    TypeName = genericTypeName
                                }
                            ];
                    }
                }
            }

            conditionParamTypeRef.TypeName = ParseTypeName(typeName);

            currentCondition.Parameters.Add(parameterName, conditionParamTypeRef.AsConditionParamTypeRef());

            base.ExitConditionParameter(context);
        }

        private TypeName ParseTypeName(string typeName)
        {
            return Enum.Parse<TypeName>(typeName);
        }

        public override void ExitConditionExpression(OpenFGAParser.ConditionExpressionContext context)
        {
            if (currentCondition != null)
            {
                currentCondition.Expression = context.GetText().Trim();
            }

            base.ExitConditionExpression(context);
        }

        public override void ExitCondition(OpenFGAParser.ConditionContext context)
        {
            if (currentCondition != null)
            {
                authorizationModel.Conditions[currentCondition.Name] = currentCondition;
                currentCondition = null;
            }

            base.ExitCondition(context);
        }

        public override void ExitTypeDef(OpenFGAParser.TypeDefContext context)
        {
            if (currentTypeDef == null)
            {
                return;
            }

            if (currentTypeDef.Metadata != null
                && currentTypeDef.Metadata.Relations != null
                && !currentTypeDef.Metadata.Relations.Any())
            {
                currentTypeDef.Metadata = null;
            }

            var typeDefinitions = authorizationModel.TypeDefinitions;
            if (typeDefinitions != null)
            {
                typeDefinitions.Add(currentTypeDef);
            }

            if (context.EXTEND() != null && this.isModularModel)
            {
                if (typeDefExtensions.ContainsKey(currentTypeDef.Type))
                {
                    parser.NotifyErrorListeners(
                        context.typeName.Start,
                        string.Format("'{0}' is already extended in file.", currentTypeDef.Type),
                        null);
                }
                else
                {
                    typeDefExtensions[currentTypeDef.Type] = currentTypeDef;
                }
            }

            currentTypeDef = null;
            base.ExitTypeDef(context);
        }

        public override void EnterRelationDeclaration(OpenFGAParser.RelationDeclarationContext context)
        {
            currentRelation = new Relation(
                null,
                new List<Userset>(),
                null,
                new RelationMetadata
                {
                    DirectlyRelatedUserTypes = new List<RelationReference>()
                });
            rewriteStack = new Stack<StackRelation>();

            base.EnterRelationDeclaration(context);
        }

        public override void ExitRelationDeclaration(OpenFGAParser.RelationDeclarationContext context)
        {
            if (context.relationName() == null || currentRelation == null || currentTypeDef == null)
            {
                return;
            }

            var relationName = context.relationName().GetText();

            var relationDef = ParseExpression(currentRelation.Rewrites, currentRelation.Operator);
            if (relationDef != null)
            {
                if (this.currentTypeDef.Relations.ContainsKey(relationName))
                {
                    var message = string.Format("'{0}' is already defined in '{1}'", relationName, currentTypeDef.Type);
                    parser.NotifyErrorListeners(context.relationName().Start, message, null);
                }

                currentTypeDef.Relations[relationName] = relationDef;
                var directlyRelatedUserTypes = currentRelation.TypeInfo.DirectlyRelatedUserTypes;
                if (currentTypeDef.Metadata?.Relations != null)
                {
                    currentTypeDef.Metadata.Relations[relationName] = new RelationMetadata
                    {
                        DirectlyRelatedUserTypes = directlyRelatedUserTypes
                    };
                }
            }

            currentRelation = null;
            base.ExitRelationDeclaration(context);
        }

        public override void EnterRelationDefDirectAssignment(OpenFGAParser.RelationDefDirectAssignmentContext context)
        {
            if (currentRelation != null)
            {
                currentRelation.TypeInfo = new RelationMetadata
                {
                    DirectlyRelatedUserTypes = new List<RelationReference>()
                };
            }

            base.EnterRelationDefDirectAssignment(context);
        }

        public override void ExitRelationDefDirectAssignment(OpenFGAParser.RelationDefDirectAssignmentContext context)
        {
            if (currentRelation != null)
            {
                var partialRewrite = new Userset
                {
                    This = new Dictionary<string, object>()
                };
                currentRelation.Rewrites.Add(partialRewrite);
            }

            base.ExitRelationDefDirectAssignment(context);
        }

        public override void ExitRelationDefTypeRestriction(OpenFGAParser.RelationDefTypeRestrictionContext context)
        {
            if (currentRelation == null)
            {
                return;
            }

            var baseRestriction = context.relationDefTypeRestrictionBase();
            if (baseRestriction == null)
            {
                return;
            }

            var _type = baseRestriction.relationDefTypeRestrictionType;
            var usersetRestriction = baseRestriction.relationDefTypeRestrictionRelation;
            var wildcardRestriction = baseRestriction.relationDefTypeRestrictionWildcard;
            var conditionName = context.conditionName();

            var relationRef = new PartialRelationReference();
            if (_type != null)
            {
                relationRef.Type = _type.GetText();
            }

            if (conditionName != null)
            {
                relationRef.Condition = conditionName.GetText();
            }

            if (usersetRestriction != null)
            {
                relationRef.Relation = usersetRestriction.GetText();
            }

            if (wildcardRestriction != null)
            {
                relationRef.Wildcard = new Dictionary<string, object>();
            }

            currentRelation.TypeInfo.DirectlyRelatedUserTypes.Add(relationRef.AsRelationReference());
            base.ExitRelationDefTypeRestriction(context);
        }

        public override void ExitRelationDefRewrite(OpenFGAParser.RelationDefRewriteContext context)
        {
            if (currentRelation == null)
            {
                return;
            }

            var computedUserset = new ObjectRelation
            {
                Relation = context.rewriteComputedusersetName.GetText()
            };

            Userset partialRewrite;
            if (context.rewriteTuplesetName == null)
            {
                partialRewrite = new Userset
                {
                    ComputedUserset = computedUserset
                };
            }
            else
            {
                partialRewrite = new Userset
                {
                    TupleToUserset = new TupleToUserset
                    {
                        ComputedUserset = computedUserset,
                        Tupleset = new ObjectRelation
                        {
                            Relation = context.rewriteTuplesetName.GetText()
                        }
                    }
                };
            }

            currentRelation.Rewrites.Add(partialRewrite);
            base.ExitRelationDefRewrite(context);
        }

        public override void ExitRelationRecurse(OpenFGAParser.RelationRecurseContext context)
        {
            if (currentRelation == null)
            {
                return;
            }

            var relationDef = ParseExpression(currentRelation.Rewrites, currentRelation.Operator);

            if (relationDef != null)
            {
                currentRelation.Rewrites = new List<Userset> { relationDef };
            }

            base.ExitRelationRecurse(context);
        }

        public override void EnterRelationRecurseNoDirect(OpenFGAParser.RelationRecurseNoDirectContext context)
        {
            if (rewriteStack != null && currentRelation != null)
            {
                rewriteStack.Push(new StackRelation(currentRelation.Rewrites, currentRelation.Operator));
            }

            if (currentRelation != null)
            {
                currentRelation.Rewrites = new List<Userset>();
            }

            base.EnterRelationRecurseNoDirect(context);
        }

        public override void ExitRelationRecurseNoDirect(OpenFGAParser.RelationRecurseNoDirectContext context)
        {
            if (currentRelation == null || rewriteStack == null)
            {
                return;
            }

            var popped = rewriteStack.Pop();

            var relationDef = ParseExpression(currentRelation.Rewrites, currentRelation.Operator);
            if (relationDef != null)
            {
                currentRelation.Operator = popped.Operator;
                currentRelation.Rewrites = new List<Userset>(popped.Rewrites) { relationDef };
            }

            base.ExitRelationRecurseNoDirect(context);
        }

        public override void EnterRelationDefPartials(OpenFGAParser.RelationDefPartialsContext context)
        {
            if (currentRelation != null)
            {
                if (context.OR() != null && context.OR().Length > 0)
                {
                    currentRelation.Operator = RELATION_DEFINITION_OPERATOR_OR;
                }
                else if (context.AND() != null && context.AND().Length > 0)
                {
                    currentRelation.Operator = RELATION_DEFINITION_OPERATOR_AND;
                }
                else if (context.BUT_NOT() != null)
                {
                    currentRelation.Operator = RELATION_DEFINITION_OPERATOR_BUT_NOT;
                }
            }

            base.EnterRelationDefPartials(context);
        }
    }
}