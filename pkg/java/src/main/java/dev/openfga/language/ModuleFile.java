package dev.openfga.language;

/** A single file that makes up a modular model, identified by its name and DSL contents. */
public final class ModuleFile {
    private final String name;
    private final String contents;

    public ModuleFile(String name, String contents) {
        this.name = name;
        this.contents = contents;
    }

    public String getName() {
        return name;
    }

    public String getContents() {
        return contents;
    }
}
