package dev.openfga.language;

import dev.openfga.language.errors.StartEnd;

/** Resolves line and column positions of symbols within a single module's DSL, for error reporting. */
final class ModuleDsl {

    private final String[] lines;

    ModuleDsl(String contents) {
        this.lines = contents.split("\n", -1);
    }

    int getConditionLineNumber(String conditionName) {
        return findLine("condition " + conditionName);
    }

    int getTypeLineNumber(String typeName) {
        return findLine("type " + typeName);
    }

    int getExtendedTypeLineNumber(String typeName) {
        return findLine("extend type " + typeName);
    }

    int getRelationLineNumber(String relationName) {
        return findLine("define " + relationName);
    }

    private int findLine(String prefix) {
        for (int i = 0; i < lines.length; i++) {
            if (lines[i].trim().startsWith(prefix)) {
                return i;
            }
        }
        return -1;
    }

    /** Builds the line and column span of {@code symbol} on {@code lineIndex}, both zero based. */
    StartEnd[] resolve(int lineIndex, String symbol) {
        if (lines.length == 0 || lineIndex == -1) {
            return new StartEnd[] {new StartEnd(0, 0), new StartEnd(0, 0)};
        }

        var wordIdx = Math.max(lines[lineIndex].indexOf(symbol), 0);
        return new StartEnd[] {
            new StartEnd(lineIndex, lineIndex), new StartEnd(wordIdx, wordIdx + symbol.length()),
        };
    }
}
