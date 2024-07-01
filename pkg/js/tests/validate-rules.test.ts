import { describe, expect, it } from "@jest/globals";
import { Validate } from "../validator/validate-rules";

describe("validation rules", () => {
  describe("Rule 'Object'", () => {
    it("should pass '<object>:<id>'", () => {
      expect(Validate.object("document:1")).toBeTruthy();
    });

    it("should fail if no ':' delimiter", () => {
      expect(Validate.object("document1")).toBeFalsy();
    });
  });

  describe("Rule 'User'", () => {
    it("shoud pass for 'UserSet'", () => {
      expect(Validate.user("group:engineering#member")).toBeTruthy();
    });
    it("shoud pass for 'UserObject'", () => {
      expect(Validate.user("group:engineering")).toBeTruthy();
    });
    it("shoud pass for 'UserWildcard'", () => {
      expect(Validate.user("group:*")).toBeTruthy();
    });
  });

  describe("Rule 'UserSet'", () => {
    it("should pass '<type>:<id>#<relation>'", () => {
      expect(Validate.userSet("group:engineering#member")).toBeTruthy();
    });

    it("should fail if no '#' delimiter", () => {
      expect(Validate.userSet("group:engineering")).toBeFalsy();
    });

    it("should fail if no ':' delimiter", () => {
      expect(Validate.userSet("group#member")).toBeFalsy();
    });
  });

  describe("Rule 'UserObject'", () => {
    it("should pass '<type>:<id>", () => {
      expect(Validate.userObject("group:engineering")).toBeTruthy();
    });

    it("should fail if contains '#", () => {
      expect(Validate.userObject("group:engineering#member")).toBeFalsy();
    });

    it("should fail if no ':' delimiter", () => {
      expect(Validate.userObject("group#member")).toBeFalsy();
    });
  });

  describe("Rule 'UserWildcard'", () => {
    it("should pass '<type>:*", () => {
      expect(Validate.userWildcard("group:*")).toBeTruthy();
    });

    it("should fail if missing '*' delimiter", () => {
      expect(Validate.userWildcard("group")).toBeFalsy();
    });
  });

  describe("Rule 'Type'", () => {
    it("should pass '<type>", () => {
      expect(Validate.type("folder")).toBeTruthy();
    });
  });
});
