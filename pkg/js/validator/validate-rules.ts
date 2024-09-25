export const Rules = {
  type: "[^:#@\\*\\s]{1,254}",
  relation: "[^:#@\\*\\s]{1,50}",
  condition: "[^\\*\\s]{1,50}",
  id: "(?![#:\\s*])[a-zA-Z0-9_|*@.+]+",
  object: "[^\\s]{2,256}",
};

export const Validator = {
  // An Object is composed of a type and identifier (e.g. 'document:1')
  object: (object: string): boolean => {
    return validateFieldValue(`^${Rules.type}:${Rules.id}$`, object) && validateFieldValue(`^${Rules.object}$`, object);
  },
  // Relation reference
  relation: (relation: string): boolean => {
    return validateFieldValue(`^${Rules.relation}$`, relation);
  },
  // User is composed of type, identifier and relation (e.g. 'group:engineering#member')
  userSet: (user: string): boolean => {
    return validateFieldValue(`^${Rules.type}:${Rules.id}#${Rules.relation}$`, user);
  },
  // User is composed of type and identifier (e.g. 'group:engineering')
  userObject: (userObject: string): boolean => {
    return (
      validateFieldValue(`^${Rules.type}:${Rules.id}$`, userObject) &&
      validateFieldValue(`^${Rules.object}$`, userObject)
    );
  },
  // User is composed of type, and a wildcard (e.g. 'group:*')
  userWildcard: (userWildcard: string): boolean => {
    return validateFieldValue(`^${Rules.type}:\\*$`, userWildcard);
  },
  // Is either a userset, userobject or a user wildcard
  user: (user: string): boolean => {
    return Validator.userSet(user) || Validator.userObject(user) || Validator.userWildcard(user);
  },
  // Condition name
  relationshipCondition: (condition: string) => {
    return validateFieldValue(`^${Rules.condition}$`, condition);
  },
  // Type name
  type: (type: string): boolean => {
    return validateFieldValue(`^${Rules.type}$`, type);
  },
};

const validateFieldValue = (rule: string, value: string): boolean => {
  return new RegExp(rule).test(value);
};
