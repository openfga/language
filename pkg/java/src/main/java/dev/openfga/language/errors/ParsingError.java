package dev.openfga.language.errors;

import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import java.util.Objects;

@Getter
@Setter
@EqualsAndHashCode(callSuper = true)
@NoArgsConstructor
public abstract class ParsingError extends SimpleError {

    private StartEnd line;

    private StartEnd column;

    private String fullMessage;

    public ParsingError(String type, ErrorProperties properties) {
        super(properties.getMessage());
        line = properties.getLine();
        column = properties.getColumn();
        fullMessage = properties.getFullMessage(type);
    }

    public StartEnd getLine() {
        return getLine(0);
    }

    public StartEnd getLine(int offset) {
        return line.withOffset(offset);
    }

    public StartEnd getColumn() {
        return getColumn(0);
    }

    public StartEnd getColumn(int offset) {
        return column.withOffset(offset);
    }

    public String toString() {
        return Objects.requireNonNullElseGet(fullMessage, this::getMessage);
    }
}