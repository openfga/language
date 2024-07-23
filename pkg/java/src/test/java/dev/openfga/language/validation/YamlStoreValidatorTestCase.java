package dev.openfga.language.validation;

import static org.junit.jupiter.api.Assertions.assertNotNull;
import static org.junit.jupiter.api.Assertions.assertNull;

import org.junit.jupiter.api.Test;
import org.snakeyaml.engine.v2.api.Load;
import org.snakeyaml.engine.v2.api.LoadSettings;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;

public class YamlStoreValidatorTestCase {

  @Test
  public void testValidation() throws JsonProcessingException {
    var validator = new YamlStoreValidator();

    var testStore = "name: Valid Store File\n" + //
        "\n" + //
        "model: |\n" + //
        "  model\n" + //
        "    schema 1.1\n" + //
        "  type user\n" + //
        "  type organization\n" + //
        "    relations\n" + //
        "      define member: [user]\n" + //
        "\n" + //
        "tuples:\n" + //
        "  - user: user:daniel\n" + //
        "    relation: member\n" + //
        "    object: organization:auth0\n" + //
        "  - user: user:jeffery\n" + //
        "    relation: member\n" + //
        "    object: organization:auth0\n" + //
        "tests:\n" + //
        "  - name: Tests for organization members\n" + //
        "    check:\n" + //
        "      - user: user:daniel\n" + //
        "        object: organization:auth0\n" + //
        "        assertions:\n" + //
        "          member: true\n" + //
        "    list_objects:\n" + //
        "      - user: user:daniel\n" + //
        "        type: organization\n" + //
        "        assertions:\n" + //
        "          member:\n" + //
        "            - organization:auth0\n" + //
        "  - name: Tests for user member of organization\n" + //
        "    list_users:\n" + //
        "      - object: organization:auth0\n" + //
        "        user_filter:\n" + //
        "          - type: user\n" + //
        "        assertions:\n" + //
        "          member:\n" + //
        "            users:\n" + //
        "              - user:daniel\n" + //
        "              - user:jeffery\n" + //
        "";

    LoadSettings settings = LoadSettings.builder()
        .setUseMarks(true)
        .build();

    var load = new Load(settings);

    var loader = load.loadFromString(testStore);

    var mapper = new ObjectMapper();

    var mapped = mapper.writeValueAsString(loader);

    for (var error : validator.validate(mapped)) {
      assertNull(error);
    }

  }
}
