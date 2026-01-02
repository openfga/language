using OpenFga.Sdk.Model;

namespace OpenFga.Language;

/// <summary>
/// Represents a partial condition parameter type reference used during parsing.
/// This class is used internally during DSL parsing to build condition parameter type references.
/// </summary>
public class PartialConditionParamTypeRef {

    /// <summary>
    /// Gets or sets the type name for this parameter.
    /// </summary>
    public TypeName TypeName { get; set; }

    /// <summary>
    /// Gets or sets the generic types for this parameter (if applicable).
    /// </summary>
    public List<ConditionParamTypeRef>? GenericTypes { get; set; }

    /// <summary>
    /// Converts this partial type reference to a complete ConditionParamTypeRef.
    /// </summary>
    /// <returns>A complete ConditionParamTypeRef instance</returns>
    public ConditionParamTypeRef AsConditionParamTypeRef() {
        return new ConditionParamTypeRef() {
            TypeName = TypeName,
            GenericTypes = GenericTypes
        };
    }
}