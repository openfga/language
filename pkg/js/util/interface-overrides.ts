// Only needed until the JS SDK is updated
import type {
  AuthorizationModel as OriginalAuthorizationModel,
  RelationReference as OriginalRelationReference,
} from "@openfga/sdk";

export interface ConditionParameterDefinition {
  type_name: string;
  generic_types?: ConditionParameterDefinition[];
}

export interface Condition {
  name: string;
  expression: string;
  parameters: Record<string, ConditionParameterDefinition>;
}

export type RelationReference = OriginalRelationReference & { condition?: string };
export type AuthorizationModel = OriginalAuthorizationModel & { conditions?: Record<string, Condition> };
