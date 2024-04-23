package dev.openfga.language;

public class ModulesToModelTransformer {
    
    public static class ModuleFile {
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

}
