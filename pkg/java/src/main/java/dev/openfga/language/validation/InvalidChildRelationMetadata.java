package dev.openfga.language.validation;

class InvalidChildRelationMetadata {
    private final int lineIndex;
    private final String symbol;
    private final String typeName;
    private final String relationName;

    public InvalidChildRelationMetadata(int lineIndex, String symbol, String typeName, String relationName) {
        this.lineIndex = lineIndex;
        this.symbol = symbol;
        this.typeName = typeName;
        this.relationName = relationName;
    }

    public int getLineIndex() {
        return lineIndex;
    }

    public String getSymbol() {
        return symbol;
    }

    public String getTypeName() {
        return typeName;
    }

    public String getRelationName() {
        return relationName;
    }
}
