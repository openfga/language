import js from "@eslint/js";
import { defineConfig, globalIgnores } from "eslint/config";
import eslintConfigPrettier from "eslint-config-prettier/flat";
import importX from "eslint-plugin-import-x";
import tseslint from "typescript-eslint";

export default defineConfig([
  // `gen/` is ANTLR-generated output; `dist/` is the tsc build output.
  globalIgnores(["dist/", "gen/"]),
  {
    files: ["**/*.{ts,tsx,mts,cts}"],
    extends: [
      js.configs.recommended,
      tseslint.configs.recommended,
      importX.flatConfigs.recommended,
      importX.flatConfigs.typescript,
    ],
    rules: {
      // Four `case` clauses declare variables without braces; silenced since f060aae (original reason not recorded).
      "no-case-declarations": "off",
      // Tests use bare `@ts-expect-error` to pass invalid inputs; silenced since f060aae (original reason not recorded).
      "@typescript-eslint/ban-ts-comment": "off",
    },
  },
  // Must stay last: turns off any rule that conflicts with Prettier.
  eslintConfigPrettier,
]);
