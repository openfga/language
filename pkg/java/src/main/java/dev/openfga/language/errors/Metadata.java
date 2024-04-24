package dev.openfga.language.errors;

public class Metadata {
    private String symbol;

    private String module;

    public Metadata() {}

    public Metadata(String symbol) {
        this.symbol = symbol;
    }

    public Metadata(String symbol, String module) {
        this.symbol = symbol;
        this.module = module;
    }

    public String getSymbol() {
        return symbol;
    }

    public void setSymbol(String symbol) {
        this.symbol = symbol;
    }

    public String getModule() {
        return module;
    }

    public void setModule(String module) {
        this.module = module;
    }
}
