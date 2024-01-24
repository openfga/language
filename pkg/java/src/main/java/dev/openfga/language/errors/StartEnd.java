package dev.openfga.language.errors;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public final class StartEnd {
    private int start;
    private int end;

    public StartEnd withOffset(int offset) {
        return new StartEnd(start + offset, end + offset);
    }
}
