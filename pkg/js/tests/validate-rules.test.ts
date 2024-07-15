import { describe, expect, it } from "@jest/globals";
import { Validate } from "../validator/validate-rules";

// These tests are a subset of bad formats that all validation rules that should fail
const validatedBadStructure = (validator: any) => {
  it("should fail if '::' is the delimiter", () => {
    expect(validator("item::1")).toBeFalsy();
  });

  it("should fail if ':' at start", () => {
    expect(validator(":item:1")).toBeFalsy();
  });

  it("should fail if ':' at end", () => {
    expect(validator("item:1:")).toBeFalsy();
  });

  it("should fail if '#' is present, but not ':'", () => {
    expect(validator("item#relation")).toBeFalsy();
  });

  it("should fail if '##' is a delimiter", () => {
    expect(validator("item:1##relation")).toBeFalsy();
  });

  it("should fail if '#' at start", () => {
    expect(validator("#item:1")).toBeFalsy();
  });

  it("should fail if '#' at end", () => {
    expect(validator("item:1#")).toBeFalsy();
  });

  it("should fail if '#' is before ':'", () => {
    expect(validator("ite#m:1")).toBeFalsy();
  });

  it("should fail if wildcard present in object name", () => {
    expect(validator("it*em:1")).toBeFalsy();
  });

  it("should fail if wildcard at start", () => {
    expect(validator("*:1")).toBeFalsy();
  });

  it("should fail if wildcard has no ':'", () => {
    expect(validator("item*thing")).toBeFalsy();
  });

  it("should fail if wildcard is proceeded by text", () => {
    expect(validator("item:*thing")).toBeFalsy();
  });

  it("should fail if wildcard appears multiple times", () => {
    expect(validator("item:**")).toBeFalsy();
  });
};

describe("Validation Rules", () => {
  describe("Rule 'Object'", () => {
    it("should pass '<object>:<id>'", () => {
      expect(Validate.object("document:1")).toBeTruthy();
    });

    it("should fail if no ':' delimiter", () => {
      expect(Validate.object("document1")).toBeFalsy();
    });

    it("should fail if includes relation", () => {
      expect(Validate.object("document:1#relation")).toBeFalsy();
    });

    validatedBadStructure(Validate.object);
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

    it("shoud fail when missing <id>'", () => {
      expect(Validate.user("group")).toBeFalsy();
    });

    validatedBadStructure(Validate.user);
  });

  describe("Rule 'UserSet'", () => {
    it("should pass '<type>:<id>#<relation>'", () => {
      expect(Validate.userSet("group:engineering#member")).toBeTruthy();
    });

    it("shoud fail for 'UserObject'", () => {
      expect(Validate.userSet("group:engineering")).toBeFalsy();
    });

    it("shoud fail for 'UserWildcard'", () => {
      expect(Validate.userSet("group:*")).toBeFalsy();
    });

    it("should fail if missing <id>", () => {
      expect(Validate.userSet("group")).toBeFalsy();
    });

    validatedBadStructure(Validate.userSet);
  });

  describe("Rule 'UserObject'", () => {
    it("should pass '<type>:<id>", () => {
      expect(Validate.userObject("group:engineering")).toBeTruthy();
    });

    it("should fail if contains '#", () => {
      expect(Validate.userObject("group:engineering#member")).toBeFalsy();
    });

    it("should fail if 'Wildcard' is present", () => {
      expect(Validate.userObject("group:*")).toBeFalsy();
    });

    it("should fail if missing <id>", () => {
      expect(Validate.userObject("group")).toBeFalsy();
    });

    validatedBadStructure(Validate.userObject);
  });

  describe("Rule 'UserWildcard'", () => {
    it("should pass '<type>:*", () => {
      expect(Validate.userWildcard("group:*")).toBeTruthy();
    });

    it("should fail if missing 'UserObject'", () => {
      expect(Validate.userWildcard("group:organization")).toBeFalsy();
    });

    it("should fail if missing '*' delimiter", () => {
      expect(Validate.userWildcard("group")).toBeFalsy();
    });

    validatedBadStructure(Validate.userWildcard);
  });

  describe("Rule 'Type'", () => {
    it("should pass '<type>", () => {
      expect(Validate.type("folder")).toBeTruthy();
    });

    it("should fail 'UserObject", () => {
      expect(Validate.type("folder:1")).toBeFalsy();
    });

    it("should fail 'UserSet'", () => {
      expect(Validate.type("folder:1#relation")).toBeFalsy();
    });

    validatedBadStructure(Validate.type);
  });
});
