package dev.openfga.language;

import dev.openfga.language.antlr.OpenFGAParser;
import dev.openfga.language.antlr.OpenFGAParserBaseListener;
import dev.openfga.sdk.api.model.*;

import java.util.*;

public class OpenFgaDslListener extends OpenFGAParserBaseListener {
    private static final String RELATION_DEFINITION_OPERATOR_OR = "or";
    private static final String RELATION_DEFINITION_OPERATOR_AND = "and";
    private static final String RELATION_DEFINITION_OPERATOR_BUT_NOT = "but not";

    private final AuthorizationModel authorizationModel = new AuthorizationModel();
    private final OpenFGAParser parser;
    private TypeDefinition currentTypeDef = null;
    private Relation currentRelation = null;
    private Condition currentCondition = null;

    private Deque<StackRelation> rewriteStack = null;

    public OpenFgaDslListener(OpenFGAParser parser) {
        this.parser = parser;
    }

    public AuthorizationModel getAuthorizationModel() {
        return authorizationModel;
    }

    private Userset parseExpression(List<Userset> rewrites, String operator) {

        if (rewrites.isEmpty()) {
            return null;
        }
        if (rewrites.size() == 1) {
            return rewrites.get(0);
        }
        Userset relationDef = null;
        switch (operator) {
            case RELATION_DEFINITION_OPERATOR_OR:
                relationDef = new Userset().union(new Usersets().child(rewrites));
                break;
            case RELATION_DEFINITION_OPERATOR_AND:
                relationDef = new Userset().intersection(new Usersets().child(rewrites));
                break;
            case RELATION_DEFINITION_OPERATOR_BUT_NOT:
                relationDef = new Userset().difference(new Difference().base(rewrites.get(0)).subtract(rewrites.get(1)));
                break;
        }
        return relationDef;
    }

    @Override
    public void enterMain(OpenFGAParser.MainContext ctx) {
        authorizationModel.setConditions(new HashMap<>());
    }

    @Override
    public void exitModelHeader(OpenFGAParser.ModelHeaderContext ctx) {
        if (ctx.SCHEMA_VERSION() != null) {
            authorizationModel.setSchemaVersion(ctx.SCHEMA_VERSION().getText());
        }
    }

    @Override
    public void enterTypeDefs(OpenFGAParser.TypeDefsContext ctx) {
        this.authorizationModel.setTypeDefinitions(new ArrayList<>());
    }

    @Override
    public void enterTypeDef(OpenFGAParser.TypeDefContext ctx) {
        if (ctx.typeName == null) {
            return;
        }

        currentTypeDef = new TypeDefinition()
                .type(ctx.typeName.getText())
                .relations(new HashMap<>())
                .metadata(new Metadata().relations(new HashMap<>()));

    }

    @Override
    public void enterConditions(OpenFGAParser.ConditionsContext ctx) {
        authorizationModel.setConditions(new HashMap<>());
    }

    @Override
    public void enterCondition(OpenFGAParser.ConditionContext ctx) {
        if (ctx.conditionName() == null) {
            return;
        }

        var conditionName = ctx.conditionName().getText();
        if (authorizationModel.getConditions().containsKey(conditionName)) {
            var message = String.format("condition '%s' is already defined in the model", conditionName);
            parser.notifyErrorListeners(ctx.conditionName().start, message, null);
        }

        currentCondition = new Condition()
                .name(conditionName)
                .expression("")
                .parameters(new HashMap<>());
    }

    @Override
    public void exitConditionParameter(OpenFGAParser.ConditionParameterContext ctx) {
        if (ctx.parameterName() == null || ctx.parameterType() == null) {
            return;
        }

        var parameterName = ctx.parameterName().getText();
        if (currentCondition.getParameters().containsKey(parameterName)) {
            var message = String.format("parameter '%s' is already defined in the condition '%s'", parameterName, currentCondition.getName());
            parser.notifyErrorListeners(ctx.parameterName().start, message, null);
        }

        var paramContainer = ctx.parameterType().CONDITION_PARAM_CONTAINER();
        var conditionParamTypeRef = new PartialConditionParamTypeRef();
        var typeName = ctx.parameterType().getText();
        if (paramContainer != null) {
            typeName = paramContainer.getText();
            conditionParamTypeRef.setTypeName(parseTypeName(paramContainer.getText()));
            if (ctx.parameterType().CONDITION_PARAM_TYPE() != null) {
                var genericTypeName = parseTypeName(ctx.parameterType().CONDITION_PARAM_TYPE().getText());
                if (genericTypeName != TypeName.UNKNOWN_DEFAULT_OPEN_API) {
                    conditionParamTypeRef.setGenericTypes(
                            new ArrayList<>() {{
                                add(new ConditionParamTypeRef().typeName(genericTypeName));
                            }}
                    );
                }

            }
        }
        conditionParamTypeRef.setTypeName(parseTypeName(typeName));

        currentCondition.getParameters().put(parameterName, conditionParamTypeRef.asConditionParamTypeRef());
    }

    private TypeName parseTypeName(String typeName) {
        return TypeName.fromValue("TYPE_NAME_" + typeName.toUpperCase());
    }
    @Override
    public void exitConditionExpression(OpenFGAParser.ConditionExpressionContext ctx) {
        currentCondition.setExpression(ctx.getText().trim());
    }

    @Override
    public void exitCondition(OpenFGAParser.ConditionContext ctx) {
        if (currentCondition != null) {
            authorizationModel.getConditions().put(currentCondition.getName(), currentCondition);
            currentCondition = null;
        }
    }

    @Override
    public void exitTypeDef(OpenFGAParser.TypeDefContext ctx) {
        if (currentTypeDef == null) {
            return;
        }

        if (currentTypeDef.getMetadata() != null
                && currentTypeDef.getMetadata().getRelations() != null
                && currentTypeDef.getMetadata().getRelations().isEmpty()) {
            currentTypeDef.setMetadata(null);
        }

        var typeDefinitions = authorizationModel.getTypeDefinitions();
        if (typeDefinitions != null) {
            typeDefinitions.add(currentTypeDef);
        }

        currentTypeDef = null;
    }

    @Override
    public void enterRelationDeclaration(OpenFGAParser.RelationDeclarationContext ctx) {
        currentRelation = new Relation(
                null,
                new ArrayList<>(),
                null,
                new RelationMetadata().directlyRelatedUserTypes(new ArrayList<>()));
        rewriteStack = new ArrayDeque<>();
    }

    @Override
    public void exitRelationDeclaration(OpenFGAParser.RelationDeclarationContext ctx) {
        if (ctx.relationName() == null) {
            return;
        }

        var relationName = ctx.relationName().getText();

        var relationDef = parseExpression(currentRelation.getRewrites(), currentRelation.getOperator());
        if(relationDef != null) {
            if (this.currentTypeDef.getRelations().get(relationName) != null) {
                var message = String.format("'%s' is already defined in '%s'", relationName, currentTypeDef.getType());
                parser.notifyErrorListeners(ctx.relationName().start, message, null);
            }

            currentTypeDef.getRelations().put(relationName, relationDef);
            var directlyRelatedUserTypes = currentRelation.getTypeInfo().getDirectlyRelatedUserTypes();
            currentTypeDef.getMetadata().getRelations().put(relationName, new RelationMetadata().directlyRelatedUserTypes(directlyRelatedUserTypes));
        }

        currentRelation = null;
    }

    @Override
    public void enterRelationDefDirectAssignment(OpenFGAParser.RelationDefDirectAssignmentContext ctx) {
        currentRelation.setTypeInfo(new RelationMetadata().directlyRelatedUserTypes(new ArrayList<>()));
    }

    @Override
    public void exitRelationDefDirectAssignment(OpenFGAParser.RelationDefDirectAssignmentContext ctx) {
        var partialRewrite = new Userset()._this(new HashMap<>());
        currentRelation.getRewrites().add(partialRewrite);
    }

    @Override
    public void exitRelationDefTypeRestriction(OpenFGAParser.RelationDefTypeRestrictionContext ctx) {

        var baseRestriction = ctx.relationDefTypeRestrictionBase();
        if (baseRestriction == null) {
            return;
        }

        var _type = baseRestriction.relationDefTypeRestrictionType;
        var usersetRestriction = baseRestriction.relationDefTypeRestrictionRelation;
        var wildcardRestriction = baseRestriction.relationDefTypeRestrictionWildcard;
        var conditionName = ctx.conditionName();

        var relationRef = new PartialRelationReference();
        if (_type != null) {
            relationRef.setType(_type.getText());
        }

        if(conditionName != null) {
            relationRef.setCondition(conditionName.getText());
        }

        if(usersetRestriction != null) {
            relationRef.setRelation(usersetRestriction.getText());
        }

        if(wildcardRestriction != null) {
            relationRef.setWildcard(new HashMap<>());
        }

        currentRelation.getTypeInfo().getDirectlyRelatedUserTypes().add(relationRef.asRelationReference());
    }

    @Override
    public void exitRelationDefRewrite(OpenFGAParser.RelationDefRewriteContext ctx) {
        var computedUserset = new ObjectRelation().relation(ctx.rewriteComputedusersetName.getText());

        var partialRewrite = ctx.rewriteTuplesetName == null
                ? new Userset().computedUserset(computedUserset)
                : new Userset().tupleToUserset(new TupleToUserset()
                .computedUserset(computedUserset)
                .tupleset(new ObjectRelation().relation(ctx.rewriteTuplesetName.getText()))
        );

        currentRelation.getRewrites().add(partialRewrite);
    }

    @Override
    public void exitRelationRecurse(OpenFGAParser.RelationRecurseContext ctx) {
        if (currentRelation == null) {
            return;
        }

        var relationDef = parseExpression(currentRelation.getRewrites(), currentRelation.getOperator());

        if (relationDef != null) {
            currentRelation.setRewrites(new ArrayList<>() {{
                add(relationDef);
            }});
        }
    }

    @Override
    public void enterRelationRecurseNoDirect(OpenFGAParser.RelationRecurseNoDirectContext ctx) {
        if (rewriteStack != null) {
            rewriteStack.add(new StackRelation(currentRelation.getRewrites(), currentRelation.getOperator()));
        }

        currentRelation.setRewrites(new ArrayList<>());
    }

    @Override
    public void exitRelationRecurseNoDirect(OpenFGAParser.RelationRecurseNoDirectContext ctx) {
        if (currentRelation == null) {
            return;
        }

        var popped = rewriteStack.removeLast();

        var relationDef = parseExpression(currentRelation.getRewrites(), currentRelation.getOperator());
        if (relationDef != null) {
            currentRelation.setOperator(popped.getOperator());
            currentRelation.setRewrites(new ArrayList<>(popped.getRewrites()) {{
                add(relationDef);
            }});
        }
    }

    @Override
    public void enterRelationDefPartials(OpenFGAParser.RelationDefPartialsContext ctx) {
        if (!ctx.OR().isEmpty()) {
            currentRelation.setOperator(RELATION_DEFINITION_OPERATOR_OR);
        } else if (!ctx.AND().isEmpty()) {
            currentRelation.setOperator(RELATION_DEFINITION_OPERATOR_AND);
        } else if (ctx.BUT_NOT() != null) {
            currentRelation.setOperator(RELATION_DEFINITION_OPERATOR_BUT_NOT);
        }
    }
}