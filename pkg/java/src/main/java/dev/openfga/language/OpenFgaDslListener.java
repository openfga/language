package dev.openfga.language;

import dev.openfga.sdk.api.model.*;
import lombok.Getter;

import java.util.ArrayList;
import java.util.HashMap;

public class OpenFgaDslListener extends OpenFGAParserBaseListener {
    private static final String RELATION_DEFINITION_OPERATOR_OR = "or";
    private static final String RELATION_DEFINITION_OPERATOR_AND = "and";
    private static final String RELATION_DEFINITION_OPERATOR_BUT_NOT = "but not";

    @Getter
    private final AuthorizationModel authorizationModel = new AuthorizationModel();
    private final OpenFGAParser parser;
    private TypeDefinition currentTypeDef = null;
    private Relation currentRelation = null;

    public OpenFgaDslListener(OpenFGAParser parser) {
        this.parser = parser;
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
    }

    @Override
    public void exitRelationDeclaration(OpenFGAParser.RelationDeclarationContext ctx) {
        if (ctx.relationName() == null) {
            return;
        }

        var relationName = ctx.relationName().getText();
        var rewrites = currentRelation != null ? currentRelation.getRewrites() : null;
        if (rewrites == null || rewrites.isEmpty()) {
            return;
        }

        Userset relationDef = null;
        if (rewrites.size() == 1) {
            relationDef = rewrites.get(0);
        } else if (currentRelation.getOperator() != null) {
            switch (currentRelation.getOperator()) {
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
        }

        if (relationDef != null) {

            if (this.currentTypeDef.getRelations().get(relationName) != null) {
                var message = String.format("'%s' is already defined in '%s'", relationName, currentTypeDef.getType());
                parser.notifyErrorListeners(ctx.relationName().start, message, null);
            }

            currentTypeDef.getRelations().put(relationName, relationDef);

            var typeInfo = currentRelation != null ? currentRelation.getTypeInfo() : null;
            var directlyRelatedUserTypes = typeInfo != null ? typeInfo.getDirectlyRelatedUserTypes() : null;
            currentTypeDef.getMetadata().getRelations().put(relationName, new RelationMetadata().directlyRelatedUserTypes(directlyRelatedUserTypes));

        }

        this.currentRelation = null;
    }

    @Override
    public void enterRelationDefDirectAssignment(OpenFGAParser.RelationDefDirectAssignmentContext ctx) {
        currentRelation.setTypeInfo(new RelationMetadata().directlyRelatedUserTypes(new ArrayList<>()));
    }

    @Override
    public void exitRelationDefDirectAssignment(OpenFGAParser.RelationDefDirectAssignmentContext ctx) {
        var partialRewrite = new Userset()._this(new HashMap<>());
        if (currentRelation != null && currentRelation.getRewrites() != null) {
            currentRelation.getRewrites().add(partialRewrite);
        }
    }

    @Override
    public void exitRelationDefTypeRestriction(OpenFGAParser.RelationDefTypeRestrictionContext ctx) {
        var relationRef = new PartialRelationReference();

        var baseRestriction = ctx.relationDefTypeRestrictionBase();
        if (baseRestriction == null) {
            return;
        }

        relationRef.setType(baseRestriction.relationDefTypeRestrictionType != null
                ? baseRestriction.relationDefTypeRestrictionType.getText()
                : null);
        var usersetRestriction = baseRestriction.relationDefTypeRestrictionRelation;
        var wildcardRestriction = baseRestriction.relationDefTypeRestrictionWildcard;

        if (ctx.conditionName() != null) {
            relationRef.setCondition(ctx.conditionName().getText());
        }

        if (usersetRestriction != null) {
            relationRef.setRelation(usersetRestriction.getText());
        }

        if (wildcardRestriction != null) {
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

        if (currentRelation != null && currentRelation.getRewrites() != null) {
            currentRelation.getRewrites().add(partialRewrite);
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