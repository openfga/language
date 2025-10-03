namespace OpenFga.Language;

public static class Utils
{
    public static U? GetNullSafe<T, U>(T? item, Func<T, U> getter) where T : class
    {
        return item == null ? default : getter(item);
    }

    public static List<U> GetNullSafeList<T, U>(T? item, Func<T, List<U>?> getter) where T : class
    {
        var list = item == null ? null : getter(item);
        return EmptyIfNull(list);
    }

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
