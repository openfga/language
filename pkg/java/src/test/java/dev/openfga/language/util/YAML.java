package dev.openfga.language.util;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.dataformat.yaml.YAMLMapper;

import java.util.List;

public class YAML {

    public static <T> List<T> parseList(String json, TypeReference<List<T>> typeReference) throws JsonProcessingException {
        return new YAMLMapper().readValue(json, typeReference);
    }
}
