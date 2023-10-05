import OpenFGAParser from '../gen/OpenFGAParser';

  // Lookup and return the corresponding literal from the parser, without quotes
function getSymbol(symbol: number): string {
  return OpenFGAParser.literalNames[symbol]!.replace(/'/g, '');
}

export type DocumentationMap = Partial<Record<string, { summary: string; link?: string }>>;

export const defaultDocumentationMap: DocumentationMap = {
  [getSymbol(OpenFGAParser.TYPE)]: {
    summary: `A type or grouping of objects that have similar characteristics. For example:
- workspace
- repository
- organization
- document`,
    link: "https://openfga.dev/docs/concepts#what-is-a-type",
  },
  [getSymbol(OpenFGAParser.RELATIONS)]: {
    summary:
      "A **relation** defines the possible relationship between an [object](https://openfga.dev/docs/concepts#what-is-an-object) and a [user](https://openfga.dev/docs/concepts#what-is-a-user).",
    link: "https://openfga.dev/docs/concepts#what-is-a-relation",
  },
  [getSymbol(OpenFGAParser.DEFINE)]: {
    summary:
      "A **relation** defines the possible relationship between an [object](https://openfga.dev/docs/concepts#what-is-an-object) and a [user](https://openfga.dev/docs/concepts#what-is-a-user).",
    link: "https://openfga.dev/docs/concepts#what-is-a-relation",
  },
  [getSymbol(OpenFGAParser.AND)]: {
    summary:
      "The intersection operator used to indicate that a relationship exists if the user is in all the sets of users.",
    link: "https://openfga.dev/docs/configuration-language#the-intersection-operator",
  },
  [getSymbol(OpenFGAParser.OR)]: {
    summary:
      "The union operator is used to indicate that a relationship exists if the user is in any of the sets of users",
    link: "https://openfga.dev/docs/configuration-language#the-union-operator",
  },
  [getSymbol(OpenFGAParser.BUT_NOT)]: {
    summary:
      "The exclusion operator is used to indicate that a relationship exists if the user is in the base userset, but not in the excluded userset.",
    link: "https://openfga.dev/docs/configuration-language#the-exclusion-operator",
  },
  [getSymbol(OpenFGAParser.FROM)]: {
    summary: "Allows referencing relations on related objects.",
    link: "https://openfga.dev/docs/configuration-language#referencing-relations-on-related-objects",
  },
  [getSymbol(OpenFGAParser.SCHEMA)]: {
    summary:
      "Defines the schema version to be used, with currently only support for '1.1'. Note that the 1.0 schema is deprecated.",
    link: "https://openfga.dev/docs/modeling/migrating/migrating-schema-1-1",
  },
};