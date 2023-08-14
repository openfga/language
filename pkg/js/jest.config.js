module.exports = {
  testEnvironment: "node",
  preset: "ts-jest",
  coveragePathIgnorePatterns: ["/node_modules/", "/tests/"],
  moduleFileExtensions: ["js", "d.ts", "ts", "json"],
  reporters: [
    "default",
    ["./node_modules/jest-html-reporter", {
      "pageTitle": "Test Report",
      includeFailureMsg: true,
      includeConsoleLog: true,
    }],
  ],
};
