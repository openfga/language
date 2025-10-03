using OpenFga.Sdk.Model;

namespace OpenFga.Language.utils;

public static class ModelUtils
{
    /// <summary>
    /// getModuleForObjectTypeRelation returns the module for the given object type and relation in that type.
    /// </summary>
    /// <param name="typeDef">A TypeDefinition object which contains metadata about the type.</param>
    /// <param name="relation">A string representing the relation whose module is to be retrieved.</param>
    /// <returns>A string representing the module for the given object type and relation.</returns>
    /// <exception cref="Exception">An error if the relation does not exist.</exception>
    public static string? GetModuleForObjectTypeRelation(TypeDefinition typeDef, string relation)
    {
        var relations = typeDef.Relations;
        if (relations == null || !relations.ContainsKey(relation))
        {
            throw new Exception($"relation {relation} does not exist in type {typeDef.Type}");
        }

        var metadata = typeDef.Metadata;
        if (metadata == null || metadata.Relations == null)
        {
            return null;
        }

        var relationMetadata = metadata.Relations.GetValueOrDefault(relation);
        if (relationMetadata == null
            || string.IsNullOrEmpty(relationMetadata.Module))
        {
            return metadata.Module;
        }

        return relationMetadata.Module;
    }

    /// <summary>
    /// isRelationAssignable returns true if the relation is assignable, as in the relation definition has a key "this" or
    /// any of its children have a key "this".
    /// </summary>
    /// <param name="relDef">A Userset object representing a relation definition.</param>
    /// <returns>A boolean representing whether the relation definition has a key "this".</returns>
    public static bool IsRelationAssignable(Userset? relDef)
    {
        if (relDef == null)
        {
            return false;
        }

        if (relDef.This != null)
        {
            return true;
        }
        else if (relDef.Union != null)
        {
            foreach (var child in relDef.Union.Child)
            {
                if (IsRelationAssignable(child))
                {
                    return true;
                }
            }
        }
        else if (relDef.Intersection != null)
        {
            foreach (var child in relDef.Intersection.Child)
            {
                if (IsRelationAssignable(child))
                {
                    return true;
                }
            }
        }
        else if (relDef.Difference != null)
        {
            var diff = relDef.Difference;
            if (IsRelationAssignable(diff.Base) || IsRelationAssignable(diff.Subtract))
            {
                return true;
            }
        }

        // ComputedUserset and TupleToUserset are not assignable
        return false;
    }
}
