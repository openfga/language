import { OpenFgaDslSyntaxMultipleError, parseDSL } from "../transformer/dsltojson";

/**
 * validateDSL - Validates model authored in FGA DSL syntax, returning all found errors
 * @param {string} dsl 
 * @returns {OpenFgaDslSyntaxMultipleError}
 */
export default function validateDsl(dsl: string): void {
    const { errorListener } = parseDSL(dsl);
  
    if (errorListener.errors.length) {
        throw new OpenFgaDslSyntaxMultipleError(errorListener.errors);
    }
}
  
  