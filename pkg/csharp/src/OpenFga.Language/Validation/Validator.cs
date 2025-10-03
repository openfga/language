namespace OpenFga.Language.Validation;

public class Validator {
    public static class Rules {
        public const string Type = "[^:#@\\*\\s]{1,254}";
        public const string Relation = "[^:#@\\*\\s]{1,50}";
        public const string Condition = "[^\\*\\s]{1,50}";
        public const string Id = "[^#:\\s*][a-zA-Z0-9_|*@.+]*";
        public const string Object = "[^\\s]{2,256}";
    }

    public static class Regexes {
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

    public static bool ValidateObject(string objectValue) {
        return Regexes.TypeId.Matches(objectValue) && Regexes.Object.Matches(objectValue);
    }

    public static bool ValidateObjectId(string objectId) {
        return Regexes.ObjectId.Matches(objectId);
    }

    public static bool ValidateRelation(string relation) {
        return Regexes.Relation.Matches(relation);
    }

    public static bool ValidateUserSet(string userset) {
        return Regexes.UserSet.Matches(userset);
    }

    public static bool ValidateUserObject(string userObject) {
        return Regexes.UserObject.Matches(userObject);
    }

    public static bool ValidateUserWildcard(string userWildcard) {
        return Regexes.UserWildcard.Matches(userWildcard);
    }

    public static bool ValidateUser(string user) {
        return Regexes.UserSet.Matches(user) || Regexes.UserObject.Matches(user) || Regexes.UserWildcard.Matches(user);
    }

    public static bool ValidateConditionName(string condition) {
        return Regexes.Condition.Matches(condition);
    }

    public static bool ValidateType(string type) {
        return Regexes.Type.Matches(type);
    }
}