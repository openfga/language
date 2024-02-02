package dev.openfga.language;

import java.util.List;
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

    private Utils() {
    }
}
