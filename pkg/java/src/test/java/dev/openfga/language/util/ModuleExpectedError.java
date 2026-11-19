package dev.openfga.language.util;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import com.fasterxml.jackson.annotation.JsonProperty;
import dev.openfga.language.errors.StartEnd;

@JsonIgnoreProperties(ignoreUnknown = true)
public final class ModuleExpectedError {
    @JsonProperty("msg")
    private String message;

    private String file;
    private StartEnd line;
    private StartEnd column;
    private String type;
    private Metadata metadata;

    @JsonIgnoreProperties(ignoreUnknown = true)
    public static final class Metadata {
        @JsonProperty("errorType")
        private String errorType;

        public String getErrorType() {
            return errorType;
        }

        public void setErrorType(String errorType) {
            this.errorType = errorType;
        }
    }

    /** True when this is a model-validation error, which the Go-equivalent module transformer does not surface. */
    public boolean isValidationError() {
        return metadata != null && metadata.getErrorType() != null;
    }

    public String getMessage() {
        return message;
    }

    public String getFile() {
        return file;
    }

    public StartEnd getLine() {
        return line;
    }

    public StartEnd getColumn() {
        return column;
    }

    public String getType() {
        return type;
    }

    public void setMessage(String message) {
        this.message = message;
    }

    public void setFile(String file) {
        this.file = file;
    }

    public void setLine(StartEnd line) {
        this.line = line;
    }

    public void setColumn(StartEnd column) {
        this.column = column;
    }

    public void setType(String type) {
        this.type = type;
    }

    public void setMetadata(Metadata metadata) {
        this.metadata = metadata;
    }
}
