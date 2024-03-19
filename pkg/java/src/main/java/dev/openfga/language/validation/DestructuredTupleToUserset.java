package dev.openfga.language.validation;

class DestructuredTupleToUserset {
    private final String decodedType;
    private final String decodedRelation;
    private final boolean wildcard;
    private final String decodedConditionName;


    public DestructuredTupleToUserset(String decodedType, String decodedRelation, boolean wildcard, String decodedConditionName) {
        this.decodedType = decodedType;
        this.decodedRelation = decodedRelation;
        this.wildcard = wildcard;
        this.decodedConditionName = decodedConditionName;
    }

    public String getDecodedType() {
        return decodedType;
    }

    public String getDecodedRelation() {
        return decodedRelation;
    }

    public boolean isWildcard() {
        return wildcard;
    }

    public String getDecodedConditionName() {
        return decodedConditionName;
    }

    public static DestructuredTupleToUserset from(String allowableType) {
        var tupleAndCondition = allowableType.split(" with ");
        var tupleString = tupleAndCondition[0];
        var decodedConditionName = tupleAndCondition.length > 1 ? tupleAndCondition[1] : null;
        var isWildcard = tupleString.contains(":*");
        var splittedWords = tupleString.replace(":*", "").split("#");
        return new DestructuredTupleToUserset(
                splittedWords[0],
                splittedWords.length > 1 ? splittedWords[1] : null,
                isWildcard,
                decodedConditionName);
    }

}
