package dev.openfga.language.utils;

import dev.openfga.sdk.api.model.AuthorizationModel;
import dev.openfga.sdk.api.model.Metadata;
import dev.openfga.sdk.api.model.RelationMetadata;
import dev.openfga.sdk.api.model.TypeDefinition;
import dev.openfga.sdk.api.model.Userset;
import java.util.Map;

public class ModelUtils {
    /**
     * getModuleForObjectTypeRelation returns the module for the given object type and relation in that type.
     *
     * @param typeDef  A TypeDefinition object which contains metadata about the type.
     * @param relation A string representing the relation whose module is to be retrieved.
     * @return A string representing the module for the given object type and relation.
     * @throws Exception An error if the relation does not exist.
     */
    public static String getModuleForObjectTypeRelation(TypeDefinition typeDef, String relation) throws Exception {
        Map<String, Userset> relations = typeDef.getRelations();
        if (relations == null || !relations.containsKey(relation)) {
            throw new Exception("relation " + relation + " does not exist in type " + typeDef.getType());
        }

        Metadata metadata = typeDef.getMetadata();
        if (metadata == null || metadata.getRelations() == null) {
            return null;
        }

        RelationMetadata relationMetadata = metadata.getRelations().get(relation);
        if (relationMetadata == null
                || relationMetadata.getModule() == null
                || relationMetadata.getModule().isEmpty()) {
            return metadata.getModule();
        }

        return relationMetadata.getModule();
    }

    /**
     * isRelationAssignable returns true if the relation is assignable, as in the relation definition has a key "this" or
     * any of its children have a key "this".
     *
     * @param relDef A Userset object representing a relation definition.
     * @return A boolean representing whether the relation definition has a key "this".
     */
    public static boolean isRelationAssignable(Userset relDef) {
        if (relDef == null) {
            return false;
        }

        if (relDef.getThis() != null) {
            return true;
        } else if (relDef.getUnion() != null) {
            for (Userset child : relDef.getUnion().getChild()) {
                if (isRelationAssignable(child)) {
                    return true;
                }
            }
        } else if (relDef.getIntersection() != null) {
            for (Userset child : relDef.getIntersection().getChild()) {
                if (isRelationAssignable(child)) {
                    return true;
                }
            }
        } else if (relDef.getDifference() != null) {
            var diff = relDef.getDifference();
            if (isRelationAssignable(diff.getBase()) || isRelationAssignable(diff.getSubtract())) {
                return true;
            }
        }

        // ComputedUserset and TupleToUserset are not assignable
        return false;
    }

    /**
     * isModelModular returns true if the model is modular: schema version 1.2 with at least one type or relation
     * that declares a module in its metadata.
     *
     * @param model An AuthorizationModel object.
     * @return Whether the model is modular.
     * @throws IllegalArgumentException If the model's schema version is not recognized.
     */
    public static boolean isModelModular(AuthorizationModel model) {
        var schemaVersion = model != null && model.getSchemaVersion() != null ? model.getSchemaVersion() : "1.1";
        switch (schemaVersion) {
            case "1.0":
            case "1.1":
                return false;
            case "1.2":
                break;
            default:
                throw new IllegalArgumentException("Unsupported schema version: " + schemaVersion);
        }

        if (model.getTypeDefinitions() == null) {
            return false;
        }

        for (var typeDef : model.getTypeDefinitions()) {
            var metadata = typeDef.getMetadata();
            if (metadata == null) {
                continue;
            }
            if (isNotBlank(metadata.getModule())) {
                return true;
            }
            if (metadata.getRelations() != null
                    && metadata.getRelations().values().stream()
                            .anyMatch(relation -> isNotBlank(relation.getModule()))) {
                return true;
            }
        }

        return false;
    }

    private static boolean isNotBlank(String value) {
        return value != null && !value.isEmpty();
    }
}
