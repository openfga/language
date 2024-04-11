package dev.openfga.language;

import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.function.Function;

public class Utils {

    public static <T, U> U getNullSafe(T item, Function<T, U> getter) {
        return item == null ? null : getter.apply(item);
    }

    public static <T, U> List<U> getNullSafeList(T item, Function<T, ? extends List<U>> getter) {
        var list = item == null ? null : getter.apply(item);
        return emptyIfNull(list);
    }

    public static <T> List<T> emptyIfNull(List<T> list) {
        return list == null ? List.of() : list;
    }

    public static Map<String, Map<String, Boolean>> deepCopy(Map<String, Map<String, Boolean>> records) {
        Map<String, Map<String, Boolean>> copy = new HashMap<>();
        records.forEach((key, value) -> copy.put(key, new HashMap<>(value)));
        return copy;
    }

    private Utils() {}
}
