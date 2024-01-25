package dev.openfga.language;

import dev.openfga.sdk.api.model.RelationMetadata;
import dev.openfga.sdk.api.model.Userset;
import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.Setter;

import java.util.List;

@Getter
@Setter
@AllArgsConstructor
final class Relation {
    private String name;
    private List<Userset> rewrites;
    private String operator;
    private RelationMetadata typeInfo;
}