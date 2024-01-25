package dev.openfga.language;

import com.fasterxml.jackson.core.JsonProcessingException;
import dev.openfga.language.errors.UnsupportedDSLNestingException;
import dev.openfga.sdk.api.model.*;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.function.Function;

import static java.util.Objects.requireNonNullElseGet;
import static java.util.stream.Collectors.joining;

public class JsonToDslTransformer {

    private static final String EOL = System.lineSeparator();

    public String transform(String json) throws JsonProcessingException {

        var model = JSON.parse(json, AuthorizationModel.class);

        return transformJSONToDSL(model);

    }

    private String transformJSONToDSL(AuthorizationModel model) {
        var schemaVersion = "1.1";
        if (model != null && model.getSchemaVersion() != null) {
            schemaVersion = model.getSchemaVersion();
        }

        var formattedTypeDefinitions = new StringBuilder();
        if (model != null && model.getTypeDefinitions() != null) {
            for (var typeDefinition : model.getTypeDefinitions()) {
                formattedTypeDefinitions.append(formatType(typeDefinition)).append(EOL);
            }
        }

        var fomattedConditions = formatConditions(model);

        return "model" + EOL + "  schema " + schemaVersion + EOL +
                formattedTypeDefinitions +
                fomattedConditions;
    }

    private String formatType(TypeDefinition typeDef) {
        var typeName = typeDef.getType();
        var formatedTypeBuilder = new StringBuilder(EOL).append("type ").append(typeName);

        var relations = requireNonNullElseGet(typeDef.getRelations(), HashMap<String, Userset>::new);
        var metadata = typeDef.getMetadata();
        var emptyMetadataRelation = new HashMap<String, RelationMetadata>();
        var metadataRelations = metadata != null ? metadata.getRelations() : emptyMetadataRelation;
        if (metadataRelations == null) {
            metadataRelations = emptyMetadataRelation;
        }

        if (!relations.isEmpty()) {
            formatedTypeBuilder
                    .append(EOL)
                    .append("  relations");
            for (var relationEntry :
                    relations.entrySet()) {
                var relationName = relationEntry.getKey();
                var relationDefinition = relationEntry.getValue();
                metadataRelations.get(relationName);
                var formattedRelationString = formatRelation(typeName, relationName, relationDefinition, metadataRelations.get(relationName));
                formatedTypeBuilder.append(EOL).append(formattedRelationString);
            }
        }

        return formatedTypeBuilder.toString();
    }

    private String formatRelation(String typeName, String relationName, Userset relationDefinition, RelationMetadata relationMetadata) {
        var formattedRelationBuilder = new StringBuilder("    define ").append(relationName).append(": ");
        var typeRestrictions = requireNonNullElseGet(relationMetadata.getDirectlyRelatedUserTypes(), ArrayList<RelationReference>::new);
        if (relationDefinition.getDifference() != null) {
            formattedRelationBuilder.append(formatDifference(typeName, relationName, relationDefinition, typeRestrictions));
        } else if (relationDefinition.getUnion() != null) {
            formattedRelationBuilder.append(formatUnion(typeName, relationName, relationDefinition, typeRestrictions));
        } else if (relationDefinition.getIntersection() != null) {
            formattedRelationBuilder.append(formatIntersection(typeName, relationName, relationDefinition, typeRestrictions));
        } else {
            formattedRelationBuilder.append(formatSubRelation(typeName, relationName, relationDefinition, typeRestrictions));
        }

        return formattedRelationBuilder.toString();
    }

    private CharSequence formatDifference(String typeName, String relationName, Userset relationDefinition, List<RelationReference> typeRestrictions) {
        var base = formatSubRelation(typeName, relationName, relationDefinition.getDifference().getBase(), typeRestrictions);
        var difference = formatSubRelation(typeName, relationName, relationDefinition.getDifference().getSubtract(), typeRestrictions);

        return new StringBuilder(base).append(" but not ").append(difference);
    }

    private CharSequence formatUnion(String typeName, String relationName, Userset relationDefinition, List<RelationReference> typeRestrictions) {
        return joinChildren(Userset::getUnion, "or", typeName, relationName, relationDefinition, typeRestrictions);
    }

    private CharSequence formatIntersection(String typeName, String relationName, Userset relationDefinition, List<RelationReference> typeRestrictions) {
        return joinChildren(Userset::getIntersection, "and", typeName, relationName, relationDefinition, typeRestrictions);
    }

    private CharSequence joinChildren(Function<Userset, Usersets> childrenAccessor, String operator, String typeName, String relationName, Userset relationDefinition, List<RelationReference> typeRestrictions) {
        List<Userset> children = null;
        if (relationDefinition != null && childrenAccessor.apply(relationDefinition) != null) {
            children = childrenAccessor.apply(relationDefinition).getChild();
        }
        children = requireNonNullElseGet(children, ArrayList::new);

        var formattedUnion = new StringBuilder();
        boolean notFirst = false;
        for (var child : children) {
            if (notFirst) {
                formattedUnion.append(" ").append(operator).append(" ");
            } else {
                notFirst = true;
            }
            formattedUnion.append(formatSubRelation(typeName, relationName, child, typeRestrictions));
        }

        return formattedUnion;
    }

    private CharSequence formatSubRelation(String typeName, String relationName, Userset relationDefinition, List<RelationReference> typeRestrictions) {
        if (relationDefinition.getThis() != null) {
            return formatThis(typeRestrictions);
        }

        if (relationDefinition.getComputedUserset() != null) {
            return formatComputedUserset(relationDefinition);
        }

        if (relationDefinition.getTupleToUserset() != null) {
            return formatTupleToUserset(relationDefinition);
        }

        throw new UnsupportedDSLNestingException(typeName, relationName);
    }

    private CharSequence formatThis(List<RelationReference> typeRestrictions) {
        return requireNonNullElseGet(typeRestrictions, ArrayList<RelationReference>::new)
                .stream()
                .map(this::formatTypeRestriction)
                .collect(joining(", ", "[", "]"));
    }

    private CharSequence formatTypeRestriction(RelationReference restriction) {
        var typeName = restriction.getType();

        Object condition = null; // not supported yet
        if (condition != null) {
            return new StringBuilder(typeName).append(" with ").append(condition);
        }

        var relation = restriction.getRelation();
        if (relation != null) {
            return new StringBuilder(typeName).append("#").append(relation);
        }

        var wildcard = restriction.getWildcard();
        if (wildcard != null) {
            return new StringBuilder(typeName).append(":*");
        }

        return typeName;
    }

    private CharSequence formatComputedUserset(Userset relationDefinition) {
        return relationDefinition.getComputedUserset().getRelation();
    }

    private CharSequence formatTupleToUserset(Userset relationDefinition) {
        String computedUserset = "";
        String tupleset = "";
        if (relationDefinition != null && relationDefinition.getTupleToUserset() != null) {
            if (relationDefinition.getTupleToUserset().getComputedUserset() != null) {
                computedUserset = relationDefinition.getTupleToUserset().getComputedUserset().getRelation();
            }
            if (relationDefinition.getTupleToUserset().getTupleset() != null) {
                tupleset = relationDefinition.getTupleToUserset().getTupleset().getRelation();
            }
        }
        return new StringBuilder(computedUserset).append(" from ").append(tupleset);
    }

    private String formatConditions(AuthorizationModel model) {
        return "";
    }

}
