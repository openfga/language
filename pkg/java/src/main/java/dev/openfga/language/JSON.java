package dev.openfga.language;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.MapperFeature;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.SerializationFeature;

class JSON {

    public static <T> T parse(String json, Class<T> type) throws JsonProcessingException {
        return new ObjectMapper().readValue(json, type);
    }

    public static String stringify(Object object) throws JsonProcessingException {
        var mapper = new ObjectMapper();
        mapper.setConfig(mapper.getSerializationConfig()
                .with(MapperFeature.SORT_PROPERTIES_ALPHABETICALLY)
                .with(SerializationFeature.ORDER_MAP_ENTRIES_BY_KEYS));

        return mapper.writeValueAsString(object);
    }

    private JSON() {}
}
