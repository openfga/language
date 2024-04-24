package dev.openfga.language;

import com.fasterxml.jackson.core.JsonProcessingException;
import dev.openfga.language.DslToJsonTransformer.ModularResult;
import dev.openfga.language.errors.ErrorProperties;
import dev.openfga.language.errors.ModuleTransformationError;
import dev.openfga.language.errors.ModuleTransformationSingleError;
import dev.openfga.language.errors.ParsingError;
import dev.openfga.language.errors.StartEnd;
import dev.openfga.language.errors.SyntaxError;
import dev.openfga.language.errors.ValidationMetadata;
import dev.openfga.language.validation.Dsl;
import dev.openfga.sdk.api.model.AuthorizationModel;
import dev.openfga.sdk.api.model.Condition;
import dev.openfga.sdk.api.model.ConditionMetadata;
import dev.openfga.sdk.api.model.Metadata;
import dev.openfga.sdk.api.model.RelationMetadata;
import dev.openfga.sdk.api.model.SourceInfo;
import dev.openfga.sdk.api.model.TypeDefinition;
import dev.openfga.sdk.api.model.Userset;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Set;
import java.util.regex.Pattern;

public class ModulesToModelTransformer {

    private final List<Object> errors = new ArrayList<Object>();
    private final ArrayList<ModuleFile> moduleFiles;
    private final String schemaVersion;

    public ModulesToModelTransformer(ArrayList<ModuleFile> moduleFiles, String schemaVersion) {
        this.moduleFiles = moduleFiles;
        this.schemaVersion = schemaVersion;
    }

    public static String transform(ArrayList<ModuleFile> moduleFiles, String schemaVersion)
            throws JsonProcessingException, ModuleTransformationError {
        return JSON.stringify(transformToModel(moduleFiles, schemaVersion));
    }

    public static AuthorizationModel transformToModel(ArrayList<ModuleFile> moduleFiles, String schemaVersion) throws ModuleTransformationError {
        Result result = new ModulesToModelTransformer(moduleFiles, schemaVersion).transform();
        return result.getAuthorizationModel();
    }

    public Result transform() throws ModuleTransformationError {
        DslToJsonTransformer transformer = new DslToJsonTransformer();

        Map<String, List<TypeDefinition>> extendedTypeDefs = new HashMap<>();
        List<TypeDefinition> typeDefs = new ArrayList<>();
        Map<String, Condition> conditions = new HashMap<>();
        Set<String> types = new HashSet<>();
        Map<String, Dsl> dslMap = new HashMap<>();

        for (ModuleFile moduleFile : moduleFiles) {
            String filename = moduleFile.getName();
            String contents = moduleFile.getContents();

            Dsl dsl = new Dsl(contents.split("\n"));
            dslMap.put(filename, dsl);

            ModularResult result = transformer.parseModularDsl(contents);
            if (result.IsFailure()) {
                for (SyntaxError error : result.getErrors()) {
                    error.setFile(filename);
                    errors.add(error);
                }
                continue;
            }
            
            AuthorizationModel model = result.getAuthorizationModel();

            for (TypeDefinition typeDef : model.getTypeDefinitions()) {
                if (result.getExtendedTypeDefinitions().containsKey(typeDef.getType())) {
                    // extendedTypeDefs.get(filename).add(typeDef);
                    extendedTypeDefs.computeIfAbsent(filename, k -> new ArrayList<>()).add(typeDef);
                    continue;
                }

                Metadata meta = typeDef.getMetadata();
                if (meta == null) {
                    meta = new Metadata();
                }
                typeDef.setMetadata(meta.sourceInfo(new SourceInfo()._file(filename)));
                types.add(typeDef.getType());
                typeDefs.add(typeDef);
            }

            if (!model.getConditions().isEmpty()) {
                for (Map.Entry<String, Condition> conditionEntry :
                        model.getConditions().entrySet()) {
                    String conditionName = conditionEntry.getKey();
                    Condition condition = conditionEntry.getValue();

                    if (conditions.containsKey(conditionName)) {
                        int lineIndex = dsl.getConditionLineNumber(conditionName);

                        addError("duplicate condition " + conditionName, dsl, lineIndex, conditionName);
                        continue;
                    }

                    ConditionMetadata meta = condition.getMetadata();
                    if (meta == null) {
                        meta = new ConditionMetadata();
                    }
                    SourceInfo sourceInfo = meta.getSourceInfo();
                    if (sourceInfo == null) {
                        sourceInfo = new SourceInfo();
                    }

                    condition.setMetadata(meta.sourceInfo(sourceInfo._file(filename)));
                    conditions.put(conditionName, condition);
                }
            }
        }

        for (Map.Entry<String, List<TypeDefinition>> entry : extendedTypeDefs.entrySet()) {
            String filename = entry.getKey();
            Dsl dsl = dslMap.get(filename);

            for (TypeDefinition typeDef : entry.getValue()) {
                if (typeDef.getRelations().isEmpty()) {
                    // TODO
                    continue;
                }

                int originalIndex = 0;
                TypeDefinition original = null;
                for (int i = 0; i < typeDefs.size(); i++) {
                    TypeDefinition t = typeDefs.get(i);
                    if (t.getType().equals(typeDef.getType())) {
                        originalIndex = i;
                        original = t;
                        break;
                    }
                }

                if (original == null) {
                    int lineIndex = dsl.getTypeLineNumber(typeDef.getType());
                    addError("extended type " + typeDef.getType() + " does not exist", dsl, lineIndex, typeDef.getType());
                    continue;
                }

                Set<String> existingRelationNames = original.getRelations().keySet();

                if (existingRelationNames.isEmpty()) {
                    original.setRelations(typeDef.getRelations());
                    original.getMetadata().setRelations(typeDef.getMetadata().getRelations());

                    for (Map.Entry<String, RelationMetadata> metaEntry :
                            original.getMetadata().getRelations().entrySet()) {
                        RelationMetadata value = metaEntry.getValue();
                        value.setSourceInfo(new SourceInfo()._file(filename));
                        original.getMetadata().putRelationsItem(metaEntry.getKey(), value);
                    }

                    typeDefs.set(originalIndex, original);
                    continue;
                }

                for (Map.Entry<String, Userset> relationsEntry :
                typeDef.getRelations().entrySet()) {
                    String name = relationsEntry.getKey();
                    Userset relation = relationsEntry.getValue();
                    if (existingRelationNames.contains(name)) {
                        var typeIndex = dsl.getExtendedTypeLineNumber(typeDef.getType());
                        int lineIndex = dsl.getRelationLineNumber(name, typeIndex);

                        addError("relation " + name + " already exists on type " + typeDef.getType(), dsl, lineIndex, name);
                        continue;
                    }

                    RelationMetadata relationsMetadata =
                            typeDef.getMetadata().getRelations().get(name);

                    SourceInfo sourceInfo = relationsMetadata.getSourceInfo();
                    if (sourceInfo == null) {
                        sourceInfo = new SourceInfo();
                    }

                    relationsMetadata.setSourceInfo(sourceInfo._file(filename));

                    original.putRelationsItem(name, relation);
                    original.getMetadata().putRelationsItem(name, relationsMetadata);
                }
                typeDefs.set(originalIndex, original);
            }
        }

        if (!errors.isEmpty()) {
            throw new ModuleTransformationError(errors);
        }

        return new Result(new AuthorizationModel()
                .schemaVersion(schemaVersion)
                .typeDefinitions(typeDefs)
                .conditions(conditions));
    }

    private void addError(String message, Dsl dsl, int lineIndex, String symbol) {

        String rawLine = dsl.getLine(lineIndex);
        var regex = Pattern.compile("\\b" + symbol + "\\b");
        var wordIdx = 0;
        var matcher = regex.matcher(rawLine);
        if (matcher.find()) {
            wordIdx = matcher.start();
        }

        StartEnd line = new StartEnd(lineIndex, lineIndex);
        StartEnd column = new StartEnd(wordIdx, wordIdx + symbol.length());

        ErrorProperties properties = new ErrorProperties(line, column, message);
        ValidationMetadata metadata = new ValidationMetadata();

        errors.add(new ModuleTransformationSingleError(properties, metadata));
    }

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

    public static final class Result {
        private final AuthorizationModel authorizationModel;

        public Result(AuthorizationModel authorizationModel) {
            this.authorizationModel = authorizationModel;
        }

        public AuthorizationModel getAuthorizationModel() {
            return authorizationModel;
        }
    }
}
