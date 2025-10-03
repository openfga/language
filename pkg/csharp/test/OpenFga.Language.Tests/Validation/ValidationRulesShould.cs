using OpenFga.Language.Validation;
using Xunit;

namespace OpenFga.Language.Tests.Validation;

public class ValidationRulesShould {
    private void ValidatedBadStructure(Func<string, bool> validate) {
        Assert.False(validate("item::1"));
        Assert.False(validate(":item:1"));
        Assert.False(validate("item:1:"));
        Assert.False(validate("item#relation"));
        Assert.False(validate("item:1##relation"));
        Assert.False(validate("#item:1"));
        Assert.False(validate("item:1#"));
        Assert.False(validate("ite#m:1"));
        Assert.False(validate("it*em:1"));
        Assert.False(validate("*:1"));
        Assert.False(validate("item*thing"));
        Assert.False(validate("item:*thing"));
        Assert.False(validate("item:**"));
    }

    [Fact]
    public void RuleObjectTest() {
        // Should pass '<type>:<id>'
        Assert.True(Validator.ValidateObject("document:1"));

        // Should fail if no ':' delimiter
        Assert.False(Validator.ValidateObject("document1"));

        // Should fail if includes relation
        Assert.False(Validator.ValidateObject("document:1#relation"));

        // Should fail if includes relation
        Assert.False(Validator.ValidateObject("document:*"));

        // Validate against bad formats
        ValidatedBadStructure(Validator.ValidateObject);
    }

    [Fact]
    public void TestValidateObjectId() {
        // Valid cases
        Assert.True(Validator.Regexes.ObjectId.Matches("document1"));
        Assert.True(Validator.Regexes.ObjectId.Matches("doc_123"));
        Assert.True(Validator.Regexes.ObjectId.Matches("user@domain.com"));
        Assert.True(Validator.Regexes.ObjectId.Matches("file.name"));
        Assert.True(Validator.Regexes.ObjectId.Matches("data+set"));
        Assert.True(Validator.Regexes.ObjectId.Matches("pipe|char"));
        Assert.True(Validator.Regexes.ObjectId.Matches("star*char"));
        Assert.True(Validator.Regexes.ObjectId.Matches("underscore_"));
        Assert.True(Validator.Regexes.ObjectId.Matches("pipe|underscore_@domain.com"));

        // Invalid cases
        Assert.False(Validator.Regexes.ObjectId.Matches("#document1"));
        Assert.False(Validator.Regexes.ObjectId.Matches(":doc123"));
        Assert.False(Validator.Regexes.ObjectId.Matches(" doc123"));
        Assert.False(Validator.Regexes.ObjectId.Matches("doc:123"));
        Assert.False(Validator.Regexes.ObjectId.Matches("doc#123"));
        Assert.False(Validator.Regexes.ObjectId.Matches("doc 123"));
        Assert.False(Validator.Regexes.ObjectId.Matches("doc:"));
        Assert.False(Validator.Regexes.ObjectId.Matches("    doc"));
    }

    [Fact]
    public void RuleUserTest() {
        // Should pass if UserSet
        Assert.True(Validator.ValidateUser("group:engineering#member"));

        // Should pass if UserObject
        Assert.True(Validator.ValidateUser("group:engineering"));

        // Should pass if UserWildcard
        Assert.True(Validator.ValidateUser("group:*"));

        // Should fail when missing <id>
        Assert.False(Validator.ValidateUser("group"));

        // Validate against bad formats
        ValidatedBadStructure(Validator.ValidateUser);
    }

    [Fact]
    public void RuleUserSetTest() {
        // Should pass if '<type>:<id>#<relation>'
        Assert.True(Validator.ValidateUserSet("group:engineering#member"));

        // Should fail for 'UserObject'
        Assert.False(Validator.ValidateUserSet("group:engineering"));

        // Should fail for 'UserWildcard'
        Assert.False(Validator.ValidateUserSet("group:*"));

        // Should fail if missing '<id>'
        Assert.False(Validator.ValidateUserSet("group"));

        ValidatedBadStructure(Validator.ValidateUserSet);
    }

    [Fact]
    public void RuleUserObjectTest() {
        // Should pass for '<type>:<id>'
        Assert.True(Validator.ValidateUserObject("group:engineering"));

        // Should fail if contains '#'
        Assert.False(Validator.ValidateUserObject("group:engineering#member"));

        // Should fail if 'Wildcard' is present
        Assert.False(Validator.ValidateUserObject("group:*"));

        // Should fail if missing '<id>'
        Assert.False(Validator.ValidateUserObject("group"));

        ValidatedBadStructure(Validator.ValidateUserObject);
    }

    [Fact]
    public void RuleUserWildcardTest() {
        // Should pass for '<type>:*'
        Assert.True(Validator.ValidateUserWildcard("group:*"));

        // Should fail for '<type>:<id>'
        Assert.False(Validator.ValidateUserWildcard("group:engineering"));

        // Should fail if contains '#'
        Assert.False(Validator.ValidateUserWildcard("group:engineering#member"));

        // Should fail if missing '*'
        Assert.False(Validator.ValidateUserObject("group"));

        ValidatedBadStructure(Validator.ValidateUserWildcard);
    }

    [Fact]
    public void RuleTypeTest() {
        // Should pass '<types>'
        Assert.True(Validator.ValidateType("folder"));

        // Should fail 'UserObject'
        Assert.False(Validator.ValidateType("folder:1"));

        // Should fail UserSet
        Assert.False(Validator.ValidateType("folder:1#relation"));

        ValidatedBadStructure(Validator.ValidateType);
    }
}