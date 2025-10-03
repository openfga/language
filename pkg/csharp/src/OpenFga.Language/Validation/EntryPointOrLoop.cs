using OpenFga.Sdk.Model;

namespace OpenFga.Language.Validation;

internal class EntryPointOrLoop
{
    public static readonly EntryPointOrLoop BothFalse = new(false, false);
    public static readonly EntryPointOrLoop HasEntryButNoLoop = new(true, false);
    public static readonly EntryPointOrLoop NoEntryWithLoop = new(false, true);

    private readonly bool _entry;
    private readonly bool _loop;

    public EntryPointOrLoop(bool hasEntry, bool isLoop)
    {
        _entry = hasEntry;
        _loop = isLoop;
    }

    // for the type/relation, whether there are any unique entry points, and if a loop is found
    // if there are unique entry points (i.e., direct relations) then it will return true
    // otherwise, it will follow its children to see if there are unique entry points
    // if there is a loop during traversal, the function will return a boolean indicating so
    public static EntryPointOrLoop Compute(
        Dictionary<string, TypeDefinition> typeMap,
        string typeName,
        string relationName,
        Userset rewrite,
        Dictionary<string, Dictionary<string, bool>> visitedRecords)
    {
        var visited = Utils.DeepCopy(visitedRecords);

        if (relationName == null)
        {
            return BothFalse;
        }

        if (!visited.ContainsKey(typeName))
        {
            visited[typeName] = new Dictionary<string, bool>();
        }
        visited[typeName][relationName] = true;

        var currentRelations = typeMap[typeName].Relations;
        if (currentRelations == null || !currentRelations.ContainsKey(relationName))
        {
            return BothFalse;
        }

        if (typeMap[typeName].Relations == null || !typeMap[typeName].Relations.ContainsKey(relationName))
        {
            return BothFalse;
        }

        var relationsMetadata = Utils.GetNullSafe(typeMap[typeName].Metadata, m => m.Relations);
        if (rewrite.This != null)
        {
            if (relationsMetadata != null)
            {
                var relationMetadata = relationsMetadata.GetValueOrDefault(relationName);
                var relatedTypes = Utils.GetNullSafeList(relationMetadata, rm => rm.DirectlyRelatedUserTypes);
                foreach (var assignableType in Dsl.GetTypeRestrictions(relatedTypes))
                {
                    var destructuredType = DestructuredTupleToUserset.From(assignableType);
                    var decodedRelation = destructuredType.DecodedRelation;
                    if (decodedRelation == null || destructuredType.IsWildcard)
                    {
                        return HasEntryButNoLoop;
                    }

                    var decodedType = destructuredType.DecodedType;
                    var assignableRelation = typeMap[decodedType].Relations?.GetValueOrDefault(decodedRelation);
                    if (assignableRelation == null)
                    {
                        return BothFalse;
                    }

                    if (Utils.GetNullSafe(visited.GetValueOrDefault(decodedType), m => m.GetValueOrDefault(decodedRelation)) != null)
                    {
                        continue;
                    }

                    var entryPointOrLoop = Compute(typeMap, decodedType, decodedRelation, assignableRelation, visited);
                    if (entryPointOrLoop.HasEntry())
                    {
                        return HasEntryButNoLoop;
                    }
                }
            }
            return BothFalse;
        }
        else if (rewrite.ComputedUserset != null)
        {
            var computedRelationName = rewrite.ComputedUserset.Relation;
            if (computedRelationName == null)
            {
                return BothFalse;
            }

            var computedRelation = typeMap[typeName].Relations?.GetValueOrDefault(computedRelationName);
            if (computedRelation == null)
            {
                return BothFalse;
            }

            // Loop detected
            if (visited[typeName].ContainsKey(computedRelationName))
            {
                return NoEntryWithLoop;
            }

            return Compute(typeMap, typeName, computedRelationName, computedRelation, visited);
        }
        else if (rewrite.TupleToUserset != null)
        {
            var tuplesetRelationName = rewrite.TupleToUserset.Tupleset?.Relation;
            var computedRelationName = rewrite.TupleToUserset.ComputedUserset?.Relation;
            if (tuplesetRelationName == null || computedRelationName == null)
            {
                return BothFalse;
            }

            if (!typeMap[typeName].Relations.ContainsKey(tuplesetRelationName))
            {
                return BothFalse;
            }
            if (relationsMetadata != null)
            {
                var relationMetadata = relationsMetadata.GetValueOrDefault(tuplesetRelationName);
                var relatedTypes = Utils.GetNullSafeList(relationMetadata, rm => rm.DirectlyRelatedUserTypes);
                foreach (var assignableType in Dsl.GetTypeRestrictions(relatedTypes))
                {
                    var assignableRelation = typeMap[assignableType].Relations?.GetValueOrDefault(computedRelationName);
                    if (assignableRelation != null)
                    {
                        if (visited.ContainsKey(assignableType) && visited[assignableType].ContainsKey(computedRelationName))
                        {
                            continue;
                        }

                        var entryOrLoop = Compute(typeMap, assignableType, computedRelationName, assignableRelation, visited);
                        if (entryOrLoop.HasEntry())
                        {
                            return HasEntryButNoLoop;
                        }
                    }
                }
            }
            return BothFalse;
        }
        else if (rewrite.Union != null)
        {
            var loop = false;

            foreach (var child in rewrite.Union.Child)
            {
                var childEntryOrLoop = Compute(typeMap, typeName, relationName, child, visited);
                if (childEntryOrLoop.HasEntry())
                {
                    return HasEntryButNoLoop;
                }
                loop = loop || childEntryOrLoop.IsLoop();
            }
            return new EntryPointOrLoop(false, loop);
        }
        else if (rewrite.Intersection != null)
        {
            foreach (var child in rewrite.Intersection.Child)
            {
                var childEntryOrLoop = Compute(typeMap, typeName, relationName, child, visited);
                if (!childEntryOrLoop.HasEntry())
                {
                    return childEntryOrLoop;
                }
            }
            return HasEntryButNoLoop;
        }
        else if (rewrite.Difference != null)
        {
            var baseEntryOrLoop = Compute(typeMap, typeName, relationName, rewrite.Difference.Base, visited);
            if (!baseEntryOrLoop.HasEntry())
            {
                return baseEntryOrLoop;
            }
            var subtractEntryOrLoop = Compute(typeMap, typeName, relationName, rewrite.Difference.Subtract, visited);
            if (!subtractEntryOrLoop.HasEntry())
            {
                return subtractEntryOrLoop;
            }
            return HasEntryButNoLoop;
        }
        return BothFalse;
    }

    public bool HasEntry() => _entry;
    public bool IsLoop() => _loop;
}
