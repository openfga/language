package dev.openfga.language;

import dev.openfga.sdk.api.model.ConditionParamTypeRef;
import dev.openfga.sdk.api.model.TypeName;

import java.util.ArrayList;
import java.util.List;

public class PartialConditionParamTypeRef {
    private TypeName typeName;
    private List<ConditionParamTypeRef> genericTypes = new ArrayList<>();

    public ConditionParamTypeRef asConditionParamTypeRef() {
        return new ConditionParamTypeRef()
                .typeName(typeName)
                .genericTypes(genericTypes);
    }

    public TypeName getTypeName() {
        return typeName;
    }

    public void setTypeName(TypeName typeName) {
        this.typeName = typeName;
    }

    public List<ConditionParamTypeRef> getGenericTypes() {
        return genericTypes;
    }

    public void setGenericTypes(List<ConditionParamTypeRef> genericTypes) {
        this.genericTypes = genericTypes;
    }
}
