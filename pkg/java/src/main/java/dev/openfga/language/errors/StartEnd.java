package dev.openfga.language.errors;

public final class StartEnd {
    private int start;
    private int end;

    // Needed for Jackson deserialization
    public StartEnd() {
    }

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
}
