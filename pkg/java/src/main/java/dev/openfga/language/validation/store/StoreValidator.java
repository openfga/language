package dev.openfga.language.validation.store;

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

public class StoreValidator extends BaseJsonValidator {
    private static ErrorMessageType ERROR_MESSAGE_TYPE = new ErrorMessageType() {
        @Override
        public String getErrorCode() {
            return "valid_store";
        }
    };

    private final String value;

    public StoreValidator(
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

        // "'model' or 'model_file' must be presenet keys.", instancePath
        if (!node.has("model") && !node.has("model_file")) {
            return Collections.singleton(message()
                    .message("'model' or 'model_file' must be presenet keys.")
                    .instanceLocation(instanceLocation)
                    .instanceNode(node)
                    .build());
        }
        return validateTestTypes(node, instanceLocation);
    }

    private Set<ValidationMessage> validateTestTypes(JsonNode node, JsonNodePath instanceLocation) {

        return null;
    }
}

/*

function validateTestTypes(store: Store, model: AuthorizationModel, instancePath: string): boolean {
  const errors = [];

  // Collect params for validity checking
  const params: string[] = [];
  for (const condition in model.conditions) {
    for (const param in model.conditions[condition].parameters) {
      params.push(param);
    }
  }

  for (const singleTest in store.tests) {
    // Collect valid tuples
    const tuples = [];
    if (store.tuples && store.tuples.length) {
      tuples.push(...store.tuples);
    }

    const test = store.tests[singleTest];
    if (test.tuples && test.tuples.length) {
      tuples.push(...test.tuples);
    }

    // Validate check
    if (test.check) {
      for (const [testNumber, singleCheckTest] of Object.entries(test.check)) {
        if (!singleCheckTest.user || !singleCheckTest.object) {
          return false;
        }
        errors.push(
          ...validateCheck(
            model,
            singleCheckTest,
            tuples,
            params,
            instancePath + `/tests/${singleTest}/check/${testNumber}`,
          ),
        );
      }
    }

    // Validate list objects
    if (test.list_objects) {
      for (const [testNumber, singleListObjectTest] of Object.entries(test.list_objects)) {
        if (!singleListObjectTest.user || !singleListObjectTest.type) {
          return false;
        }
        errors.push(
          ...validateListObject(
            model,
            singleListObjectTest,
            tuples,
            params,
            instancePath + `/tests/${singleTest}/list_objects/${testNumber}`,
          ),
        );
      }
    }

    // Validate list users
    if (test.list_users) {
      for (const [testNumber, singleListUsersTest] of Object.entries(test.list_users)) {
        if (!singleListUsersTest.object || !singleListUsersTest.user_filter) {
          return false;
        }
        errors.push(
          ...validateListUsers(
            model,
            singleListUsersTest,
            tuples,
            params,
            instancePath + `/tests/${singleTest}/list_users/${testNumber}`,
          ),
        );
      }
    }
  }

  if (errors.length) {
    validateStore.errors?.push(...errors);
    return false;
  }
  return true;
}
 */
