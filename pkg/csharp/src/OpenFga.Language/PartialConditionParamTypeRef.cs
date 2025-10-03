
using OpenFga.Sdk.Model;

namespace OpenFga.Language;

public class PartialConditionParamTypeRef {
    public TypeName TypeName { get; set; }
    public List<ConditionParamTypeRef>? GenericTypes { get; set; }

    public ConditionParamTypeRef AsConditionParamTypeRef() {
        return new ConditionParamTypeRef()
            {
                TypeName = TypeName,
                GenericTypes = GenericTypes
            };
    }
}
