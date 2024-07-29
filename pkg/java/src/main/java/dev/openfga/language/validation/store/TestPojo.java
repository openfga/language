package dev.openfga.language.validation.store;

import com.fasterxml.jackson.annotation.JsonProperty;
import dev.openfga.sdk.api.model.*;


public class TestPojo {

    public String name;

    @JsonProperty(value="tuple_file")
    public String tupleFile;

    public TupleKey[] tuples;

    public CheckPojo[] check;

    @JsonProperty(value = "list_objects")
    public ListObjectPojo[] listObjects;

    @JsonProperty(value = "list_users")
    public ListUsersPojo[] listUsers;

}
