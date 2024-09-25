package dev.openfga.language.validation;

public class Validator {

    public class Rules {
        public static final String TYPE = "[^:#@\\*\\s]{1,254}";
        public static final String RELATION = "[^:#@\\*\\s]{1,50}";
        public static final String CONDITION = "[^\\*\\s]{1,50}";
        public static final String ID = "[^#:\\s*][a-zA-Z0-9_|*@.+]*";
        public static final String OBJECT = "[^\\s]{2,256}";
    }

    public static class Regexes {
        public static final ValidationRegex object =
                ValidationRegex.build("object", String.format("^%s$", Rules.OBJECT));

        public static final ValidationRegex objectId = ValidationRegex.build("object", String.format("^%s$", Rules.ID));

        public static final ValidationRegex typeId =
                ValidationRegex.build("object", String.format("^%s:%s$", Rules.TYPE, Rules.ID));

        public static final ValidationRegex relation =
                ValidationRegex.build("relation", String.format("^%s$", Rules.RELATION));

        public static final ValidationRegex userSet =
                ValidationRegex.build("userSet", String.format("^%s:%s#%s$", Rules.TYPE, Rules.ID, Rules.RELATION));

        public static final ValidationRegex userObject =
                ValidationRegex.build("userObject", String.format("^%s:%s$", Rules.TYPE, Rules.ID));

        public static final ValidationRegex userWildcard =
                ValidationRegex.build("userWildcard", String.format("^%s:\\*$", Rules.TYPE));

        public static final ValidationRegex condition =
                ValidationRegex.build("condition", String.format("^%s$", Rules.CONDITION));

        public static final ValidationRegex type =
                ValidationRegex.build("condition", String.format("^%s$", Rules.TYPE));
    }

    public static boolean validateObject(String object) {
        return Regexes.typeId.matches(object) && Regexes.object.matches(object);
    }

    public static boolean validateObjectId(String objectId) {
        return Regexes.objectId.matches(objectId);
    }

    public static boolean validateRelation(String relation) {
        return Regexes.relation.matches(relation);
    }

    public static boolean validateUserSet(String userset) {
        return Regexes.userSet.matches(userset);
    }

    public static boolean validateUserObject(String userObject) {
        return Regexes.userObject.matches(userObject);
    }

    public static boolean validateUserWildcard(String userWildcard) {
        return Regexes.userWildcard.matches(userWildcard);
    }

    public static boolean validateUser(String user) {
        return Regexes.userSet.matches(user) || Regexes.userObject.matches(user) || Regexes.userWildcard.matches(user);
    }

    public static boolean validateConditionName(String condition) {
        return Regexes.condition.matches(condition);
    }

    public static boolean validateType(String type) {
        return Regexes.type.matches(type);
    }
}
