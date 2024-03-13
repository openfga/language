package dev.openfga.language.validation;

public class RelationTargetParserResult {

    private final String target;
    private final String from;
    private final RewriteType rewrite;

    public RelationTargetParserResult(String target, String from, RewriteType rewrite) {
        this.target = target;
        this.from = from;
        this.rewrite = rewrite;
    }

    public String getTarget() {
        return target;
    }

    public String getFrom() {
        return from;
    }

    public RewriteType getRewrite() {
        return rewrite;
    }
}
