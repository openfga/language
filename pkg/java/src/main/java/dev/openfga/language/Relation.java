package dev.openfga.language;

import dev.openfga.sdk.api.model.RelationMetadata;
import dev.openfga.sdk.api.model.Userset;

import java.util.List;

final class Relation {
    private String name;
    private List<Userset> rewrites;
    private String operator;
    private RelationMetadata typeInfo;

    public Relation(String name, List<Userset> rewrites, String operator, RelationMetadata typeInfo) {
        this.name = name;
        this.rewrites = rewrites;
        this.operator = operator;
        this.typeInfo = typeInfo;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public List<Userset> getRewrites() {
        return rewrites;
    }

    public void setRewrites(List<Userset> rewrites) {
        this.rewrites = rewrites;
    }

    public String getOperator() {
        return operator;
    }

    public void setOperator(String operator) {
        this.operator = operator;
    }

    public RelationMetadata getTypeInfo() {
        return typeInfo;
    }

    public void setTypeInfo(RelationMetadata typeInfo) {
        this.typeInfo = typeInfo;
    }
}