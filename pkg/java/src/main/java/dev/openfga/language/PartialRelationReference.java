package dev.openfga.language;

import dev.openfga.sdk.api.model.RelationReference;
import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
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
}