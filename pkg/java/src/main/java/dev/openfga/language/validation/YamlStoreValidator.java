package dev.openfga.language.validation;

import java.util.List;
import java.util.Set;
import java.util.stream.Collectors;

import com.networknt.schema.ExecutionContext;
import com.networknt.schema.Format;
import com.networknt.schema.InputFormat;
import com.networknt.schema.JsonMetaSchema;
import com.networknt.schema.JsonSchema;
import com.networknt.schema.JsonSchemaFactory;
import com.networknt.schema.SchemaValidatorsConfig;
import com.networknt.schema.SpecVersion.VersionFlag;
import com.networknt.schema.ValidationMessage;
import com.networknt.schema.serialization.JsonNodeReader;

import dev.openfga.language.errors.StoreValidationSingleError;

public class YamlStoreValidator {

  private final static String JSON_SCHEMA_STRING = "{\n" + //
      "  \"type\": \"object\",\n" + //
      "  \"required\": [\n" + //
      "    \"tests\"\n" + //
      "  ],\n" + //
      "  \"additionalProperties\": false,\n" + //
      "  \"valid_store\": true,\n" + //
      "  \"properties\": {\n" + //
      "    \"name\": {\n" + //
      "      \"type\": \"string\",\n" + //
      "      \"description\": \"the store name\"\n" + //
      "    },\n" + //
      "    \"model_file\": {\n" + //
      "      \"type\": \"string\",\n" + //
      "      \"description\": \"the authorization model file path\"\n" + //
      "    },\n" + //
      "    \"model\": {\n" + //
      "      \"type\": \"string\",\n" + //
      "      \"description\": \"the authorization model (takes precedence over model_file)\"\n" + //
      "    },\n" + //
      "    \"tuple_file\": {\n" + //
      "      \"type\": \"string\",\n" + //
      "      \"description\": \"the tuple file path\"\n" + //
      "    },\n" + //
      "    \"tuples\": {\n" + //
      "      \"type\": \"array\",\n" + //
      "      \"description\": \"the tuples (takes precedence over tuples_file)\",\n" + //
      "      \"items\": {\n" + //
      "        \"type\": \"object\",\n" + //
      "        \"additionalProperties\": false,\n" + //
      "        \"required\": [\n" + //
      "          \"user\",\n" + //
      "          \"relation\",\n" + //
      "          \"object\"\n" + //
      "        ],\n" + //
      "        \"valid_tuple\": true,\n" + //
      "        \"properties\": {\n" + //
      "          \"user\": {\n" + //
      "            \"type\": \"string\",\n" + //
      "            \"format\": \"user\",\n" + //
      "            \"description\": \"the user\"\n" + //
      "          },\n" + //
      "          \"relation\": {\n" + //
      "            \"type\": \"string\",\n" + //
      "            \"format\": \"relation\",\n" + //
      "            \"description\": \"the relation\"\n" + //
      "          },\n" + //
      "          \"object\": {\n" + //
      "            \"type\": \"string\",\n" + //
      "            \"format\": \"object\",\n" + //
      "            \"description\": \"the object\"\n" + //
      "          },\n" + //
      "          \"condition\": {\n" + //
      "            \"type\": \"object\",\n" + //
      "            \"additionalProperties\": false,\n" + //
      "            \"required\": [\n" + //
      "              \"name\"\n" + //
      "            ],\n" + //
      "            \"properties\": {\n" + //
      "              \"name\": {\n" + //
      "                \"type\": \"string\",\n" + //
      "                \"format\": \"condition\"\n" + //
      "              },\n" + //
      "              \"context\": {\n" + //
      "                \"type\": \"object\"\n" + //
      "              }\n" + //
      "            }\n" + //
      "          }\n" + //
      "        }\n" + //
      "      }\n" + //
      "    },\n" + //
      "    \"tests\": {\n" + //
      "      \"type\": \"array\",\n" + //
      "      \"items\": {\n" + //
      "        \"type\": \"object\",\n" + //
      "        \"additionalProperties\": false,\n" + //
      "        \"properties\": {\n" + //
      "          \"name\": {\n" + //
      "            \"type\": \"string\",\n" + //
      "            \"description\": \"the test name\"\n" + //
      "          },\n" + //
      "          \"description\": {\n" + //
      "            \"type\": \"string\",\n" + //
      "            \"description\": \"the test description\"\n" + //
      "          },\n" + //
      "          \"tuple_file\": {\n" + //
      "            \"type\": \"string\",\n" + //
      "            \"description\": \"the tuple file with additional tuples for this test\"\n" + //
      "          },\n" + //
      "          \"tuples\": {\n" + //
      "            \"type\": \"array\",\n" + //
      "            \"description\": \"the tuples (takes precedence over tuples_file)\",\n" + //
      "            \"items\": {\n" + //
      "              \"type\": \"object\",\n" + //
      "              \"additionalProperties\": false,\n" + //
      "              \"required\": [\n" + //
      "                \"user\",\n" + //
      "                \"relation\",\n" + //
      "                \"object\"\n" + //
      "              ],\n" + //
      "              \"valid_tuple\": true,\n" + //
      "              \"properties\": {\n" + //
      "                \"user\": {\n" + //
      "                  \"type\": \"string\",\n" + //
      "                  \"format\": \"user\",\n" + //
      "                  \"description\": \"the user\"\n" + //
      "                },\n" + //
      "                \"relation\": {\n" + //
      "                  \"type\": \"string\",\n" + //
      "                  \"format\": \"relation\",\n" + //
      "                  \"description\": \"the relation\"\n" + //
      "                },\n" + //
      "                \"object\": {\n" + //
      "                  \"type\": \"string\",\n" + //
      "                  \"format\": \"object\",\n" + //
      "                  \"description\": \"the object\"\n" + //
      "                },\n" + //
      "                \"condition\": {\n" + //
      "                  \"type\": \"object\",\n" + //
      "                  \"additionalProperties\": false,\n" + //
      "                  \"required\": [\n" + //
      "                    \"name\"\n" + //
      "                  ],\n" + //
      "                  \"properties\": {\n" + //
      "                    \"name\": {\n" + //
      "                      \"type\": \"string\",\n" + //
      "                      \"format\": \"condition\"\n" + //
      "                    },\n" + //
      "                    \"context\": {\n" + //
      "                      \"type\": \"object\"\n" + //
      "                    }\n" + //
      "                  }\n" + //
      "                }\n" + //
      "              }\n" + //
      "            }\n" + //
      "          },\n" + //
      "          \"check\": {\n" + //
      "            \"type\": \"array\",\n" + //
      "            \"items\": {\n" + //
      "              \"type\": \"object\",\n" + //
      "              \"additionalProperties\": false,\n" + //
      "              \"required\": [\n" + //
      "                \"user\",\n" + //
      "                \"object\",\n" + //
      "                \"assertions\"\n" + //
      "              ],\n" + //
      "              \"properties\": {\n" + //
      "                \"user\": {\n" + //
      "                  \"type\": \"string\",\n" + //
      "                  \"format\": \"user\",\n" + //
      "                  \"description\": \"the user\"\n" + //
      "                },\n" + //
      "                \"object\": {\n" + //
      "                  \"type\": \"string\",\n" + //
      "                  \"format\": \"object\",\n" + //
      "                  \"description\": \"the object\"\n" + //
      "                },\n" + //
      "                \"assertions\": {\n" + //
      "                  \"type\": \"object\",\n" + //
      "                  \"patternProperties\": {\n" + //
      "                    \".*\": {\n" + //
      "                      \"type\": \"boolean\"\n" + //
      "                    }\n" + //
      "                  }\n" + //
      "                },\n" + //
      "                \"context\": {\n" + //
      "                  \"type\": \"object\"\n" + //
      "                }\n" + //
      "              }\n" + //
      "            }\n" + //
      "          },\n" + //
      "          \"list_objects\": {\n" + //
      "            \"type\": \"array\",\n" + //
      "            \"items\": {\n" + //
      "              \"type\": \"object\",\n" + //
      "              \"additionalProperties\": false,\n" + //
      "              \"required\": [\n" + //
      "                \"user\",\n" + //
      "                \"type\",\n" + //
      "                \"assertions\"\n" + //
      "              ],\n" + //
      "              \"properties\": {\n" + //
      "                \"user\": {\n" + //
      "                  \"type\": \"string\",\n" + //
      "                  \"format\": \"user\",\n" + //
      "                  \"description\": \"the user\"\n" + //
      "                },\n" + //
      "                \"relation\": {\n" + //
      "                  \"type\": \"string\",\n" + //
      "                  \"format\": \"relation\",\n" + //
      "                  \"description\": \"the relation\"\n" + //
      "                },\n" + //
      "                \"type\": {\n" + //
      "                  \"type\": \"string\",\n" + //
      "                  \"format\": \"type\",\n" + //
      "                  \"description\": \"the object type\"\n" + //
      "                },\n" + //
      "                \"assertions\": {\n" + //
      "                  \"type\": \"object\",\n" + //
      "                  \"patternProperties\": {\n" + //
      "                    \".*\": {\n" + //
      "                      \"type\": \"array\",\n" + //
      "                      \"items\": {\n" + //
      "                        \"type\": \"string\"\n" + //
      "                      }\n" + //
      "                    }\n" + //
      "                  }\n" + //
      "                },\n" + //
      "                \"context\": {\n" + //
      "                  \"type\": \"object\"\n" + //
      "                }\n" + //
      "              }\n" + //
      "            }\n" + //
      "          },\n" + //
      "          \"list_users\": {\n" + //
      "            \"type\": \"array\",\n" + //
      "            \"items\": {\n" + //
      "              \"type\": \"object\",\n" + //
      "              \"additionalProperties\": false,\n" + //
      "              \"required\": [\n" + //
      "                \"object\",\n" + //
      "                \"user_filter\",\n" + //
      "                \"assertions\"\n" + //
      "              ],\n" + //
      "              \"properties\": {\n" + //
      "                \"object\": {\n" + //
      "                  \"type\": \"string\",\n" + //
      "                  \"format\": \"object\"\n" + //
      "                },\n" + //
      "                \"user_filter\": {\n" + //
      "                  \"type\": \"array\",\n" + //
      "                  \"items\": {\n" + //
      "                    \"type\": \"object\",\n" + //
      "                    \"required\": [\n" + //
      "                      \"type\"\n" + //
      "                    ],\n" + //
      "                    \"properties\": {\n" + //
      "                      \"type\": {\n" + //
      "                        \"type\": \"string\"\n" + //
      "                      },\n" + //
      "                      \"relation\": {\n" + //
      "                        \"type\": \"string\"\n" + //
      "                      }\n" + //
      "                    }\n" + //
      "                  }\n" + //
      "                },\n" + //
      "                \"context\": {\n" + //
      "                  \"type\": \"object\"\n" + //
      "                },\n" + //
      "                \"assertions\": {\n" + //
      "                  \"type\": \"object\",\n" + //
      "                  \"patternProperties\": {\n" + //
      "                    \".*\": {\n" + //
      "                      \"type\": \"object\",\n" + //
      "                      \"additionalProperties\": false,\n" + //
      "                      \"properties\": {\n" + //
      "                        \"users\": {\n" + //
      "                          \"type\": \"array\",\n" + //
      "                          \"items\": {\n" + //
      "                            \"type\": \"string\",\n" + //
      "                            \"format\": \"user\"\n" + //
      "                          }\n" + //
      "                        }\n" + //
      "                      }\n" + //
      "                    }\n" + //
      "                  }\n" + //
      "                }\n" + //
      "              }\n" + //
      "            }\n" + //
      "          }\n" + //
      "        }\n" + //
      "      }\n" + //
      "    }\n" + //
      "  }\n" + //
      "}";

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

  public YamlStoreValidator() {
    JsonMetaSchema metaSchema = JsonMetaSchema.builder(JsonMetaSchema.getV202012())
        .format(new UserFormat())
        .format(new RelationFormat())
        .format(new ObjectFormat())
        .format(new ConditionFormat())
        .format(new TypeFormat())
        .build();
    JsonSchemaFactory factory = JsonSchemaFactory.getInstance(VersionFlag.V202012,
        builder -> builder.metaSchema(metaSchema).jsonNodeReader(JsonNodeReader.builder().locationAware().build()));
    SchemaValidatorsConfig config = SchemaValidatorsConfig.builder().build();
    this.schema = factory.getSchema(JSON_SCHEMA_STRING, InputFormat.JSON, config);
  }

  public List<StoreValidationSingleError> validate(final String store) {
    Set<ValidationMessage> messages = schema.validate(store, InputFormat.JSON, executionContext -> {
      executionContext.getExecutionConfig().setFormatAssertionsEnabled(true);
    });

    return messages.stream().map(m -> new StoreValidationSingleError(m.getMessage()))
        .collect(Collectors.toList());
  }
}
