package dev.openfga.language;

import dev.openfga.sdk.api.model.RelationReference;

class PartialRelationReference {
    private String type;
    private String relation;
    private Object wildcard;
    private String condition;

    public RelationReference asRelationReference() {
        return new RelationReference()
                .type(type)
                .relation(relation)
                .wildcard(wildcard)
                .condition(condition);
    }

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }

    public String getRelation() {
        return relation;
    }

    public void setRelation(String relation) {
        this.relation = relation;
    }

    public Object getWildcard() {
        return wildcard;
    }

    public void setWildcard(Object wildcard) {
        this.wildcard = wildcard;
    }

    public String getCondition() {
        return condition;
    }

    public void setCondition(String condition) {
        this.condition = condition;
    }
}
