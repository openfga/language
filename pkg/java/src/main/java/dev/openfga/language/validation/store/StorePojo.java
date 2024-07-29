package dev.openfga.language.validation.store;

import com.fasterxml.jackson.annotation.JsonProperty;
import dev.openfga.sdk.api.model.TupleKey;

public class StorePojo {

    public String name;

    @JsonProperty(value="model_file")
    public String modelFile;

    public String model;

    @JsonProperty(value="tuple_file")
    public String tupleFile;

    public TupleKey[] tuples;

    public TestPojo[] tests;

}
