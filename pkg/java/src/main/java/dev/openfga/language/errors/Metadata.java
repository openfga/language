package dev.openfga.language.errors;

public class Metadata {
    private String symbol;

    public Metadata() {}

    public Metadata(String symbol) {
        this.symbol = symbol;
    }

    public String getSymbol() {
        return symbol;
    }

    public void setSymbol(String symbol) {
        this.symbol = symbol;
    }
}
