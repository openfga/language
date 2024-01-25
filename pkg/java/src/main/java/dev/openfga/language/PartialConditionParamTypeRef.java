package dev.openfga.language;

import dev.openfga.sdk.api.model.ConditionParamTypeRef;
import dev.openfga.sdk.api.model.TypeName;
import lombok.Getter;
import lombok.Setter;

import java.util.ArrayList;
import java.util.List;

@Getter
@Setter
public class PartialConditionParamTypeRef {
    private TypeName typeName;
    private List<ConditionParamTypeRef> genericTypes = new ArrayList<>();

    public ConditionParamTypeRef asConditionParamTypeRef() {
        return new ConditionParamTypeRef()
                .typeName(typeName)
                .genericTypes(genericTypes);
    }
}
