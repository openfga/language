using OpenFga.Sdk.Model;

namespace OpenFga.Language;

/// <summary>
/// OpenFGA DSL Listener that uses OpenFGA SDK model classes
/// This is the main version that works with the actual SDK
/// </summary>
public class OpenFgaDslListener(OpenFGAParser parser)
    : OpenFGAParserBaseListener {

    private const string RelationDefinitionOperatorOr = "or";
    private const string RelationDefinitionOperatorAnd = "and";
    private const string RelationDefinitionOperatorButNot = "but not";

    private readonly AuthorizationModel _authorizationModel = new();
    private TypeDefinition? _currentTypeDef;
    private Relation? _currentRelation;
    private Condition? _currentCondition;
    private bool _isModularModel;
    private readonly Dictionary<string, TypeDefinition> _typeDefExtensions = new();

    private Stack<StackRelation>? _rewriteStack;

    public AuthorizationModel GetAuthorizationModel() {
        return _authorizationModel;
    }

    private Userset? ParseExpression(List<Userset> rewrites, string? @operator) {
        if (rewrites == null || !rewrites.Any()) {
            return null;
        }

        if (rewrites.Count == 1) {
            return rewrites[0];
        }

        Userset? relationDef = null;
        switch (@operator) {
            case RelationDefinitionOperatorOr:
                relationDef = new Userset {
                    Union = new Usersets {
                        Child = rewrites
                    }
                };
                break;
            case RelationDefinitionOperatorAnd:
                relationDef = new Userset {
                    Intersection = new Usersets {
                        Child = rewrites
                    }
                };
                break;
            case RelationDefinitionOperatorButNot:
                relationDef = new Userset {
                    Difference = new Difference {
                        Base = rewrites[0],
                        Subtract = rewrites[1]
                    }
                };
                break;
        }

        return relationDef;
    }

    public override void EnterMain(OpenFGAParser.MainContext context) {
        _authorizationModel.Conditions = new Dictionary<string, Condition>();
        base.EnterMain(context);
    }

    public override void ExitMain(OpenFGAParser.MainContext context) {
        // TO MAKE TEST PASS: If there are no conditions, set the Conditions to null
        if ((_authorizationModel.Conditions?.Count ?? 0) == 0) {
            _authorizationModel.Conditions = null;
        }
        base.ExitMain(context);
    }

    public override void ExitModelHeader(OpenFGAParser.ModelHeaderContext context) {
        if (context.SCHEMA_VERSION() != null) {
            _authorizationModel.SchemaVersion = context.SCHEMA_VERSION().GetText();
        }

        base.ExitModelHeader(context);
    }

    public override void ExitModuleHeader(OpenFGAParser.ModuleHeaderContext context) {
        this._isModularModel = true;
        base.ExitModuleHeader(context);
    }

    public override void EnterTypeDefs(OpenFGAParser.TypeDefsContext context) {
        this._authorizationModel.TypeDefinitions = new List<TypeDefinition>();
        base.EnterTypeDefs(context);
    }

    public override void ExitTypeDefs(OpenFGAParser.TypeDefsContext context) {
        // TO MAKE TEST PASS: If there are no type definitions, set the TypeDefinitions to null
        if (this._authorizationModel.TypeDefinitions.Count == 0) {
            this._authorizationModel.TypeDefinitions = null!;
        }
        base.ExitTypeDefs(context);
    }

    public override void EnterTypeDef(OpenFGAParser.TypeDefContext context) {
        if (context.typeName == null) {
            return;
        }

        if (context.EXTEND() != null && !this._isModularModel) {
            parser.NotifyErrorListeners(context.typeName.Start, "extend can only be used in a modular model", null);
        }

        _currentTypeDef = new TypeDefinition {
            Type = context.typeName.GetText(),
            Relations = new Dictionary<string, Userset>(),
            Metadata = new Metadata {
                Relations = new Dictionary<string, RelationMetadata>()
            }
        };

        base.EnterTypeDef(context);
    }

    public override void EnterConditions(OpenFGAParser.ConditionsContext context) {
        _authorizationModel.Conditions = new Dictionary<string, Condition>();
        base.EnterConditions(context);
    }

    public override void EnterCondition(OpenFGAParser.ConditionContext context) {
        if (context.conditionName() == null) {
            return;
        }

        var conditionName = context.conditionName().GetText();
        if (_authorizationModel.Conditions?.ContainsKey(conditionName) ?? false) {
            var message = $"condition '{conditionName}' is already defined in the model";
            parser.NotifyErrorListeners(context.conditionName().Start, message, null);
        }

        _currentCondition = new Condition {
            Name = conditionName,
            Expression = "",
            Parameters = new Dictionary<string, ConditionParamTypeRef>()
        };

        base.EnterCondition(context);
    }

    public override void ExitConditionParameter(OpenFGAParser.ConditionParameterContext context) {
        if (context.parameterName() == null || context.parameterType() == null) {
            return;
        }

        var parameterName = context.parameterName().GetText();
        if (_currentCondition!.Parameters!.ContainsKey(parameterName)) {
            var message = string.Format(
                "parameter '{0}' is already defined in the condition '{1}'",
                parameterName, _currentCondition.Name);
            parser.NotifyErrorListeners(context.parameterName().Start, message, null);
        }

        var paramContainer = context.parameterType().CONDITION_PARAM_CONTAINER();
        var conditionParamTypeRef = new PartialConditionParamTypeRef();
        var typeName = context.parameterType().GetText();
        if (paramContainer != null) {
            typeName = paramContainer.GetText();
            conditionParamTypeRef.TypeName = ParseTypeName(paramContainer.GetText());
            if (context.parameterType().CONDITION_PARAM_TYPE() != null) {
                var genericTypeName =
                    ParseTypeName(context.parameterType().CONDITION_PARAM_TYPE().GetText());
                if (genericTypeName != (TypeName)13) {
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

        _currentCondition.Parameters[parameterName] = conditionParamTypeRef.AsConditionParamTypeRef();

        base.ExitConditionParameter(context);
    }

    private TypeName ParseTypeName(string typeName) {
        return Enum.Parse<TypeName>("TYPENAME" + typeName.ToUpper(), true);
    }

    public override void ExitConditionExpression(OpenFGAParser.ConditionExpressionContext context) {
        if (_currentCondition != null) {
            _currentCondition.Expression = context.GetText().Trim();
        }

        base.ExitConditionExpression(context);
    }

    public override void ExitCondition(OpenFGAParser.ConditionContext context) {
        if (_currentCondition != null) {
            _authorizationModel.Conditions![_currentCondition.Name] = _currentCondition;
            _currentCondition = null;
        }

        base.ExitCondition(context);
    }

    public override void ExitTypeDef(OpenFGAParser.TypeDefContext context) {
        if (_currentTypeDef == null) {
            return;
        }

        if (_currentTypeDef.Metadata is { Relations: not null }
            && !_currentTypeDef.Metadata.Relations.Any()) {
            _currentTypeDef.Metadata = null;
        }

        var typeDefinitions = _authorizationModel.TypeDefinitions;
        if (typeDefinitions != null) {
            typeDefinitions.Add(_currentTypeDef);
        }

        if (context.EXTEND() != null && this._isModularModel) {
            if (_typeDefExtensions.ContainsKey(_currentTypeDef.Type)) {
                parser.NotifyErrorListeners(
                    context.typeName.Start,
                    string.Format("'{0}' is already extended in file.", _currentTypeDef.Type),
                    null);
            }
            else {
                _typeDefExtensions[_currentTypeDef.Type] = _currentTypeDef;
            }
        }

        _currentTypeDef = null;
        base.ExitTypeDef(context);
    }

    public override void EnterRelationDeclaration(OpenFGAParser.RelationDeclarationContext context) {
        _currentRelation = new Relation(
            null,
            new List<Userset>(),
            null,
            new RelationMetadata {
                DirectlyRelatedUserTypes = new List<RelationReference>()
            });
        _rewriteStack = new Stack<StackRelation>();

        base.EnterRelationDeclaration(context);
    }

    public override void ExitRelationDeclaration(OpenFGAParser.RelationDeclarationContext context) {
        if (context.relationName() == null || _currentRelation == null || _currentTypeDef == null) {
            return;
        }

        var relationName = context.relationName().GetText();

        var relationDef = ParseExpression(_currentRelation.Rewrites, _currentRelation.Operator);
        if (relationDef != null) {
            if (_currentTypeDef.Relations!.ContainsKey(relationName)) {
                var message = $"'{relationName}' is already defined in '{_currentTypeDef.Type}'";
                parser.NotifyErrorListeners(context.relationName().Start, message, null);
            }

            _currentTypeDef.Relations[relationName] = relationDef;
            var directlyRelatedUserTypes = _currentRelation.TypeInfo.DirectlyRelatedUserTypes;
            if (_currentTypeDef.Metadata?.Relations != null) {
                _currentTypeDef.Metadata.Relations[relationName] = new RelationMetadata {
                    DirectlyRelatedUserTypes = directlyRelatedUserTypes
                };
            }
        }

        _currentRelation = null;
        base.ExitRelationDeclaration(context);
    }

    public override void EnterRelationDefDirectAssignment(OpenFGAParser.RelationDefDirectAssignmentContext context) {
        if (_currentRelation != null) {
            _currentRelation.TypeInfo = new RelationMetadata {
                DirectlyRelatedUserTypes = new List<RelationReference>()
            };
        }

        base.EnterRelationDefDirectAssignment(context);
    }

    public override void ExitRelationDefDirectAssignment(OpenFGAParser.RelationDefDirectAssignmentContext context) {
        if (_currentRelation != null) {
            var partialRewrite = new Userset {
                This = new Dictionary<string, object>()
            };
            _currentRelation.Rewrites.Add(partialRewrite);
        }

        base.ExitRelationDefDirectAssignment(context);
    }

    public override void ExitRelationDefTypeRestriction(OpenFGAParser.RelationDefTypeRestrictionContext context) {
        if (_currentRelation == null) {
            return;
        }

        var baseRestriction = context.relationDefTypeRestrictionBase();
        if (baseRestriction == null) {
            return;
        }

        var t = baseRestriction.relationDefTypeRestrictionType;
        var usersetRestriction = baseRestriction.relationDefTypeRestrictionRelation;
        var wildcardRestriction = baseRestriction.relationDefTypeRestrictionWildcard;
        var conditionName = context.conditionName();

        var relationRef = new PartialRelationReference();
        if (t != null) {
            relationRef.Type = t.GetText();
        }

        if (conditionName != null) {
            relationRef.Condition = conditionName.GetText();
        }

        if (usersetRestriction != null) {
            relationRef.Relation = usersetRestriction.GetText();
        }

        if (wildcardRestriction != null) {
            relationRef.Wildcard = new Dictionary<string, object>();
        }

        _currentRelation.TypeInfo.DirectlyRelatedUserTypes!.Add(relationRef.AsRelationReference());
        base.ExitRelationDefTypeRestriction(context);
    }

    public override void ExitRelationDefRewrite(OpenFGAParser.RelationDefRewriteContext context) {
        if (_currentRelation == null) {
            return;
        }

        var computedUserset = new ObjectRelation {
            Relation = context.rewriteComputedusersetName.GetText()
        };

        Userset partialRewrite;
        if (context.rewriteTuplesetName == null) {
            partialRewrite = new Userset {
                ComputedUserset = computedUserset
            };
        }
        else {
            partialRewrite = new Userset {
                TupleToUserset = new TupleToUserset {
                    ComputedUserset = computedUserset,
                    Tupleset = new ObjectRelation {
                        Relation = context.rewriteTuplesetName.GetText()
                    }
                }
            };
        }

        _currentRelation.Rewrites.Add(partialRewrite);
        base.ExitRelationDefRewrite(context);
    }

    public override void ExitRelationRecurse(OpenFGAParser.RelationRecurseContext context) {
        if (_currentRelation == null) {
            return;
        }

        var relationDef = ParseExpression(_currentRelation.Rewrites, _currentRelation.Operator);

        if (relationDef != null) {
            _currentRelation.Rewrites = new List<Userset> { relationDef };
        }

        base.ExitRelationRecurse(context);
    }

    public override void EnterRelationRecurseNoDirect(OpenFGAParser.RelationRecurseNoDirectContext context) {
        if (_rewriteStack != null && _currentRelation != null) {
            _rewriteStack.Push(new StackRelation(_currentRelation.Rewrites, _currentRelation.Operator));
        }

        if (_currentRelation != null) {
            _currentRelation.Rewrites = new List<Userset>();
        }

        base.EnterRelationRecurseNoDirect(context);
    }

    public override void ExitRelationRecurseNoDirect(OpenFGAParser.RelationRecurseNoDirectContext context) {
        if (_currentRelation == null || _rewriteStack == null) {
            return;
        }

        var popped = _rewriteStack.Pop();

        var relationDef = ParseExpression(_currentRelation.Rewrites, _currentRelation.Operator);
        if (relationDef != null) {
            _currentRelation.Operator = popped.Operator;
            _currentRelation.Rewrites = new List<Userset>(popped.Rewrites) { relationDef };
        }

        base.ExitRelationRecurseNoDirect(context);
    }

    public override void EnterRelationDefPartials(OpenFGAParser.RelationDefPartialsContext context) {
        if (_currentRelation != null) {
            if (context.OR() != null && context.OR().Length > 0) {
                _currentRelation.Operator = RelationDefinitionOperatorOr;
            }
            else if (context.AND() != null && context.AND().Length > 0) {
                _currentRelation.Operator = RelationDefinitionOperatorAnd;
            }
            else if (context.BUT_NOT() != null) {
                _currentRelation.Operator = RelationDefinitionOperatorButNot;
            }
        }

        base.EnterRelationDefPartials(context);
    }
}