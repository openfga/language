package dev.openfga.language;

import com.fasterxml.jackson.core.JsonLocation;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.JsonToken;
import com.fasterxml.jackson.dataformat.yaml.YAMLFactory;
import com.fasterxml.jackson.dataformat.yaml.YAMLParser;
import dev.openfga.language.errors.ErrorProperties;
import dev.openfga.language.errors.ModFileValidationError;
import dev.openfga.language.errors.ModFileValidationSingleError;
import dev.openfga.language.errors.StartEnd;
import java.io.IOException;
import java.net.URLDecoder;
import java.util.ArrayList;
import java.util.List;

public class FgaModTransformer {

    private final List<ModFileValidationSingleError> errors = new ArrayList<ModFileValidationSingleError>();
    private String modFileContents;

    public FgaModTransformer(String modFileContents) {
        this.modFileContents = modFileContents;
    }

    public static String transform(String modFileContents)
            throws IOException, JsonProcessingException, ModFileValidationError {
        return JSON.stringify(parse(modFileContents));
    }

    public static FgaModFile parse(String modFileContents)
            throws IOException, JsonProcessingException, ModFileValidationError {
        return new FgaModTransformer(modFileContents).parse();
    }

    public FgaModFile parse() throws IOException, JsonProcessingException, ModFileValidationError {
        YAMLParser parser = new YAMLFactory().createParser(modFileContents);
        FgaModFile modFile = new FgaModFile();
        ArrayList<String> seenFields = new ArrayList<String>();

        while (parser.nextToken() != JsonToken.END_OBJECT) {
            JsonToken currentToken = parser.currentToken();
            String name = parser.getCurrentName();

            if (currentToken != JsonToken.FIELD_NAME) {
                continue;
            }

            parser.nextToken();
            currentToken = parser.currentToken();
            JsonLocation location = parser.getCurrentLocation();

            if (name.equals("schema")) {
                seenFields.add("schema");
                handleSchema(parser, modFile, currentToken, location);
            } else if (name.equals("contents")) {
                seenFields.add("contents");
                handleContents(parser, modFile, currentToken, location);
            }
        }

        if (!seenFields.contains("schema")) {
            addError("missing schema field", new StartEnd(0, 0), new StartEnd(0, 0));
        }

        if (!seenFields.contains("contents")) {
            addError("missing contents field", new StartEnd(0, 0), new StartEnd(0, 0));
        }

        if (!errors.isEmpty()) {
            throw new ModFileValidationError(errors);
        }

        return modFile;
    }

    private void handleSchema(YAMLParser parser, FgaModFile modFile, JsonToken currentToken, JsonLocation location)
            throws IOException {
        String value = parser.getText();

        if (currentToken != JsonToken.VALUE_STRING) {
            StartEnd line = getLine(location);
            StartEnd column = getColumn(location, value);
            addError("unexpected schema type, expected string got value " + value, line, column);
        } else if (!value.equals("1.2")) {
            StartEnd line = getLine(location);
            StartEnd column = getColumn(location, "\"" + value + "\"");
            addError("unsupported schema version, fga.mod only supported in version `1.2`", line, column);
        } else {
            StartEnd line = getLine(location);
            // We have to wrap our value in quotes here as
            StartEnd column = getColumn(location, "\"" + value + "\"");

            ModFileStringProperty p =
                    new ModFileStringProperty().value(value).line(line).column(column);
            modFile.schema(p);
        }
    }

    private void handleContents(
            YAMLParser parser, FgaModFile modFile, JsonToken currentToken, JsonLocation startLocation)
            throws IOException {
        ModFileArrayProperty p = new ModFileArrayProperty();

        if (currentToken != JsonToken.START_ARRAY) {
            String value = parser.getText();
            StartEnd line = getLine(startLocation);
            StartEnd column = getColumn(startLocation, value);
            addError("unexpected contents type, expected list of strings got value " + value, line, column);
            return;
        }

        List<ModFileStringProperty> contents = new ArrayList<ModFileStringProperty>();
        while (parser.nextToken() != JsonToken.END_ARRAY) {
            currentToken = parser.currentToken();
            JsonLocation currentLoc = parser.getCurrentLocation();
            String rawValue = parser.getValueAsString();

            StartEnd line = getLine(currentLoc);
            StartEnd column = getColumn(currentLoc, rawValue);

            if (currentToken != JsonToken.VALUE_STRING) {
                addError("unexpected contents item type, expected string got value " + rawValue, line, column);
                continue;
            }

            String value = URLDecoder.decode(rawValue, "utf-8");

            value = value.replaceAll("\\\\", "/");

            if (value.contains("../") || value.startsWith("/")) {
                addError("invalid contents item " + rawValue, line, column);
                continue;
            }

            if (!value.endsWith(".fga")) {
                addError("contents items should use fga file extension, got " + value, line, column);
                continue;
            }

            contents.add(new ModFileStringProperty().value(value).line(line).column(column));
        }
        JsonLocation endLoc = parser.getCurrentLocation();

        p.line(new StartEnd(startLocation.getLineNr() - 1, endLoc.getLineNr() - 1))
                .column(new StartEnd(startLocation.getColumnNr() - 1, endLoc.getColumnNr() - 1))
                .value(contents);

        modFile.contents(p);
    }

    private StartEnd getLine(JsonLocation loc) {
        Integer line = loc.getLineNr() - 1;
        return new StartEnd(line, line);
    }

    private StartEnd getColumn(JsonLocation loc, String text) {
        // As the yaml parser does not expose an easy way of checking if the value is wrapped in
        // in quotes, scan the modFile string to see if it is and add an offset to the string length.
        Integer quotesOffset = 0;
        if (modFileContents.contains("'" + text + "'") || modFileContents.contains("\"" + text + "\"")) {
            quotesOffset = 2;
        }

        Integer columnEnd = loc.getColumnNr() - 1;
        Integer columnStart = columnEnd - (text.length() + quotesOffset);
        return new StartEnd(columnStart, columnEnd);
    }

    private void addError(String message, StartEnd line, StartEnd column) {
        errors.add(new ModFileValidationSingleError(new ErrorProperties(line, column, message)));
    }
}
