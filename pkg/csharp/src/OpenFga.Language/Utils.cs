namespace OpenFga.Language;

/// <summary>
/// Utility class providing helper methods for common operations.
/// </summary>
public static class Utils {
    /// <summary>
    /// Creates a deep copy of a nested dictionary structure.
    /// </summary>
    /// <param name="records">The dictionary to copy</param>
    /// <returns>A deep copy of the input dictionary</returns>
    public static Dictionary<string, Dictionary<string, bool>> DeepCopy(Dictionary<string, Dictionary<string, bool>> records) {
        var copy = new Dictionary<string, Dictionary<string, bool>>();
        foreach (var kvp in records) {
            copy[kvp.Key] = new Dictionary<string, bool>(kvp.Value);
        }
        return copy;
    }
}