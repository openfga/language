package dev.openfga.language.errors;

import java.util.Objects;

public final class StartEnd {
    private int start;
    private int end;

    // Needed for Jackson deserialization
    public StartEnd() {}

    public StartEnd(int start, int end) {
        this.start = start;
        this.end = end;
    }

    public int getStart() {
        return start;
    }

    public void setStart(int start) {
        this.start = start;
    }

    public int getEnd() {
        return end;
    }

    public void setEnd(int end) {
        this.end = end;
    }

    public StartEnd withOffset(int offset) {
        return new StartEnd(start + offset, end + offset);
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        StartEnd startEnd = (StartEnd) o;
        return start == startEnd.start && end == startEnd.end;
    }

    @Override
    public int hashCode() {
        return Objects.hash(start, end);
    }
}
