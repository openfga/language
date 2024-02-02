package dev.openfga.language.errors;

public class ErrorProperties {

    private StartEnd line;

    private StartEnd column;

    private String message;

    public ErrorProperties(StartEnd line, StartEnd column, String message) {
        this.line = line;
        this.column = column;
        this.message = message;
    }

    String getFullMessage(String type) {
        return String.format("%s error at line=%d, column=%d: %s", type, line.getStart(), column.getStart(), message);
    }

    public StartEnd getLine() {
        return line;
    }

    public void setLine(StartEnd line) {
        this.line = line;
    }

    public StartEnd getColumn() {
        return column;
    }

    public void setColumn(StartEnd column) {
        this.column = column;
    }

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }
}