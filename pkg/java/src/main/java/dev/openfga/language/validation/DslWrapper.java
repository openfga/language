package dev.openfga.language.validation;

import java.util.function.Predicate;
import java.util.stream.IntStream;

class DslWrapper {

    private final String[] lines;

    DslWrapper(String[] lines) {
        this.lines = lines;
    }

    private int findLine(Predicate<String> predicate, int skipIndex) {
        return IntStream.range(skipIndex, lines.length)
                .filter(index -> predicate.test(lines[index]))
                .findFirst().orElse(-1);
    }

    public int getConditionLineNumber(String conditionName) {
        return getConditionLineNumber(conditionName, 0);
    }

    public int getConditionLineNumber(String conditionName, int skipIndex) {
        return findLine(
                line -> line.trim().startsWith("condition " + conditionName),
                skipIndex);
    }

    public int getRelationLineNumber(String relationName, int skipIndex) {
        return findLine(
                line -> line.trim().replaceAll(" {2,}", " ").startsWith("define " + relationName),
                skipIndex);
    }

    public int getSchemaLineNumber(String schemaVersion) {
        return findLine(
                line -> line.trim().replaceAll(" {2,}", " ").startsWith("schema " + schemaVersion),
                0);
    }

    public int getTypeLineNumber(String typeName) {
        return getTypeLineNumber(typeName, 0);
    }

    public int getTypeLineNumber(String typeName, int skipIndex) {
        return findLine(
                line -> line.trim().startsWith("type " + typeName),
                skipIndex);
    }

}
