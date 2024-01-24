package dev.openfga.language.errors;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
public class ErrorProperties {

    private StartEnd line;

    private StartEnd column;

    private String message;
    String getFullMessage(String type) {
        return String.format("%s error at line=%d, column=%d: %s", type, line.getStart(), column.getStart(), message);
    }
}