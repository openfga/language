package dev.openfga.language.validation;

import static dev.openfga.language.Utils.getNullSafe;
import static dev.openfga.language.Utils.getNullSafeList;
import static dev.openfga.language.validation.Dsl.getTypeRestrictions;

import dev.openfga.language.Utils;
import dev.openfga.sdk.api.model.Metadata;
import dev.openfga.sdk.api.model.RelationMetadata;
import dev.openfga.sdk.api.model.TypeDefinition;
import dev.openfga.sdk.api.model.Userset;
import java.util.HashMap;
import java.util.Map;

class EntryPointOrLoop {
    public static final EntryPointOrLoop BOTH_FALSE = new EntryPointOrLoop(false, false);
    public static final EntryPointOrLoop HAS_ENTRY_BUT_NO_LOOP = new EntryPointOrLoop(true, false);
    public static final EntryPointOrLoop NO_ENTRY_WITH_LOOP = new EntryPointOrLoop(false, true);

    private final boolean entry;
    private final boolean loop;

    public EntryPointOrLoop(boolean hasEntry, boolean isLoop) {
        this.entry = hasEntry;
        this.loop = isLoop;
    }

    // for the type/relation, whether there are any unique entry points, and if a loop is found
    // if there are unique entry points (i.e., direct relations) then it will return true
    // otherwise, it will follow its children to see if there are unique entry points
    // if there is a loop during traversal, the function will return a boolean indicating so
    public static EntryPointOrLoop compute(
            Map<String, TypeDefinition> typeMap,
            String typeName,
            String relationName,
            Userset rewrite,
            Map<String, Map<String, Boolean>> visitedRecords) {
        var visited = Utils.deepCopy(visitedRecords);

        if (relationName == null) {
            return BOTH_FALSE;
        }

        if (!visited.containsKey(typeName)) {
            visited.put(typeName, new HashMap<>());
        }
        visited.get(typeName).put(relationName, true);

        var currentRelations = typeMap.get(typeName).getRelations();
        if (currentRelations == null || !currentRelations.containsKey(relationName)) {
            return BOTH_FALSE;
        }

        if (typeMap.get(typeName).getRelations() == null
                || !typeMap.get(typeName).getRelations().containsKey(relationName)) {
            return BOTH_FALSE;
        }

        var relationsMetada = getNullSafe(typeMap.get(typeName).getMetadata(), Metadata::getRelations);
        if (rewrite.getThis() != null) {
            if (relationsMetada != null) {
                var relationMetadata = relationsMetada.get(relationName);
                var relatedTypes = getNullSafeList(relationMetadata, RelationMetadata::getDirectlyRelatedUserTypes);
                for (var assignableType : getTypeRestrictions(relatedTypes)) {
                    var destructuredType = DestructuredTupleToUserset.from(assignableType);
                    var decodedRelation = destructuredType.getDecodedRelation();
                    if (decodedRelation == null || destructuredType.isWildcard()) {
                        return HAS_ENTRY_BUT_NO_LOOP;
                    }

                    var decodedType = destructuredType.getDecodedType();
                    var assignableRelation =
                            typeMap.get(decodedType).getRelations().get(decodedRelation);
                    if (assignableRelation == null) {
                        return BOTH_FALSE;
                    }

                    if (getNullSafe(visited.get(decodedType), m -> m.get(decodedRelation)) != null) {
                        continue;
                    }

                    var entryPointOrLoop = compute(typeMap, decodedType, decodedRelation, assignableRelation, visited);
                    if (entryPointOrLoop.hasEntry()) {
                        return HAS_ENTRY_BUT_NO_LOOP;
                    }
                }
            }
            return BOTH_FALSE;
        } else if (rewrite.getComputedUserset() != null) {
            var computedRelationName = rewrite.getComputedUserset().getRelation();
            if (computedRelationName == null) {
                return BOTH_FALSE;
            }

            var computedRelation = typeMap.get(typeName).getRelations().get(computedRelationName);
            if (computedRelation == null) {
                return BOTH_FALSE;
            }

            // Loop detected
            if (visited.get(typeName).containsKey(computedRelationName)) {
                return NO_ENTRY_WITH_LOOP;
            }

            return compute(typeMap, typeName, computedRelationName, computedRelation, visited);
        } else if (rewrite.getTupleToUserset() != null) {
            var tuplesetRelationName = rewrite.getTupleToUserset().getTupleset().getRelation();
            var computedRelationName =
                    rewrite.getTupleToUserset().getComputedUserset().getRelation();
            if (tuplesetRelationName == null || computedRelationName == null) {
                return BOTH_FALSE;
            }

            if (!typeMap.get(typeName).getRelations().containsKey(tuplesetRelationName)) {
                return BOTH_FALSE;
            }
            if (relationsMetada != null) {
                var relationMetadata = relationsMetada.get(tuplesetRelationName);
                var relatedTypes = getNullSafeList(relationMetadata, RelationMetadata::getDirectlyRelatedUserTypes);
                for (var assignableType : getTypeRestrictions(relatedTypes)) {
                    var assignableRelation =
                            typeMap.get(assignableType).getRelations().get(computedRelationName);
                    if (assignableRelation != null) {
                        if (visited.containsKey(assignableType)
                                && visited.get(assignableType).containsKey(computedRelationName)) {
                            continue;
                        }

                        var entryOrLoop =
                                compute(typeMap, assignableType, computedRelationName, assignableRelation, visited);
                        if (entryOrLoop.hasEntry()) {
                            return HAS_ENTRY_BUT_NO_LOOP;
                        }
                    }
                }
            }
            return BOTH_FALSE;
        } else if (rewrite.getUnion() != null) {
            var loop = false;

            for (var child : rewrite.getUnion().getChild()) {
                var childEntryOrLoop = compute(typeMap, typeName, relationName, child, visited);
                if (childEntryOrLoop.hasEntry()) {
                    return HAS_ENTRY_BUT_NO_LOOP;
                }
                loop = loop || childEntryOrLoop.isLoop();
            }
            return new EntryPointOrLoop(false, loop);
        } else if (rewrite.getIntersection() != null) {
            for (var child : rewrite.getIntersection().getChild()) {
                var childEntryOrLoop = compute(typeMap, typeName, relationName, child, visited);
                if (!childEntryOrLoop.hasEntry()) {
                    return childEntryOrLoop;
                }
            }
            return HAS_ENTRY_BUT_NO_LOOP;
        } else if (rewrite.getDifference() != null) {
            var baseEntryOrLoop = compute(
                    typeMap, typeName, relationName, rewrite.getDifference().getBase(), visited);
            if (!baseEntryOrLoop.hasEntry()) {
                return baseEntryOrLoop;
            }
            var substractEntryOrLoop = compute(
                    typeMap, typeName, relationName, rewrite.getDifference().getSubtract(), visited);
            if (!substractEntryOrLoop.hasEntry()) {
                return substractEntryOrLoop;
            }
            return HAS_ENTRY_BUT_NO_LOOP;
        }
        return BOTH_FALSE;
    }

    public boolean hasEntry() {
        return entry;
    }

    public boolean isLoop() {
        return loop;
    }
}
