package dev.openfga.language.validation;

import com.fasterxml.jackson.databind.JsonNode;
import com.networknt.schema.BaseJsonValidator;
import com.networknt.schema.ErrorMessageType;
import com.networknt.schema.ExecutionContext;
import com.networknt.schema.JsonNodePath;
import com.networknt.schema.JsonSchema;
import com.networknt.schema.Keyword;
import com.networknt.schema.SchemaLocation;
import com.networknt.schema.ValidationContext;
import com.networknt.schema.ValidationMessage;
import java.util.Collections;
import java.util.Set;

public class TupleValidator extends BaseJsonValidator {
    private static ErrorMessageType ERROR_MESSAGE_TYPE = new ErrorMessageType() {
        @Override
        public String getErrorCode() {
            return "valid_tuple";
        }
    };

    private final String value;

    public TupleValidator(
            SchemaLocation schemaLocation,
            JsonNodePath evaluationPath,
            JsonNode schemaNode,
            JsonSchema parentSchema,
            Keyword keyword,
            ValidationContext validationContext,
            boolean suppressSubSchemaRetrieval) {
        super(
                schemaLocation,
                evaluationPath,
                schemaNode,
                parentSchema,
                ERROR_MESSAGE_TYPE,
                keyword,
                validationContext,
                suppressSubSchemaRetrieval);
        this.value = schemaNode.textValue();
    }

    @Override
    public Set<ValidationMessage> validate(
            ExecutionContext executionContext, JsonNode node, JsonNode rootNode, JsonNodePath instanceLocation) {
        if (!node.asText().equals(value)) {
            return Collections.singleton(message()
                    .message("{0}: must be equal to ''{1}''")
                    .arguments(value)
                    .instanceLocation(instanceLocation)
                    .instanceNode(node)
                    .build());
        }
        return Collections.emptySet();
    }
}
