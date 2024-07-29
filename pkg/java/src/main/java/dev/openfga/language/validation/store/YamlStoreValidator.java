package dev.openfga.language.validation.store;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.JsonNode;
import com.networknt.schema.ExecutionContext;
import com.networknt.schema.Format;
import com.networknt.schema.InputFormat;
import com.networknt.schema.JsonMetaSchema;
import com.networknt.schema.JsonNodePath;
import com.networknt.schema.JsonSchema;
import com.networknt.schema.JsonSchemaException;
import com.networknt.schema.JsonSchemaFactory;
import com.networknt.schema.JsonValidator;
import com.networknt.schema.Keyword;
import com.networknt.schema.SchemaLocation;
import com.networknt.schema.SchemaValidatorsConfig;
import com.networknt.schema.SpecVersion.VersionFlag;
import com.networknt.schema.ValidationContext;
import com.networknt.schema.ValidationMessage;
import com.networknt.schema.serialization.JsonNodeReader;
import dev.openfga.language.errors.StoreValidationSingleError;
import dev.openfga.language.validation.Validator;
import java.io.IOException;
import java.util.List;
import java.util.Set;
import java.util.stream.Collectors;

public class YamlStoreValidator {

    private JsonSchema schema;

    public static class UserFormat implements Format {

        @Override
        public boolean matches(ExecutionContext executionContext, String value) {
            return Validator.validateUser(value);
        }

        @Override
        public String getName() {
            return "user";
        }
    }

    public static class RelationFormat implements Format {

        @Override
        public boolean matches(ExecutionContext executionContext, String value) {
            return Validator.validateRelation(value);
        }

        @Override
        public String getName() {
            return "relation";
        }
    }

    public static class ObjectFormat implements Format {

        @Override
        public boolean matches(ExecutionContext executionContext, String value) {
            return Validator.validateObject(value);
        }

        @Override
        public String getName() {
            return "object";
        }
    }

    public static class ConditionFormat implements Format {

        @Override
        public boolean matches(ExecutionContext executionContext, String value) {
            return Validator.validateConditionName(value);
        }

        @Override
        public String getName() {
            return "condition";
        }
    }

    public static class TypeFormat implements Format {

        @Override
        public boolean matches(ExecutionContext executionContext, String value) {
            return Validator.validateType(value);
        }

        @Override
        public String getName() {
            return "type";
        }
    }

    public static class TupleKeyword implements Keyword {

        @Override
        public String getValue() {
            return "valid_tuple";
        }

        @Override
        public JsonValidator newValidator(
                SchemaLocation location, JsonNodePath path, JsonNode node, JsonSchema schema, ValidationContext context)
                throws JsonSchemaException, Exception {
            return new TupleValidator(location, path, node, schema, this, context, false);
        }
    }

    public static class StoreKeyword implements Keyword {

        @Override
        public String getValue() {
            return "valid_store";
        }

        @Override
        public JsonValidator newValidator(
                SchemaLocation location, JsonNodePath path, JsonNode node, JsonSchema schema, ValidationContext context)
                throws JsonSchemaException, Exception {
            return new StoreValidator(location, path, node, schema, this, context, false);
        }
    }

    public YamlStoreValidator() throws IOException {
        JsonMetaSchema metaSchema = JsonMetaSchema.builder(JsonMetaSchema.getV202012())
                .format(new UserFormat())
                .format(new RelationFormat())
                .format(new ObjectFormat())
                .format(new ConditionFormat())
                .format(new TypeFormat())
                .keyword(new TupleKeyword())
                .keyword(new StoreKeyword())
                .build();
        JsonSchemaFactory factory =
                JsonSchemaFactory.getInstance(VersionFlag.V202012, builder -> builder.metaSchema(metaSchema)
                        .jsonNodeReader(JsonNodeReader.builder().locationAware().build()));
        SchemaValidatorsConfig config = SchemaValidatorsConfig.builder().build();

        var jsonSchema = new String(
                this.getClass().getResourceAsStream("/json-schema.json").readAllBytes());

        this.schema = factory.getSchema(jsonSchema, InputFormat.JSON, config);
    }

    public List<StoreValidationSingleError> validate(final String store) throws JsonProcessingException {
        Set<ValidationMessage> messages = schema.validate(store, InputFormat.JSON, executionContext -> {
            executionContext.getExecutionConfig().setFormatAssertionsEnabled(true);
        });

        return messages.stream().map(m -> new StoreValidationSingleError(m)).collect(Collectors.toList());
    }
}
