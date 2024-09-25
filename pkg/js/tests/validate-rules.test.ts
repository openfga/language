import { describe, expect, it } from "@jest/globals";
import { Validator } from "../validator/validate-rules";

// These tests are a subset of bad formats that all validation rules that should fail
const validatedBadStructure = (validator: (value: string) => boolean) => {
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
      expect(Validator.object("document:1")).toBeTruthy();
    });

    it("should fail if no ':' delimiter", () => {
      expect(Validator.object("document1")).toBeFalsy();
    });

    it("should fail if includes relation", () => {
      expect(Validator.object("document:1#relation")).toBeFalsy();
    });

    it("should fail if includes '*'", () => {
      expect(Validator.object("document:*")).toBeFalsy();
    });

    validatedBadStructure(Validator.object);
  });

  describe("Rule 'User'", () => {
    it("shoud pass for 'UserSet'", () => {
      expect(Validator.user("group:engineering#member")).toBeTruthy();
    });

    it("shoud pass for 'UserObject'", () => {
      expect(Validator.user("group:engineering")).toBeTruthy();
    });

    it("shoud pass for 'UserWildcard'", () => {
      expect(Validator.user("group:*")).toBeTruthy();
    });

    it("shoud fail when missing <id>'", () => {
      expect(Validator.user("group")).toBeFalsy();
    });

    validatedBadStructure(Validator.user);
  });

  describe("Rule 'UserSet'", () => {
    it("should pass '<type>:<id>#<relation>'", () => {
      expect(Validator.userSet("group:engineering#member")).toBeTruthy();
    });

    it("shoud fail for 'UserObject'", () => {
      expect(Validator.userSet("group:engineering")).toBeFalsy();
    });

    it("shoud fail for 'UserWildcard'", () => {
      expect(Validator.userSet("group:*")).toBeFalsy();
    });

    it("should fail if missing <id>", () => {
      expect(Validator.userSet("group")).toBeFalsy();
    });

    validatedBadStructure(Validator.userSet);
  });

  describe("Rule 'UserObject'", () => {
    it("should pass '<type>:<id>", () => {
      expect(Validator.userObject("group:engineering")).toBeTruthy();
    });

    it("should fail if contains '#", () => {
      expect(Validator.userObject("group:engineering#member")).toBeFalsy();
    });

    it("should fail if 'Wildcard' is present", () => {
      expect(Validator.userObject("group:*")).toBeFalsy();
    });

    it("should fail if missing <id>", () => {
      expect(Validator.userObject("group")).toBeFalsy();
    });

    validatedBadStructure(Validator.userObject);
  });

  describe("Rule 'UserWildcard'", () => {
    it("should pass '<type>:*", () => {
      expect(Validator.userWildcard("group:*")).toBeTruthy();
    });

    it("should fail if missing 'UserObject'", () => {
      expect(Validator.userWildcard("group:organization")).toBeFalsy();
    });

    it("should fail if contains '#", () => {
      expect(Validator.userWildcard("group:engineering#member")).toBeFalsy();
    });

    it("should fail if missing '*' delimiter", () => {
      expect(Validator.userWildcard("group")).toBeFalsy();
    });

    validatedBadStructure(Validator.userWildcard);
  });

  describe("Rule 'Type'", () => {
    it("should pass '<type>", () => {
      expect(Validator.type("folder")).toBeTruthy();
    });

    it("should fail 'UserObject", () => {
      expect(Validator.type("folder:1")).toBeFalsy();
    });

    it("should fail 'UserSet'", () => {
      expect(Validator.type("folder:1#relation")).toBeFalsy();
    });

    validatedBadStructure(Validator.type);
  });

  describe("Rule 'id'", () => {
    it("should pass 'document1'", () => {
      expect(Validator.objectId("document1")).toBeTruthy();
    });

    it("should pass 'doc_123'", () => {
      expect(Validator.objectId("doc_123")).toBeTruthy();
    });

    it("should pass 'user@domain.com'", () => {
      expect(Validator.objectId("user@domain.com")).toBeTruthy();
    });

    it("should pass 'file.name'", () => {
      expect(Validator.objectId("file.name")).toBeTruthy();
    });

    it("should pass 'data+set'", () => {
      expect(Validator.objectId("data+set")).toBeTruthy();
    });

    it("should pass 'pipe|char'", () => {
      expect(Validator.objectId("pipe|char")).toBeTruthy();
    });

    it("should pass 'star*char'", () => {
      expect(Validator.objectId("star*char")).toBeTruthy();
    });

    it("should pass 'underscore_'", () => {
      expect(Validator.objectId("underscore_")).toBeTruthy();
    });

    it("should pass 'pipe|underscore_@domain.com'", () => {
      expect(Validator.objectId("pipe|underscore_@domain.com")).toBeTruthy();
    });

    it("should fail '#document1'", () => {
      expect(Validator.objectId("#document1")).toBeFalsy();
    });

    it("should fail ':doc123'", () => {
      expect(Validator.objectId(":doc123")).toBeFalsy();
    });

    it("should fail ' doc123'", () => {
      expect(Validator.objectId(" doc123")).toBeFalsy();
    });

    it("should fail 'doc*123'", () => {
      expect(Validator.objectId("doc*123")).toBeTruthy();
    });

    it("should fail 'doc:123'", () => {
      expect(Validator.objectId("doc:123")).toBeFalsy();
    });

    it("should fail 'doc#123'", () => {
      expect(Validator.objectId("doc#123")).toBeFalsy();
    });

    it("should fail 'doc 123'", () => {
      expect(Validator.objectId("doc 123")).toBeFalsy();
    });

    it("should fail 'doc*'", () => {
      expect(Validator.objectId("doc*")).toBeTruthy();
    });

    it("should fail 'doc:'", () => {
      expect(Validator.objectId("doc:")).toBeFalsy();
    });
  });
});
