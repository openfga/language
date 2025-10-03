using System.Text;
using OpenFga.Sdk.Model;
using OpenFgaLanguage.Errors;
using OpenFgaLanguage.Utils;

namespace OpenFgaLanguage.Transformers;

public class JsonToDslTransformer
{
    private static readonly string EOL = Environment.NewLine;

    public string Transform(string json)
    {
        AuthorizationModel? model = null;
        if (!string.IsNullOrEmpty(json) && json != "null")
        {
            model = Json.Parse<AuthorizationModel>(json);
        }
        return TransformJsonToDsl(model);
    }

    private string TransformJsonToDsl(AuthorizationModel? model)
    {
        var schemaVersion = "1.1";
        if (model?.SchemaVersion != null)
        {
            schemaVersion = model.SchemaVersion;
        }

        var formattedTypeDefinitions = new StringBuilder();
        if (model?.TypeDefinitions != null)
        {
            foreach (var typeDefinition in model.TypeDefinitions)
            {
                formattedTypeDefinitions.Append(FormatType(typeDefinition)).Append(EOL);
            }
        }

        var formattedConditions = FormatConditions(model);

        return $"model{EOL}  schema {schemaVersion}{EOL}{formattedTypeDefinitions}{formattedConditions}";
    }

    private string FormatType(TypeDefinition typeDef)
    {
        var typeName = typeDef.Type;
        var formattedTypeBuilder = new StringBuilder(EOL).Append("type ").Append(typeName);

        var relations = typeDef.Relations ?? new Dictionary<string, Userset>();
        var metadata = typeDef.Metadata;
        var emptyMetadataRelation = new Dictionary<string, RelationMetadata>();
        var metadataRelations = metadata?.Relations ?? emptyMetadataRelation;

        if (relations.Count > 0)
        {
            formattedTypeBuilder.Append(EOL).Append("  relations");
            foreach (var relationEntry in relations)
            {
                var relationName = relationEntry.Key;
                var relationDefinition = relationEntry.Value;
                var formattedRelationString = FormatRelation(typeName, relationName, relationDefinition, 
                    metadataRelations.GetValueOrDefault(relationName));
                formattedTypeBuilder.Append(EOL).Append(formattedRelationString);
            }
        }

        return formattedTypeBuilder.ToString();
    }

    private string FormatRelation(string typeName, string relationName, Userset relationDefinition, RelationMetadata? relationMetadata)
    {
        var validator = new DirectAssignmentValidator();

        var typeRestrictions = relationMetadata?.DirectlyRelatedUserTypes ?? new List<RelationReference>();

        Func<string, string, Userset, List<RelationReference>, DirectAssignmentValidator, StringBuilder> formatter = FormatSubRelation;

        if (relationDefinition.Difference != null)
        {
            formatter = FormatDifference;
        }
        else if (relationDefinition.Union != null)
        {
            formatter = FormatUnion;
        }
        else if (relationDefinition.Intersection != null)
        {
            formatter = FormatIntersection;
        }

        var formattedRelation = formatter(typeName, relationName, relationDefinition, typeRestrictions, validator);
        
        if (validator.Occurrences() == 0 || (validator.Occurrences() == 1 && validator.IsFirstPosition(relationDefinition)))
        {
            return $"    define {relationName}: {formattedRelation}";
        }

        throw new UnsupportedDSLNestingException(typeName, relationName);
    }

    private StringBuilder FormatDifference(string typeName, string relationName, Userset relationDefinition, 
        List<RelationReference> typeRestrictions, DirectAssignmentValidator validator)
    {
        var baseRelation = FormatSubRelation(typeName, relationName, relationDefinition.Difference.Base, typeRestrictions, validator);
        var difference = FormatSubRelation(typeName, relationName, relationDefinition.Difference.Subtract, typeRestrictions, validator);
        return new StringBuilder(baseRelation.ToString()).Append(" but not ").Append(difference);
    }

    private StringBuilder FormatUnion(string typeName, string relationName, Userset relationDefinition, 
        List<RelationReference> typeRestrictions, DirectAssignmentValidator validator)
    {
        return JoinChildren(relationDefinition.Union?.Child, "or", typeName, relationName, relationDefinition, typeRestrictions, validator);
    }

    private StringBuilder FormatIntersection(string typeName, string relationName, Userset relationDefinition, 
        List<RelationReference> typeRestrictions, DirectAssignmentValidator validator)
    {
        return JoinChildren(relationDefinition.Intersection?.Child, "and", typeName, relationName, relationDefinition, typeRestrictions, validator);
    }

    private StringBuilder JoinChildren(List<Userset>? children, string op, string typeName, string relationName, 
        Userset relationDefinition, List<RelationReference> typeRestrictions, DirectAssignmentValidator validator)
    {
        children = PrioritizeDirectAssignment(children ?? new List<Userset>());

        var formattedUnion = new StringBuilder();
        var notFirst = false;
        foreach (var child in children)
        {
            if (notFirst)
            {
                formattedUnion.Append($" {op} ");
            }
            else
            {
                notFirst = true;
            }
            formattedUnion.Append(FormatSubRelation(typeName, relationName, child, typeRestrictions, validator));
        }

        return formattedUnion;
    }

    private static List<Userset> PrioritizeDirectAssignment(List<Userset> usersets)
    {
        if (usersets.Count > 0)
        {
            var thisPosition = usersets.FindIndex(u => u.This != null);
            if (thisPosition > 0)
            {
                var thisUserset = usersets[thisPosition];
                usersets.RemoveAt(thisPosition);
                usersets.Insert(0, thisUserset);
            }
        }

        return usersets;
    }

    private class DirectAssignmentValidator
    {
        private int _occurred = 0;

        public void Incr()
        {
            _occurred++;
        }

        public int Occurrences()
        {
            return _occurred;
        }

        public bool IsFirstPosition(Userset userset)
        {
            if (userset.This != null)
            {
                return true;
            }

            if (userset.Difference?.Base != null)
            {
                if (userset.Difference.Base.This != null)
                {
                    return true;
                }
                else
                {
                    return IsFirstPosition(userset.Difference.Base);
                }
            }
            else if (userset.Intersection?.Child != null && userset.Intersection.Child.Count > 0)
            {
                if (userset.Intersection.Child[0].This != null)
                {
                    return true;
                }
                else
                {
                    return IsFirstPosition(userset.Intersection.Child[0]);
                }
            }
            else if (userset.Union?.Child != null && userset.Union.Child.Count > 0)
            {
                if (userset.Union.Child[0].This != null)
                {
                    return true;
                }
                else
                {
                    return IsFirstPosition(userset.Union.Child[0]);
                }
            }
            return false;
        }
    }

    private StringBuilder FormatSubRelation(string typeName, string relationName, Userset relationDefinition, 
        List<RelationReference> typeRestrictions, DirectAssignmentValidator validator)
    {
        if (relationDefinition.This != null)
        {
            validator.Incr();
            return new StringBuilder(FormatThis(typeRestrictions));
        }

        if (relationDefinition.ComputedUserset != null)
        {
            return new StringBuilder(FormatComputedUserset(relationDefinition));
        }

        if (relationDefinition.TupleToUserset != null)
        {
            return new StringBuilder(FormatTupleToUserset(relationDefinition));
        }

        if (relationDefinition.Union != null)
        {
            return FormatUnion(typeName, relationName, relationDefinition, typeRestrictions, validator)
                .Insert(0, '(')
                .Append(')');
        }

        if (relationDefinition.Intersection != null)
        {
            return FormatIntersection(typeName, relationName, relationDefinition, typeRestrictions, validator)
                .Insert(0, '(')
                .Append(')');
        }

        if (relationDefinition.Difference != null)
        {
            return FormatDifference(typeName, relationName, relationDefinition, typeRestrictions, validator)
                .Insert(0, '(')
                .Append(')');
        }

        throw new UnsupportedDSLNestingException(typeName, relationName);
    }

    private string FormatThis(List<RelationReference> typeRestrictions)
    {
        var restrictions = typeRestrictions ?? new List<RelationReference>();
        return "[" + string.Join(", ", restrictions.Select(FormatTypeRestriction)) + "]";
    }

    private string FormatTypeRestriction(RelationReference restriction)
    {
        var typeName = restriction.Type;
        var relation = restriction.Relation;
        var wildcard = restriction.Wildcard;
        var condition = restriction.Condition;

        var formattedTypeRestriction = new StringBuilder(typeName);

        if (wildcard != null)
        {
            formattedTypeRestriction.Append(":*");
        }

        if (!string.IsNullOrEmpty(relation))
        {
            formattedTypeRestriction.Append('#').Append(relation);
        }

        if (!string.IsNullOrEmpty(condition))
        {
            formattedTypeRestriction.Append(" with ").Append(condition);
        }
        
        return formattedTypeRestriction.ToString();
    }

    private string FormatComputedUserset(Userset relationDefinition)
    {
        return relationDefinition.ComputedUserset?.Relation ?? string.Empty;
    }

    private string FormatTupleToUserset(Userset relationDefinition)
    {
        var computedUserset = string.Empty;
        var tupleset = string.Empty;
        
        if (relationDefinition?.TupleToUserset != null)
        {
            if (relationDefinition.TupleToUserset.ComputedUserset != null)
            {
                computedUserset = relationDefinition.TupleToUserset.ComputedUserset.Relation ?? string.Empty;
            }
            if (relationDefinition.TupleToUserset.Tupleset != null)
            {
                tupleset = relationDefinition.TupleToUserset.Tupleset.Relation ?? string.Empty;
            }
        }
        
        return new StringBuilder(computedUserset).Append(" from ").Append(tupleset).ToString();
    }

    private string FormatConditions(AuthorizationModel model)
    {
        var conditions = model?.Conditions;
        if (conditions == null || conditions.Count == 0)
        {
            return string.Empty;
        }

        var formattedConditions = new StringBuilder();
        var sortedConditions = new SortedDictionary<string, Condition>(conditions);

        foreach (var conditionEntry in sortedConditions)
        {
            var conditionName = conditionEntry.Key;
            var conditionDef = conditionEntry.Value;

            var formattedCondition = FormatCondition(conditionName, conditionDef);
            formattedConditions.Append(EOL).Append(formattedCondition);
        }

        return formattedConditions.ToString();
    }

    private string FormatCondition(string conditionName, Condition conditionDef)
    {
        if (!conditionName.Equals(conditionDef.Name))
        {
            throw new ArgumentException("conditionName must match condition.Name()");
        }

        var formattedParameters = FormatConditionParameters(conditionDef.Parameters);
        return new StringBuilder("condition ")
            .Append(conditionDef.Name)
            .Append('(')
            .Append(formattedParameters)
            .Append(") {")
            .Append(EOL)
            .Append("  ")
            .Append(conditionDef.Expression)
            .Append(EOL)
            .Append('}')
            .Append(EOL)
            .ToString();
    }

    private string FormatConditionParameters(Dictionary<string, ConditionParamTypeRef>? parameters)
    {
        if (parameters == null || parameters.Count == 0)
        {
            return string.Empty;
        }

        var sortedParameters = new SortedDictionary<string, ConditionParamTypeRef>(parameters);
        return string.Join(", ", sortedParameters.Select(entry =>
        {
            var parameterName = entry.Key;
            var parameterType = entry.Value;
            var formattedParameterType = parameterType.TypeName.ToString()
                .Replace("TYPE_NAME_", "")
                .ToLowerInvariant();
            
            if (formattedParameterType == "list" || formattedParameterType == "map")
            {
                var genericTypeString = string.Empty;
                if (parameterType.GenericTypes != null && parameterType.GenericTypes.Count > 0)
                {
                    genericTypeString = parameterType.GenericTypes[0].TypeName.ToString()
                        .Replace("TYPE_NAME_", "")
                        .ToLowerInvariant();
                }
                formattedParameterType = $"{formattedParameterType}<{genericTypeString}>";
            }
            
            return $"{parameterName}: {formattedParameterType}";
        }));
    }
}
