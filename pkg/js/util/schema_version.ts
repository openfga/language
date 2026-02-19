// Instead of importing from @openfga/sdk and creating a runtime dependency, define locally.
function assertNever(x: never): never {
  throw new Error(`Unexpected value: ${x}`);
}

/**
 * SchemaVersion represents the version of the schema used in the authorization model.
 * Currently, OpenFGA supports schema versions 1.1, and 1.2, 1.0 has been deprecated and is no longer supported.
 */
export enum SchemaVersion {
  V1_0 = "1.0",
  V1_1 = "1.1",
  V1_2 = "1.2",
}

/**
 * isSchemaVersionSupported checks if the provided schema version is supported by OpenFGA.
 * @param {SchemaVersion} version - The schema version to check.
 * @return {boolean} - Returns true if the schema version is supported, false otherwise.
 */
export function isSchemaVersionSupported(version: SchemaVersion): boolean {
  switch (version) {
    case SchemaVersion.V1_1:
    case SchemaVersion.V1_2:
      return true;
    case SchemaVersion.V1_0:
      return false;
    default:
      assertNever(version);
  }
}

/**
 * checkSchemaVersionSupportsModules checks if the provided schema version supports modules.
 * @param {SchemaVersion} version - The schema version to check.
 * @returns {boolean} - Returns true if the schema version supports modules, false otherwise.
 */
export function checkSchemaVersionSupportsModules(version: SchemaVersion): boolean {
  switch (version) {
    case SchemaVersion.V1_0:
    case SchemaVersion.V1_1:
      return false;
    case SchemaVersion.V1_2:
      return true;
    default:
      assertNever(version);
  }
}

/**
 * getSchemaVersionFromString converts a string to a SchemaVersion enum value, or throws an error if the string is not a valid schema version.
 * @param {string} version - The string representation of the schema version.
 * @returns {SchemaVersion} - The corresponding SchemaVersion enum value.
 * @throws {Error} - Throws an error if the provided string does not correspond to a valid schema version.
 */
export function getSchemaVersionFromString(version: string): SchemaVersion {
  switch (version) {
    case "1.0":
      return SchemaVersion.V1_0;
    case "1.1":
      return SchemaVersion.V1_1;
    case "1.2":
      return SchemaVersion.V1_2;
    default:
      throw new Error(`Unsupported schema version: ${version}`);
  }
}
