package dev.openfga.language;

import dev.openfga.language.errors.StartEnd;
import java.util.regex.Pattern;

/** Resolves line and column positions of symbols within a single module's DSL, for error reporting. */
final class ModuleDsl {

    // Space, tab and form feed are exactly what the lexer's WHITESPACE rule admits between tokens (OpenFGALexer.g4).
    private static final Pattern INLINE_WHITESPACE = Pattern.compile("[ \\t\\f]+");

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

    /**
     * Returns the index of the first line that declares {@code prefix}, or -1. The prefix always spans whole words
     * ({@code type <name>}, {@code define <name>}, ...), so requiring the name to end where the prefix ends is what
     * stops {@code document} matching a declaration of {@code documentation}. Runs of inline whitespace are collapsed
     * first so {@code define\towner:} is found like {@code define owner:}.
     */
    private int findLine(String prefix) {
        for (int i = 0; i < lines.length; i++) {
            var line = INLINE_WHITESPACE.matcher(lines[i].trim()).replaceAll(" ");
            if (line.startsWith(prefix)
                    && (line.length() == prefix.length() || !isNameChar(line.charAt(prefix.length())))) {
                return i;
            }
        }
        return -1;
    }

    // IDENTIFIER admits letters, digits, `_` and MINUS; EXTENDED_IDENTIFIER adds SLASH and DOT (OpenFGALexer.g4), so
    // `core.doc` is one name rather than `core` followed by a terminator.
    private static boolean isNameChar(char c) {
        return (c >= 'a' && c <= 'z')
                || (c >= 'A' && c <= 'Z')
                || (c >= '0' && c <= '9')
                || c == '_'
                || c == '-'
                || c == '/'
                || c == '.';
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
