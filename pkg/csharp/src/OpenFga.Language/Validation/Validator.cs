namespace OpenFga.Language.Validation;

/// <summary>
/// Provides validation rules and methods for OpenFGA authorization model components.
/// Contains predefined regex patterns and validation methods for types, relations, objects, and users.
/// </summary>
public class Validator {
    internal static class Rules {
        public const string Type = "[^:#@\\*\\s]{1,254}";
        public const string Relation = "[^:#@\\*\\s]{1,50}";
        public const string Condition = "[^\\*\\s]{1,50}";
        public const string Id = "[^#:\\s*][a-zA-Z0-9_|*@.+]*";
        public const string Object = "[^\\s]{2,256}";
    }

    internal static class Regexes {
        public static readonly ValidationRegex Object =
            ValidationRegex.Build("object", $"^{Rules.Object}$");

        public static readonly ValidationRegex ObjectId =
            ValidationRegex.Build("object", $"^{Rules.Id}$");

        public static readonly ValidationRegex TypeId =
            ValidationRegex.Build("object", $"^{Rules.Type}:{Rules.Id}$");

        public static readonly ValidationRegex Relation =
            ValidationRegex.Build("relation", $"^{Rules.Relation}$");

        public static readonly ValidationRegex UserSet =
            ValidationRegex.Build("userSet", $"^{Rules.Type}:{Rules.Id}#{Rules.Relation}$");

        public static readonly ValidationRegex UserObject =
            ValidationRegex.Build("userObject", $"^{Rules.Type}:{Rules.Id}$");

        public static readonly ValidationRegex UserWildcard =
            ValidationRegex.Build("userWildcard", $"^{Rules.Type}:\\*$");

        public static readonly ValidationRegex Condition =
            ValidationRegex.Build("condition", $"^{Rules.Condition}$");

        public static readonly ValidationRegex Type =
            ValidationRegex.Build("condition", $"^{Rules.Type}$");
    }

    /// <summary>
    /// Validates that an object string matches the required format (type:ID).
    /// </summary>
    /// <param name="objectValue">The object string to validate</param>
    /// <returns>True if the object string is valid, false otherwise</returns>
    public static bool ValidateObject(string objectValue) {
        return Regexes.TypeId.Matches(objectValue) && Regexes.Object.Matches(objectValue);
    }

    /// <summary>
    /// Validates that an object ID matches the required format.
    /// </summary>
    /// <param name="objectId">The object ID to validate</param>
    /// <returns>True if the object ID is valid, false otherwise</returns>
    public static bool ValidateObjectId(string objectId) {
        return Regexes.ObjectId.Matches(objectId);
    }

    /// <summary>
    /// Validates that a relation name matches the required format.
    /// </summary>
    /// <param name="relation">The relation name to validate</param>
    /// <returns>True if the relation name is valid, false otherwise</returns>
    public static bool ValidateRelation(string relation) {
        return Regexes.Relation.Matches(relation);
    }

    /// <summary>
    /// Validates that a user set string matches the required format (type:ID#relation).
    /// </summary>
    /// <param name="userset">The user set string to validate</param>
    /// <returns>True if the user set string is valid, false otherwise</returns>
    public static bool ValidateUserSet(string userset) {
        return Regexes.UserSet.Matches(userset);
    }

    /// <summary>
    /// Validates that a user object string matches the required format (type:ID).
    /// </summary>
    /// <param name="userObject">The user object string to validate</param>
    /// <returns>True if the user object string is valid, false otherwise</returns>
    public static bool ValidateUserObject(string userObject) {
        return Regexes.UserObject.Matches(userObject);
    }

    /// <summary>
    /// Validates that a user wildcard string matches the required format (type:*).
    /// </summary>
    /// <param name="userWildcard">The user wildcard string to validate</param>
    /// <returns>True if the user wildcard string is valid, false otherwise</returns>
    public static bool ValidateUserWildcard(string userWildcard) {
        return Regexes.UserWildcard.Matches(userWildcard);
    }

    /// <summary>
    /// Validates that a user string matches any valid user format (user set, user object, or user wildcard).
    /// </summary>
    /// <param name="user">The user string to validate</param>
    /// <returns>True if the user string is valid, false otherwise</returns>
    public static bool ValidateUser(string user) {
        return Regexes.UserSet.Matches(user) || Regexes.UserObject.Matches(user) || Regexes.UserWildcard.Matches(user);
    }

    /// <summary>
    /// Validates that a condition name matches the required format.
    /// </summary>
    /// <param name="condition">The condition name to validate</param>
    /// <returns>True if the condition name is valid, false otherwise</returns>
    public static bool ValidateConditionName(string condition) {
        return Regexes.Condition.Matches(condition);
    }

    /// <summary>
    /// Validates that a type name matches the required format.
    /// </summary>
    /// <param name="type">The type name to validate</param>
    /// <returns>True if the type name is valid, false otherwise</returns>
    public static bool ValidateType(string type) {
        return Regexes.Type.Matches(type);
    }
}