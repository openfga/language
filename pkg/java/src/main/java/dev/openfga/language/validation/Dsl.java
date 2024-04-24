package dev.openfga.language.validation;

import static dev.openfga.language.Utils.getNullSafe;
import static java.util.stream.Collectors.toList;

import dev.openfga.sdk.api.model.ObjectRelation;
import dev.openfga.sdk.api.model.RelationReference;
import dev.openfga.sdk.api.model.Userset;
import java.util.Collection;
import java.util.List;
import java.util.function.Predicate;
import java.util.stream.IntStream;

class Dsl {

    private final String[] lines;

    Dsl(String[] lines) {
        this.lines = lines;
    }

    private int findLine(Predicate<String> predicate, int skipIndex) {
        if (lines == null) {
            return -1;
        }

        return IntStream.range(skipIndex, lines.length)
                .filter(index -> predicate.test(lines[index]))
                .findFirst()
                .orElse(-1);
    }

    public int getConditionLineNumber(String conditionName) {
        return getConditionLineNumber(conditionName, 0);
    }

    public int getConditionLineNumber(String conditionName, int skipIndex) {
        return findLine(line -> line.trim().startsWith("condition " + conditionName), skipIndex);
    }

    public int getRelationLineNumber(String relationName, int skipIndex) {
        return findLine(line -> line.trim().replaceAll(" {2,}", " ").startsWith("define " + relationName), skipIndex);
    }

    public int getSchemaLineNumber(String schemaVersion) {
        return findLine(line -> line.trim().replaceAll(" {2,}", " ").startsWith("schema " + schemaVersion), 0);
    }

    public int getTypeLineNumber(String typeName) {
        return getTypeLineNumber(typeName, 0);
    }

    public int getTypeLineNumber(String typeName, int skipIndex) {
        return findLine(line -> line.trim().startsWith("type " + typeName), skipIndex);
    }

    public static String getRelationDefName(Userset userset) {
        var relationDefName = getNullSafe(userset.getComputedUserset(), ObjectRelation::getRelation);
        var parserResult = getRelationalParserResult(userset);
        if (parserResult.getRewrite() == RewriteType.ComputedUserset) {
            relationDefName = parserResult.getTarget();
        } else if (parserResult.getRewrite() == RewriteType.TupleToUserset) {
            relationDefName = parserResult.getTarget() + " from " + parserResult.getFrom();
        }
        return relationDefName;
    }

    public static RelationTargetParserResult getRelationalParserResult(Userset userset) {
        String target = null, from = null;

        if (userset.getComputedUserset() != null) {
            target = userset.getComputedUserset().getRelation();
        } else {
            if (userset.getTupleToUserset() != null
                    && userset.getTupleToUserset().getComputedUserset() != null) {
                target = userset.getTupleToUserset().getComputedUserset().getRelation();
            }
            if (userset.getTupleToUserset() != null
                    && userset.getTupleToUserset().getTupleset() != null) {
                from = userset.getTupleToUserset().getTupleset().getRelation();
            }
        }

        var rewrite = RewriteType.Direct;
        if (target != null) {
            rewrite = RewriteType.ComputedUserset;
        }

        if (from != null) {
            rewrite = RewriteType.TupleToUserset;
        }
        return new RelationTargetParserResult(target, from, rewrite);
    }

    public static List<String> getTypeRestrictions(Collection<RelationReference> relatedTypes) {
        return relatedTypes.stream().map(Dsl::getTypeRestrictionString).collect(toList());
    }

    public static String getTypeRestrictionString(RelationReference typeRestriction) {
        var typeRestrictionString = typeRestriction.getType();
        if (typeRestriction.getWildcard() != null) {
            typeRestrictionString += ":*";
        } else if (typeRestriction.getRelation() != null) {
            typeRestrictionString += "#" + typeRestriction.getRelation();
        }

        if (typeRestriction.getCondition() != null) {
            typeRestrictionString += " with " + typeRestriction.getCondition();
        }

        return typeRestrictionString;
    }
}
