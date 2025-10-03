using OpenFga.Sdk.Model;

namespace OpenFga.Language.Validation;

internal class Dsl
{
    private readonly string[]? _lines;

    public Dsl(string[]? lines)
    {
        _lines = lines;
    }

    private int FindLine(Func<string, bool> predicate, int skipIndex)
    {
        if (_lines == null)
        {
            return -1;
        }

        for (int i = skipIndex; i < _lines.Length; i++)
        {
            if (predicate(_lines[i]))
            {
                return i;
            }
        }
        return -1;
    }

    public int GetConditionLineNumber(string conditionName)
    {
        return GetConditionLineNumber(conditionName, 0);
    }

    public int GetConditionLineNumber(string conditionName, int skipIndex)
    {
        return FindLine(line => line.Trim().StartsWith($"condition {conditionName}"), skipIndex);
    }

    public int GetRelationLineNumber(string relationName, int skipIndex)
    {
        return FindLine(line => System.Text.RegularExpressions.Regex.Replace(line.Trim(), @" {2,}", " ").StartsWith($"define {relationName}"), skipIndex);
    }

    public int GetSchemaLineNumber(string schemaVersion)
    {
        return FindLine(line => System.Text.RegularExpressions.Regex.Replace(line.Trim(), @" {2,}", " ").StartsWith($"schema {schemaVersion}"), 0);
    }

    public int GetTypeLineNumber(string typeName)
    {
        return GetTypeLineNumber(typeName, 0);
    }

    public int GetTypeLineNumber(string typeName, int skipIndex)
    {
        return FindLine(line => System.Text.RegularExpressions.Regex.IsMatch(line.Trim(), $"type {typeName}"), skipIndex);
    }

    public static string? GetRelationDefName(Userset userset)
    {
        var relationDefName = Utils.GetNullSafe(userset.ComputedUserset, cu => cu.Relation);
        var parserResult = GetRelationalParserResult(userset);
        if (parserResult.Rewrite == RewriteType.ComputedUserset)
        {
            relationDefName = parserResult.Target;
        }
        else if (parserResult.Rewrite == RewriteType.TupleToUserset)
        {
            relationDefName = $"{parserResult.Target} from {parserResult.From}";
        }
        return relationDefName;
    }

    public static RelationTargetParserResult GetRelationalParserResult(Userset userset)
    {
        string? target = null, from = null;

        if (userset.ComputedUserset != null)
        {
            target = userset.ComputedUserset.Relation;
        }
        else
        {
            if (userset.TupleToUserset?.ComputedUserset != null)
            {
                target = userset.TupleToUserset.ComputedUserset.Relation;
            }
            if (userset.TupleToUserset?.Tupleset != null)
            {
                from = userset.TupleToUserset.Tupleset.Relation;
            }
        }

        var rewrite = RewriteType.Direct;
        if (target != null)
        {
            rewrite = RewriteType.ComputedUserset;
        }

        if (from != null)
        {
            rewrite = RewriteType.TupleToUserset;
        }
        return new RelationTargetParserResult(target, from, rewrite);
    }

    public static List<string> GetTypeRestrictions(IEnumerable<RelationReference> relatedTypes)
    {
        return relatedTypes.Select(GetTypeRestrictionString).ToList();
    }

    public static string GetTypeRestrictionString(RelationReference typeRestriction)
    {
        var typeRestrictionString = typeRestriction.Type;
        if (typeRestriction.Wildcard != null)
        {
            typeRestrictionString += ":*";
        }
        else if (typeRestriction.Relation != null)
        {
            typeRestrictionString += "#" + typeRestriction.Relation;
        }

        if (typeRestriction.Condition != null)
        {
            typeRestrictionString += " with " + typeRestriction.Condition;
        }

        return typeRestrictionString;
    }
}
