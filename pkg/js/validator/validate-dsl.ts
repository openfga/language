import { OpenFgaDslSyntaxMultipleError, parseDSL } from "../transformer/dsltojson";

/**
 * validateDSL - Validates model authored in FGA DSL syntax, returning all found errors
 * @param {string} dsl 
 * @returns {OpenFgaDslSyntaxMultipleError}
 */
export default function validateDsl(dsl: string): OpenFgaDslSyntaxMultipleError {
    const { errorListener } = parseDSL(dsl);
  
    return new OpenFgaDslSyntaxMultipleError(errorListener.errors);
}
  
  