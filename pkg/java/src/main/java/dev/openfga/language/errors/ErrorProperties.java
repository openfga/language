package dev.openfga.language.errors;

public class ErrorProperties {

    private StartEnd line;

    private StartEnd column;

    private String file;

    private String module;

    private String message;

    public ErrorProperties(StartEnd line, StartEnd column, String message) {
        this.line = line;
        this.column = column;
        this.message = message;
    }

    public ErrorProperties(StartEnd line, StartEnd column, String message, String file, String module) {
        this.line = line;
        this.column = column;
        this.message = message;
        this.file = file;
        this.module = module;
    }

    String getFullMessage(String type) {
        if (line != null && column != null) {
            return String.format(
                    "%s error at line=%d, column=%d: %s", type, line.getStart(), column.getStart(), message);
        } else {
            return String.format("%s error: %s", type, message);
        }
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

    public String getFile() {
        return file;
    }

    public void setFile(String file) {
        this.file = file;
    }

    public String getModule() {
        return module;
    }

    public void setModule(String module) {
        this.module = module;
    }

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }
}
