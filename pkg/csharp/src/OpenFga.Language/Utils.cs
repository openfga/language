namespace OpenFga.Language;

public static class Utils
{
    public static List<T> EmptyIfNull<T>(List<T>? list)
    {
        return list ?? new List<T>();
    }

    public static Dictionary<string, Dictionary<string, bool>> DeepCopy(Dictionary<string, Dictionary<string, bool>> records)
    {
        var copy = new Dictionary<string, Dictionary<string, bool>>();
        foreach (var kvp in records)
        {
            copy[kvp.Key] = new Dictionary<string, bool>(kvp.Value);
        }
        return copy;
    }
}
