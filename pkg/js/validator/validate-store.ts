import {
  AuthorizationModel,
  CheckRequestTupleKey,
  Condition,
  ListObjectsRequest,
  ListUsersRequest,
  RelationReference,
  TupleKey,
  TypeDefinition,
} from "@openfga/sdk";
import Ajv, { Schema, ValidateFunction, SchemaValidateFunction } from "ajv";
import { Validator } from "./validate-rules";

export function isStringValue(str: unknown) {
  return typeof str == "string" || String;
}

type Store = {
  name: string;
  model_file?: string;
  model?: string;
  tuple_file?: string;
  tuples?: TupleKey[];
  tests: Test[];
};

type Test = {
  tuple_file?: string;
  tuples: TupleKey[];
  check: CheckTest[];
  list_objects: ListObjectTest[];
  list_users: ListUsersTest[];
};

type CheckTest = Omit<CheckRequestTupleKey, "relation"> & {
  assertions: Record<string, boolean>;
  // eslint-disable-next-line  @typescript-eslint/no-explicit-any
  context: Record<string, any>;
};

type ListObjectTest = Omit<ListObjectsRequest, "relation"> & {
  // eslint-disable-next-line  @typescript-eslint/no-explicit-any
  assertions: Record<string, any>;
};

type ListUsersTest = Omit<ListUsersRequest, "relation" | "user_filters"> & {
  object: string;
  // eslint-disable-next-line  @typescript-eslint/no-explicit-any
  user_filter: Record<string, any>;
  // eslint-disable-next-line  @typescript-eslint/no-explicit-any
  assertions: Record<string, any>;
};

type BaseError = { keyword: string; message: string; instancePath: string };

// Errors for tuples validation
const invalidTuple = (message: string, instancePath: string): BaseError => {
  return {
    keyword: "valid_tuple",
    message,
    instancePath,
  };
};

const invalidType = (user: string, types: string[], instancePath: string) => {
  return invalidTuple(`invalid type '${user}'. Valid types are [${types}]`, instancePath);
};

const relationMustExistOnType = (relation: string, type: string, instancePath: string) => {
  return invalidTuple(`relation '${relation}' is not a relation on type '${type}'.`, instancePath);
};

const userNotTypeRestriction = (user: string, tuple: TupleKey, instancePath: string) => {
  return invalidTuple(
    `'${user}' is not a type restriction on relation '${tuple.relation}' of type '${tuple.object.split(":")[0]}'.`,
    instancePath + "/user",
  );
};

const conditionDoesntExist = (tuple: TupleKey, instancePath: string) => {
  return invalidTuple(`condition '${tuple.condition?.name}' is not defined.`, instancePath + "/condition/name");
};

const notAParameter = (param: string, tuple: TupleKey, instancePath: string) => {
  return invalidTuple(
    `'${param}' is not a parameter on condition '${tuple.condition?.name}'.`,
    instancePath + `/condition/context/${param}`,
  );
};

// Errors for store validation
const invalidStore = (message: string, instancePath: string) => {
  return {
    keyword: "valid_store",
    message,
    instancePath,
  };
};

const modelOrModelFile = (instancePath: string) => {
  return invalidStore("'model' or 'model_file' must be presenet keys.", instancePath);
};

const invalidUserField = (instancePath: string) => {
  return invalidStore("`user` field must be one of [<type>:<id>, <type>:* or <type>:<id>#<relation>]", instancePath);
};

const invalidTypeUser = (type: string, types: string[], instancePath: string) => {
  return invalidStore(`invalid type '${type}'. Valid types are [${types}]`, instancePath);
};

const invalidRelationUser = (relation: string, relations: string[], instancePath: string) => {
  return invalidStore(`invalid relation '${relation}'. Valid relations are [${relations}]`, instancePath);
};

const nonMatchingRelationType = (relation: string, user: string, values: string[], instancePath: string) => {
  if (values.length) {
    return invalidStore(
      `\`${relation}\` is not a relation on \`${user}\`, and does not exist in model ` +
        `- valid relations are [${values}].`,
      instancePath,
    );
  }
  return invalidStore(`\`${relation}\` is not a relation on \`${user}\`, and does not exist in model.`, instancePath);
};

const invalidAssertion = (assertion: string, object: string, instancePath: string) => {
  return invalidStore(`\`${assertion}\` is not a relationship for type \`${object}\`.`, instancePath);
};

const unidentifiedTestParam = (testParam: string, instancePath: string) => {
  return invalidStore(
    `\`${testParam}\` is not a recognized paramaeter for any condition defined in the model.`,
    instancePath,
  );
};

const undefinedTypeTuple = (user: string, instancePath: string) => {
  return {
    keyword: "valid_store_warning",
    message: `${user} does not match any existing tuples; the check is still valid ` +
      "- but double check to ensure this is intended.",
    instancePath,
  };
};

// Helpers

function isValidTuple(tuple: TupleKey): boolean {
  if (
    !tuple.user ||
    !isStringValue(tuple.user) ||
    !tuple.object ||
    !isStringValue(tuple.object) ||
    !tuple.relation ||
    !isStringValue(tuple.relation)
  ) {
    return false;
  }

  return true;
}

function mapTuple(tuple: TupleKey): RelationReference {
  return {
    type: tuple.user.split(":")[0],
    relation: tuple.user.includes("#") ? tuple.user.split("#")[1] : undefined,
    wildcard: tuple.user.includes(":*") ? {} : undefined,
    condition: tuple.condition?.name,
  };
}

function parseObjectString(object: string): { type: string; id: string } {
  const result = object.split(":");
  return { type: result[0], id: result[1] };
}

function parseUserString(user: string): {
  type: string;
  id: string | undefined;
  relation: string | undefined;
  wildcard: boolean;
} {
  const result = user.split(":");
  const type = result[0];

  if (result[1] === "*") {
    return { type, id: undefined, relation: undefined, wildcard: true };
  }

  if (result[1].includes("#")) {
    const [id, relation] = result[1].split("#");
    return { type, id, relation, wildcard: false };
  }

  return { type, id: result[1], relation: undefined, wildcard: false };
}

// Validation for Types

function validateTypes(tuple: TupleKey, types: string[], instancePath: string): boolean {
  const errors = [];

  const userType = parseUserString(tuple.user).type;

  // Ensure valid type of user
  if (!types.includes(userType)) {
    errors.push(invalidType(userType, types, instancePath + "/user"));
  }

  const objectType = parseObjectString(tuple.object).type;

  // Ensure valid type of object
  if (!types.includes(objectType)) {
    errors.push(invalidType(objectType, types, instancePath + "/object"));
  }

  // Report all errors
  if (errors.length) {
    validateTuple.errors?.push(...errors);
    return false;
  }
  return true;
}

// Validate Relation
function validateRelation(tuple: TupleKey, typeDefs: TypeDefinition[], instancePath: string): boolean {
  const errors = [];

  // Check if relation exists on given type
  let doesExistOnType = false;
  if (tuple.user.includes("#")) {
    const userResult = parseUserString(tuple.user);
    for (const typeDef of typeDefs) {
      if (
        typeDef &&
        typeDef.type === userResult.type &&
        userResult.relation &&
        typeDef.relations?.[userResult.relation]
      ) {
        doesExistOnType = true;
        break;
      }
    }

    if (!doesExistOnType && userResult.relation) {
      errors.push(relationMustExistOnType(userResult.relation, userResult.type, instancePath + "/user"));
    }
  }

  // Check if relation exists on given object
  const objectType = parseObjectString(tuple.object).type;
  let doesExistOnObject = false;
  for (const typeDef of typeDefs) {
    if (typeDef && typeDef.type === objectType && typeDef.relations?.[tuple.relation]) {
      doesExistOnObject = true;
      break;
    }
  }
  if (!doesExistOnObject) {
    errors.push(relationMustExistOnType(tuple.relation, objectType, instancePath + "/relation"));
  }

  if (errors.length) {
    validateTuple.errors?.push(...errors);
    return false;
  }
  return true;
}

function getRelationReferenceString(relationReference: RelationReference) {
  let relationReferenceString = relationReference.type;
  if (relationReference.wildcard) {
    relationReferenceString += ":*";
  } else if (relationReference.relation) {
    relationReferenceString += `#${relationReference.relation}`;
  }

  if (relationReference.condition) {
    relationReferenceString += `' with '${relationReference.condition}`;
  }

  return relationReferenceString;
}

function validateTypeRestrictions(
  tuple: TupleKey,
  typeDefs: TypeDefinition[],
  conditions: Record<string, Condition> | undefined,
  instancePath: string,
): boolean {
  validateTuple.errors = validateTuple.errors || [];

  const mappedTuple = mapTuple(tuple);
  const type = parseObjectString(tuple.object).type;
  const foundType = typeDefs.filter((t) => t.type === type)[0];
  const userTypes = foundType?.metadata?.relations?.[tuple.relation].directly_related_user_types;

  if (
    !userTypes?.filter(
      (userType) =>
        userType.type === mappedTuple.type && // type matches
        !!mappedTuple.wildcard === !!userType.wildcard && // and the wildcard matches (either both true or both false)
        userType.relation === mappedTuple.relation && // and the relation matches
        userType.condition === mappedTuple.condition, // and the condition matches
    ).length
  ) {
    validateTuple.errors.push(userNotTypeRestriction(getRelationReferenceString(mappedTuple), tuple, instancePath));
    return false;
  }

  if (mappedTuple.condition && conditions) {
    // Check parameters for matching condition
    return validateConditionParams(tuple, mappedTuple.condition, conditions, instancePath);
  }
  return true;
}

function validateConditionExists(
  tuple: TupleKey,
  conditions: Record<string, Condition> | undefined,
  instancePath: string,
): boolean {
  if (!tuple.condition) {
    return true;
  }

  validateTuple.errors = validateTuple.errors || [];

  // Condition on tuple not found
  if (!conditions || !conditions[tuple.condition.name]) {
    validateTuple.errors.push(conditionDoesntExist(tuple, instancePath));
    return false;
  }
  return true;
}

function validateConditionParams(
  tuple: TupleKey,
  condition: string,
  conditions: Record<string, Condition>,
  instancePath: string,
): boolean {
  validateTuple.errors = validateTuple.errors || [];

  if (tuple.condition && tuple.condition.context && conditions[condition].parameters) {
    for (const param of Object.keys(tuple.condition.context)) {
      if (!conditions[condition].parameters![param]) {
        validateTuple.errors.push(notAParameter(param, tuple, instancePath));
        return false;
      }
    }
  }
  return true;
}

// Validation for tuples
const validateTuple: SchemaValidateFunction = function (
  this: { jsonModel: AuthorizationModel },
  tuple: TupleKey,
  ctx: { instancePath: string },
): boolean {
  validateTuple.errors = validateTuple.errors || [];

  if (!isValidTuple(tuple)) {
    return false;
  }

  if (!this.jsonModel) {
    return false;
  }

  const jsonModel: AuthorizationModel = this.jsonModel;

  // Validate
  const types = jsonModel.type_definitions.map((d) => d.type);
  return (
    validateTypes(tuple, types, ctx.instancePath) &&
    validateRelation(tuple, jsonModel.type_definitions, ctx.instancePath) &&
    validateConditionExists(tuple, jsonModel.conditions, ctx.instancePath) &&
    validateTypeRestrictions(tuple, jsonModel.type_definitions, jsonModel.conditions, ctx.instancePath)
  );
};

// Validate the type field for check and list_objects
function validateUserField(model: AuthorizationModel, types: string[], userField: string, instancePath: string) {
  const errors = [];
  const foundType = parseUserString(userField).type;

  // Ensure valid type
  if (!types.includes(foundType)) {
    errors.push(invalidUserField(instancePath + "/user"));
    errors.push(invalidTypeUser(foundType, types, instancePath + "/user"));
  }

  // Valdidate check user
  if (userField.includes("#")) {
    const [type, relation] = userField.split("#");

    const userRelations = model.type_definitions
      .filter((typeDef) => typeDef.type === type)
      .flatMap((typeDef) => {
        const relationArray: string[] = [];
        for (const rel in typeDef.relations) {
          relationArray.push(type + "#" + rel);
        }
        return relationArray;
      });

    if (!userRelations.includes(userField)) {
      errors.push(
        nonMatchingRelationType(
          relation,
          type,
          userRelations.map((rel) => rel.split("#")[1]),
          instancePath + "/user",
        ),
      );
    }
  }
  return errors;
}

function validateAssertionField(
  model: AuthorizationModel,
  typeField: string,
  // eslint-disable-next-line  @typescript-eslint/no-explicit-any
  assertions: Record<string, any>,
  instancePath: string,
) {
  const errors = [];

  // Validate assertions exist as relations
  const typesRelations = model.type_definitions
    .filter((tuple) => tuple.type === typeField)
    .map((tuple) => tuple.relations);
  for (const assertion in assertions) {
    for (const relation in typesRelations) {
      if (!typesRelations[relation]?.[assertion]) {
        errors.push(invalidAssertion(assertion, typeField, instancePath + `/assertions/${assertion}`));
      }
    }
  }

  return errors;
}

// Validate Check Tuple
function validateCheck(
  model: AuthorizationModel,
  checkTest: CheckTest,
  tuples: TupleKey[],
  params: string[],
  instancePath: string,
) {
  const userErrors = [];

  const types = model.type_definitions.map((d) => d.type);

  if (checkTest.user && isStringValue(checkTest.user)) {
    const checkUser = checkTest.user;
    userErrors.push(...validateUserField(model, types, checkUser, instancePath));
    if (!userErrors.length) {
      if (!tuples.map((tuple) => tuple.user).filter((user) => user === checkUser).length) {
        userErrors.push(undefinedTypeTuple(checkUser, instancePath + "/user"));
      }
    }
  }

  const objectErrors = [];

  if (checkTest.object && isStringValue(checkTest.object)) {
    const checkObject = checkTest.object;
    const objectType = parseObjectString(checkObject).type;

    // Ensure valid type of object
    if (!types.includes(objectType)) {
      objectErrors.push(invalidTypeUser(objectType, types, instancePath + "/object"));
    }

    if (!objectErrors.length) {
      if (!tuples.map((tuple) => tuple.object).filter((object) => object === checkObject).length) {
        objectErrors.push(undefinedTypeTuple(checkObject, instancePath + "/object"));
      }
    }

    objectErrors.push(...validateAssertionField(model, objectType, checkTest.assertions, instancePath));
  }

  const context = checkTest.context;
  for (const testParam in context) {
    if (!params.includes(testParam)) {
      objectErrors.push(unidentifiedTestParam(testParam, instancePath + `/context/${testParam}`));
    }
  }

  return [...userErrors, ...objectErrors];
}

// Validate List Object
function validateListObject(
  model: AuthorizationModel,
  listObjects: ListObjectTest,
  tuples: TupleKey[],
  params: string[],
  instancePath: string,
) {
  const errors = [];

  const types = model.type_definitions.map((d) => d.type);

  if (listObjects.user && isStringValue(listObjects.user)) {
    const listUser = listObjects.user;

    errors.push(...validateUserField(model, types, listUser, instancePath));

    if (!errors.length) {
      if (!tuples.map((tuple) => tuple.user).filter((user) => user === listUser).length) {
        errors.push(undefinedTypeTuple(listUser, instancePath + "/user"));
      }
    }
  }

  if (listObjects.type && isStringValue(listObjects.type)) {
    const listType = listObjects.type;

    // Ensure valid type of object
    if (!types.includes(listType)) {
      errors.push(invalidTypeUser(listType, types, instancePath + "/type"));
    }

    // Validate assertions exist as relations
    errors.push(...validateAssertionField(model, listType, listObjects.assertions, instancePath));
  }

  const context = listObjects.context;
  for (const testParam in context) {
    if (!params.includes(testParam)) {
      errors.push(unidentifiedTestParam(testParam, instancePath + `/context/${testParam}`));
    }
  }
  return errors;
}

// Validate List Users
function validateListUsers(
  model: AuthorizationModel,
  listUsers: ListUsersTest,
  tuples: TupleKey[],
  params: string[],
  instancePath: string,
) {
  const errors = [];

  const types = model.type_definitions.map((d) => d.type);

  if (listUsers && isStringValue(listUsers.object)) {
    const listUserObj = listUsers.object;

    const objectType = parseObjectString(listUserObj).type;

    // Ensure valid type of object
    if (!types.includes(objectType)) {
      errors.push(invalidTypeUser(objectType, types, instancePath + "/object"));
    }

    if (!errors.length) {
      if (!tuples.map((tuple) => tuple.object).filter((object) => object === listUserObj).length) {
        errors.push(undefinedTypeTuple(listUserObj, instancePath + "/object"));
      }
    }
  }

  // Check user filter
  if (listUsers.user_filter) {
    for (const index in listUsers.user_filter) {
      const listType = listUsers.user_filter[index].type;

      if (listType && isStringValue(listType)) {
        // Ensure valid type of object
        if (!types.includes(listType)) {
          errors.push(invalidTypeUser(listType, types, instancePath + `/user_filter/${index}/type`));
        }
      }

      // Check relations if present
      const relation = listUsers.user_filter[index].relation;
      if (relation && isStringValue(relation)) {
        const typeRelations = model.type_definitions.filter((rel) => rel.type === listType).map((rel) => rel.relations);

        if (typeRelations.length && typeRelations[0] && Object.keys(typeRelations[0]).length) {
          if (!Object.keys(typeRelations[0]).includes(relation)) {
            errors.push(
              invalidRelationUser(
                relation,
                Object.keys(typeRelations[0]),
                instancePath + `/user_filter/${index}/relation`,
              ),
            );
          }
        } else {
          errors.push(relationMustExistOnType(relation, listType, instancePath + `/user_filter/${index}/relation`));
        }
      }
    }

    // Check assertions
    if (listUsers.assertions) {
      for (const assertion of Object.keys(listUsers.assertions)) {
        if (listUsers.assertions[assertion].users) {
          for (const user of listUsers.assertions[assertion].users) {
            if (!tuples.some((tuple) => tuple.user === user)) {
              errors.push(undefinedTypeTuple(user, instancePath + `/assertions/${assertion}/users`));
            }
          }
        }
      }
    }
  }

  // Check context params
  if (listUsers.context) {
    for (const testParam in listUsers.context) {
      if (!params.includes(testParam)) {
        errors.push(unidentifiedTestParam(testParam, instancePath + `/context/${testParam}`));
      }
    }
  }

  return errors;
}

// Validation for types in check
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

const validateStore: SchemaValidateFunction = function (
  this: { jsonModel: AuthorizationModel },
  store: Store,
  ctx: { instancePath: string },
): boolean {
  validateStore.errors = validateStore.errors || [];

  // Require model or model_file
  if (!store.model && !store.model_file) {
    validateStore.errors.push(modelOrModelFile(ctx.instancePath));
    return false;
  }

  return validateTestTypes(store, this.jsonModel, ctx.instancePath);
};

// YAML validation using ajv
export function YamlTuplesValidator(): ValidateFunction {
  return new Ajv({
    allErrors: true,
    verbose: true,
    passContext: true,
    $data: true,
  })
    .addFormat("user", {
      validate: Validator.user,
    })
    .addFormat("relation", {
      validate: Validator.relation,
    })
    .addFormat("object", {
      validate: Validator.object,
    })
    .addFormat("condition", {
      validate: Validator.relationshipCondition,
    })
    .addKeyword({
      keyword: "valid_tuple",
      type: "object",
      schema: false,
      errors: true,
      validate: validateTuple,
    })
    .compile(OPENFGA_TUPLES_SCHEMA);
}

// YAML tuple validation
const OPENFGA_TUPLES_SCHEMA: Schema = {
  type: "array",
  description: "the tuples (takes precedence over tuples_file)",
  items: {
    type: "object",
    additionalProperties: false,
    required: ["user", "relation", "object"],
    valid_tuple: true,
    properties: {
      user: {
        type: "string",
        format: "user",
        description: "the user",
      },
      relation: {
        type: "string",
        format: "relation",
        description: "the relation",
      },
      object: {
        type: "string",
        format: "object",
        description: "the object",
      },
      condition: {
        type: "object",
        additionalProperties: false,
        required: ["name"],
        properties: {
          name: {
            type: "string",
            format: "condition",
          },
          context: {
            type: "object",
          },
        },
      },
    },
  },
};

// YAML validation using ajv
export function YamlStoreValidator(): ValidateFunction {
  return new Ajv({
    allErrors: true,
    verbose: true,
    passContext: true,
    $data: true,
  })
    .addFormat("user", {
      validate: Validator.user,
    })
    .addFormat("relation", {
      validate: Validator.relation,
    })
    .addFormat("object", {
      validate: Validator.object,
    })
    .addFormat("condition", {
      validate: Validator.relationshipCondition,
    })
    .addFormat("type", {
      validate: Validator.type,
    })
    .addKeyword({
      keyword: "valid_tuple",
      type: "object",
      schema: false,
      errors: true,
      validate: validateTuple,
    })
    .addKeyword({
      keyword: "valid_store",
      type: "object",
      schema: false,
      errors: true,
      validate: validateStore,
    })
    .compile(OPENFGA_YAML_SCHEMA);
}

// YAML Schema
const OPENFGA_YAML_SCHEMA: Schema = {
  type: "object",
  required: ["tests"],
  additionalProperties: false,
  valid_store: true,
  properties: {
    name: {
      type: "string",
      description: "the store name",
    },
    model_file: {
      type: "string",
      description: "the authorization model file path",
    },
    model: {
      type: "string",
      description: "the authorization model (takes precedence over model_file)",
    },
    tuple_file: {
      type: "string",
      description: "the tuple file path",
    },
    tuples: OPENFGA_TUPLES_SCHEMA,
    tests: {
      type: "array",
      items: {
        type: "object",
        additionalProperties: false,
        properties: {
          name: {
            type: "string",
            description: "the test name",
          },
          description: {
            type: "string",
            description: "the test description",
          },
          tuple_file: {
            type: "string",
            description: "the tuple file with additional tuples for this test",
          },
          tuples: OPENFGA_TUPLES_SCHEMA,
          check: {
            type: "array",
            items: {
              type: "object",
              additionalProperties: false,
              required: ["user", "object", "assertions"],
              properties: {
                user: {
                  type: "string",
                  format: "user",
                  description: "the user",
                },
                object: {
                  type: "string",
                  format: "object",
                  description: "the object",
                },
                assertions: {
                  type: "object",
                  patternProperties: {
                    ".*": {
                      type: "boolean",
                    },
                  },
                },
                context: {
                  type: "object",
                },
              },
            },
          },
          list_objects: {
            type: "array",
            items: {
              type: "object",
              additionalProperties: false,
              required: ["user", "type", "assertions"],
              properties: {
                user: {
                  type: "string",
                  format: "user",
                  description: "the user",
                },
                relation: {
                  type: "string",
                  format: "relation",
                  description: "the relation",
                },
                type: {
                  type: "string",
                  format: "type",
                  description: "the object type",
                },
                assertions: {
                  type: "object",
                  patternProperties: {
                    ".*": {
                      type: "array",
                      items: {
                        type: "string",
                      },
                    },
                  },
                },
                context: {
                  type: "object",
                },
              },
            },
          },
          list_users: {
            type: "array",
            items: {
              type: "object",
              additionalProperties: false,
              required: ["object", "user_filter", "assertions"],
              properties: {
                object: {
                  type: "string",
                  format: "object",
                },
                user_filter: {
                  type: "array",
                  items: {
                    type: "object",
                    required: ["type"],
                    properties: {
                      type: {
                        type: "string",
                      },
                      relation: {
                        type: "string",
                      },
                    },
                  },
                },
                context: {
                  type: "object",
                },
                assertions: {
                  type: "object",
                  patternProperties: {
                    ".*": {
                      type: "object",
                      additionalProperties: false,
                      properties: {
                        users: {
                          type: "array",
                          items: {
                            type: "string",
                            format: "user",
                          },
                        },
                      },
                    },
                  },
                },
              },
            },
          },
        },
      },
    },
  },
};
