import { string } from "yaml/dist/schema/common/string";

export const Rules = {
  type: "[^:#@\\s]{1,254}",
  relation: "[^:#@\\s]{1,50}",
  condition: "[^\\s]{2,256}",
  object: "[^\\s]{2,256}",
  id: "[^#:\\s]+",
};

export const Validate = {
  // An Object is composed of a type and identifier (e.g. 'document:1')
  object: (object: string): boolean => {
    return validateFieldValue(`^${Rules.object}:${Rules.id}$`, object);
  },
  // Relation reference
  relation: (relation: string): boolean => {
    return validateFieldValue(`^${Rules.relation}$`, relation);
  },
  // User is composed of type and identifier (e.g. 'group:engineering')
  userSet: (user: string): boolean => {
    return validateFieldValue(`^${Rules.type}:${Rules.id}#${Rules.relation}$`, user);
  },
  // User is composed of type and identifier (e.g. 'group:engineering')
  userObject: (userObject: string): boolean => {
    return validateFieldValue(`^${Rules.type}:${Rules.id}$`, userObject);
  },
  // User is composed of type, identifier and relation (e.g. 'group:engineering#member')
  userWildcard: (userWildcard: string): boolean => {
    return validateFieldValue(`^${Rules.type}:\\*$`, userWildcard);
  },
  // Is either a userset, userobject or a user wildcard
  user: (user: string): boolean => {
    return Validate.userSet(user) || Validate.userObject(user) || Validate.userWildcard(user);
  },
  // Condition name
  relationshipCondition: (condition: string) => {
    return validateFieldValue(`^${Rules.condition}$`, condition);
  },
  type: (type: string): boolean => {
    return validateFieldValue(`^${Rules.type}$`, type);
  },
};

const validateFieldValue = (rule: string, value: string): boolean => {
  return new RegExp(rule).test(value);
};
