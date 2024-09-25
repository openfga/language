package dev.openfga.language.validation;

import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertTrue;

import java.util.function.Function;
import org.junit.jupiter.api.Test;

public class ValidationRules {

    public void validatedBadStructure(Function<String, Boolean> validate) {
        assertFalse(validate.apply("item::1"));
        assertFalse(validate.apply(":item:1"));
        assertFalse(validate.apply("item:1:"));
        assertFalse(validate.apply("item#relation"));
        assertFalse(validate.apply("item:1##relation"));
        assertFalse(validate.apply("#item:1"));
        assertFalse(validate.apply("item:1#"));
        assertFalse(validate.apply("ite#m:1"));
        assertFalse(validate.apply("it*em:1"));
        assertFalse(validate.apply("*:1"));
        assertFalse(validate.apply("item*thing"));
        assertFalse(validate.apply("item:*thing"));
        assertFalse(validate.apply("item:**"));
    }

    @Test
    public void ruleObjectTest() {

        // Should pass '<type>:<id>'
        assertTrue(Validator.validateObject("document:1"));

        // Should fail if no ':' delimiter
        assertFalse(Validator.validateObject("document1"));

        // Should fail if includes relation
        assertFalse(Validator.validateObject("document:1#relation"));

        // Should fail if includes relation
        assertFalse(Validator.validateObject("document:*"));

        // Validate against bad formats
        validatedBadStructure(Validator::validateObject);
    }

    @Test
    public void testValidateObjectId() {
        // Valid cases
        assertTrue(Validator.Regexes.objectId.matches("document1"));
        assertTrue(Validator.Regexes.objectId.matches("doc_123"));
        assertTrue(Validator.Regexes.objectId.matches("user@domain.com"));
        assertTrue(Validator.Regexes.objectId.matches("file.name"));
        assertTrue(Validator.Regexes.objectId.matches("data+set"));
        assertTrue(Validator.Regexes.objectId.matches("pipe|char"));
        assertTrue(Validator.Regexes.objectId.matches("star*char"));
        assertTrue(Validator.Regexes.objectId.matches("underscore_"));
        assertTrue(Validator.Regexes.objectId.matches("pipe|underscore_@domain.com"));

        // Invalid cases
        assertFalse(Validator.Regexes.objectId.matches("#document1"));
        assertFalse(Validator.Regexes.objectId.matches(":doc123"));
        assertFalse(Validator.Regexes.objectId.matches(" doc123"));
        assertFalse(Validator.Regexes.objectId.matches("doc:123"));
        assertFalse(Validator.Regexes.objectId.matches("doc#123"));
        assertFalse(Validator.Regexes.objectId.matches("doc 123"));
        assertFalse(Validator.Regexes.objectId.matches("doc:"));
        assertFalse(Validator.Regexes.objectId.matches("    doc"));
    }

    @Test
    public void ruleUserTest() {

        // Should pass if UserSet
        assertTrue(Validator.validateUser("group:engineering#member"));

        // Should pass if UserObject
        assertTrue(Validator.validateUser("group:engineering"));

        // Should pass if UserWildcard
        assertTrue(Validator.validateUser("group:*"));

        // Should fail when missing <id>
        assertFalse(Validator.validateUser("group"));

        // Validate against bad formats
        validatedBadStructure(Validator::validateUser);
    }

    @Test
    public void ruleUserSetTest() {

        // Should pass if '<type>:<id>#<relation>'
        assertTrue(Validator.validateUserSet("group:engineering#member"));

        // Shoud fail for 'UserObject'
        assertFalse(Validator.validateUserSet("group:engineering"));

        // Shoud fail for 'UserWildcard'
        assertFalse(Validator.validateUserSet("group:*"));

        // Shoud fail if missing '<id>'
        assertFalse(Validator.validateUserSet("group"));

        validatedBadStructure(Validator::validateUserSet);
    }

    @Test
    public void ruleUserObjectTest() {

        // Should pass for '<type>:<id>'
        assertTrue(Validator.validateUserObject("group:engineering"));

        // Should fail if contains '#'
        assertFalse(Validator.validateUserObject("group:engineering#member"));

        // Should fail if 'Wildcard' is present
        assertFalse(Validator.validateUserObject("group:*"));

        // Should fail if missing '<id>'
        assertFalse(Validator.validateUserObject("group"));

        validatedBadStructure(Validator::validateUserObject);
    }

    @Test
    public void ruleUserWildcardTest() {

        // Should pass for '<type>:*'
        assertTrue(Validator.validateUserWildcard("group:*"));

        // Should fail for '<type>:<id>'
        assertFalse(Validator.validateUserWildcard("group:engineering"));

        // Should fail if contains '#'
        assertFalse(Validator.validateUserWildcard("group:engineering#member"));

        // Should fail if missing '*'
        assertFalse(Validator.validateUserObject("group"));

        validatedBadStructure(Validator::validateUserWildcard);
    }

    @Test
    public void ruleTypeTest() {

        // Should pass '<types>'
        assertTrue(Validator.validateType("folder"));

        // Should ail 'UserObject'
        assertFalse(Validator.validateType("folder:1"));

        // Should fail UserSet
        assertFalse(Validator.validateType("folder:1#relation"));

        validatedBadStructure(Validator::validateType);
    }
}
