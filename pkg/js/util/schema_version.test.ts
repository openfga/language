import { describe, expect, it } from "@jest/globals";
import {
  SchemaVersion,
  isSchemaVersionSupported,
  checkSchemaVersionSupportsModules,
  getSchemaVersionFromString,
} from "./schema_version";

describe("SchemaVersion", () => {
  describe("SchemaVersion enum", () => {
    const testCases = [
      { enumKey: SchemaVersion.V1_0, expectedValue: "1.0" },
      { enumKey: SchemaVersion.V1_1, expectedValue: "1.1" },
      { enumKey: SchemaVersion.V1_2, expectedValue: "1.2" },
    ] as const;

    testCases.forEach(({ enumKey, expectedValue }) => {
      it(`should have ${enumKey} with value '${expectedValue}'`, () => {
        expect(enumKey.toString()).toBe(expectedValue);
      });
    });
  });

  describe("isSchemaVersionSupported", () => {
    const testCases = [
      { version: SchemaVersion.V1_0, expected: false, description: "1.0" },
      { version: SchemaVersion.V1_1, expected: true, description: "1.1" },
      { version: SchemaVersion.V1_2, expected: true, description: "1.2" },
    ];

    testCases.forEach(({ version, expected, description }) => {
      it(`should return ${expected} for schema version ${description}`, () => {
        expect(isSchemaVersionSupported(version)).toBe(expected);
      });
    });
  });

  describe("checkSchemaVersionSupportsModules", () => {
    const testCases = [
      { version: SchemaVersion.V1_0, expected: false, description: "1.0" },
      { version: SchemaVersion.V1_1, expected: false, description: "1.1" },
      { version: SchemaVersion.V1_2, expected: true, description: "1.2" },
    ];

    testCases.forEach(({ version, expected, description }) => {
      it(`should return ${expected} for schema version ${description}`, () => {
        expect(checkSchemaVersionSupportsModules(version)).toBe(expected);
      });
    });
  });

  describe("getSchemaVersionFromString", () => {
    describe("valid versions", () => {
      const testCases = [
        { input: "1.0", expected: SchemaVersion.V1_0 },
        { input: "1.1", expected: SchemaVersion.V1_1 },
        { input: "1.2", expected: SchemaVersion.V1_2 },
      ];

      testCases.forEach(({ input, expected }) => {
        it(`should return ${expected} for string '${input}'`, () => {
          expect(getSchemaVersionFromString(input)).toBe(expected);
        });
      });
    });

    describe("invalid versions", () => {
      const testCases = [
        { input: "1.5", description: "unsupported version '1.5'" },
        { input: "0.9", description: "unsupported version '0.9'" },
        { input: "invalid", description: "invalid version string 'invalid'" },
        { input: "", description: "empty string" },
      ];

      testCases.forEach(({ input, description }) => {
        it(`should throw error for ${description}`, () => {
          expect(() => getSchemaVersionFromString(input)).toThrow(`Unsupported schema version: ${input}`);
        });
      });
    });
  });
});
