package dev.openfga.language.validation;

import static dev.openfga.language.Utils.getNullSafe;
import static java.util.stream.Collectors.toList;

import dev.openfga.sdk.api.model.ObjectRelation;
import dev.openfga.sdk.api.model.RelationReference;
import dev.openfga.sdk.api.model.Userset;
import java.util.Collection;
import java.util.List;
import java.util.function.Predicate;
import java.util.regex.Pattern;
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

        return IntStream.range(Math.max(skipIndex, 0), lines.length)
                .filter(index -> predicate.test(lines[index]))
                .findFirst()
                .orElse(-1);
    }

    public int getConditionLineNumber(String conditionName) {
        return getConditionLineNumber(conditionName, 0);
    }

    public int getConditionLineNumber(String conditionName, int skipIndex) {
        // Require `(` after the name so a condition name that is a prefix of
        // another (e.g. `less` vs `less_than`) cannot match the wrong line.
        return findLine(
                line -> line.trim().matches("condition " + Pattern.quote(conditionName) + "\\s*\\(.*"), skipIndex);
    }

    public int getRelationLineNumber(String relationName, int skipIndex) {
        // Require `:` after the name so a relation name that is a prefix of
        // another (e.g. `writer` vs `writers`) cannot match the wrong line.
        return findLine(
                line -> line.trim()
                        .replaceAll(" {2,}", " ")
                        .matches("define " + Pattern.quote(relationName) + "\\s*:.*"),
                skipIndex);
    }

    public int getSchemaLineNumber(String schemaVersion) {
        // Allow only whitespace or a trailing comment after the version so
        // e.g. `1.1` cannot match `schema 1.10`. A comment must be preceded by
        // whitespace so a `#` glued to the version isn't treated as a comment.
        return findLine(
                line -> line.trim()
                        .replaceAll(" {2,}", " ")
                        .matches("schema " + Pattern.quote(schemaVersion) + "(\\s+#.*)?"),
                0);
    }

    public int getTypeLineNumber(String typeName) {
        return getTypeLineNumber(typeName, 0);
    }

    public int getTypeLineNumber(String typeName, int skipIndex) {
        // Allow an optional trailing comment (e.g. `type page # module: ...`) after the type name.
        // The comment must be preceded by whitespace so a `#` glued to the name isn't treated as a comment.
        // Quote the type name so regex metacharacters (e.g. `.`) are matched literally.
        return findLine(line -> line.trim().matches("type " + Pattern.quote(typeName) + "(\\s+#.*)?"), skipIndex);
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
