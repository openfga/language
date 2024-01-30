package dev.openfga.language.validation;

import lombok.AllArgsConstructor;
import lombok.Getter;

@Getter
@AllArgsConstructor
public class RelationTargetParserResult {

    private String target;
    private String from;
    private RewriteType rewrite;
}
