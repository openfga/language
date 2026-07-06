import js from "@eslint/js";
import tsPlugin from "@typescript-eslint/eslint-plugin";
import tsParser from "@typescript-eslint/parser";
import importPlugin from "eslint-plugin-import";
import eslintConfigPrettier from "eslint-config-prettier";

export default [
  {
    ignores: ["dist/", "gen/", "*.config.js", "*.config.mjs", ".eslintrc.js", "jest.config.js"],
  },
  js.configs.recommended,
  // TypeScript config
  {
    files: ["**/*.ts", "**/*.tsx", "**/*.mts", "**/*.cts"],
    plugins: {
      "@typescript-eslint": tsPlugin,
      import: importPlugin,
    },
    languageOptions: {
      ecmaVersion: 2021,
      sourceType: "module",
      parser: tsParser,
      parserOptions: {
        ecmaVersion: 2021,
        sourceType: "module",
      },
      globals: {
        window: "readonly",
        document: "readonly",
        navigator: "readonly",
        console: "readonly",
        process: "readonly",
        __dirname: "readonly",
        __filename: "readonly",
        Buffer: "readonly",
        global: "readonly",
        module: "readonly",
        require: "readonly",
        exports: "readonly",
        Promise: "readonly",
        Symbol: "readonly",
        WeakMap: "readonly",
        WeakSet: "readonly",
        Map: "readonly",
        Set: "readonly",
        Proxy: "readonly",
        Reflect: "readonly",
      },
    },

    settings: {
      "import/resolver": {
        typescript: {
          alwaysTryTypes: true,
          project: "./tsconfig.json",
        },
        node: {
          extensions: [".js", ".jsx", ".ts", ".tsx"],
        },
      },
    },

    rules: {
      ...tsPlugin.configs["eslint-recommended"].overrides[0].rules,
      ...tsPlugin.configs.recommended.rules,

      ...importPlugin.configs.recommended.rules,
      ...importPlugin.configs.typescript.rules,

      "no-case-declarations": "off",
      "linebreak-style": ["error", "unix"],
      "@typescript-eslint/ban-ts-comment": "off",
      quotes: ["error", "double"],
      semi: ["error", "always"],
      "max-len": [
        "warn",
        {
          code: 120,
        },
      ],
      "object-curly-spacing": ["error", "always"],
      "no-trailing-spaces": "error",
    },
  },

  // Must be last
  eslintConfigPrettier,
];
