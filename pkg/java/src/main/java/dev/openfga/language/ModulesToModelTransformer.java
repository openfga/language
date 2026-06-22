package dev.openfga.language;

import com.fasterxml.jackson.core.JsonProcessingException;
import dev.openfga.language.errors.ErrorProperties;
import dev.openfga.language.errors.ModuleTransformationError;
import dev.openfga.language.errors.ModuleTransformationSingleError;
import dev.openfga.language.errors.ParsingError;
import dev.openfga.language.errors.StartEnd;
import dev.openfga.sdk.api.model.AuthorizationModel;
import dev.openfga.sdk.api.model.Condition;
import dev.openfga.sdk.api.model.ConditionMetadata;
import dev.openfga.sdk.api.model.Metadata;
import dev.openfga.sdk.api.model.SourceInfo;
import dev.openfga.sdk.api.model.TypeDefinition;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

/** Merges a set of module files into a single {@link AuthorizationModel}, mirroring the OpenFGA modular DSL. */
public class ModulesToModelTransformer {

    private final List<ParsingError> errors = new ArrayList<>();
    private final List<ModuleFile> moduleFiles;
    private final String schemaVersion;

    public ModulesToModelTransformer(List<ModuleFile> moduleFiles, String schemaVersion) {
        this.moduleFiles = moduleFiles;
        this.schemaVersion = schemaVersion;
    }

    public static String transform(List<ModuleFile> moduleFiles, String schemaVersion)
            throws JsonProcessingException, ModuleTransformationError {
        return JSON.stringify(transformToModel(moduleFiles, schemaVersion));
    }

    public static AuthorizationModel transformToModel(List<ModuleFile> moduleFiles, String schemaVersion)
            throws ModuleTransformationError {
        return new ModulesToModelTransformer(moduleFiles, schemaVersion).transform();
    }

    public AuthorizationModel transform() throws ModuleTransformationError {
        var transformer = new DslToJsonTransformer();

        var typeDefs = new ArrayList<TypeDefinition>();
        var types = new HashSet<String>();
        var conditions = new LinkedHashMap<String, Condition>();
        var extendedTypeDefs = new LinkedHashMap<String, List<TypeDefinition>>();
        var dslByFile = new HashMap<String, ModuleDsl>();

        for (var moduleFile : moduleFiles) {
            var filename = moduleFile.getName();
            var dsl = new ModuleDsl(moduleFile.getContents());
            dslByFile.put(filename, dsl);

            var result = transformer.parseModularDsl(moduleFile.getContents());
            if (result.isFailure()) {
                result.getErrors().forEach(error -> {
                    error.setFile(filename);
                    errors.add(error);
                });
                continue;
            }

            var extensions = result.getTypeDefExtensions();
            for (var typeDef : result.getAuthorizationModel().getTypeDefinitions()) {
                if (extensions.contains(typeDef.getType())) {
                    extendedTypeDefs
                            .computeIfAbsent(filename, k -> new ArrayList<>())
                            .add(typeDef);
                    continue;
                }

                if (types.contains(typeDef.getType())) {
                    addError(
                            "duplicate type definition " + typeDef.getType(),
                            dsl,
                            dsl.getTypeLineNumber(typeDef.getType()),
                            typeDef.getType(),
                            filename);
                    continue;
                }

                if (typeDef.getMetadata() == null) {
                    addError("file is not a module", new StartEnd(0, 0), new StartEnd(0, 0), filename);
                    continue;
                }

                typeDef.getMetadata().setSourceInfo(new SourceInfo()._file(filename));
                types.add(typeDef.getType());
                typeDefs.add(typeDef);
            }

            result.getAuthorizationModel().getConditions().forEach((name, condition) -> {
                if (conditions.containsKey(name)) {
                    addError("duplicate condition " + name, dsl, dsl.getConditionLineNumber(name), name, filename);
                    return;
                }

                var metadata = condition.getMetadata() != null ? condition.getMetadata() : new ConditionMetadata();
                condition.setMetadata(metadata.sourceInfo(new SourceInfo()._file(filename)));
                conditions.put(name, condition);
            });
        }

        mergeExtensions(typeDefs, extendedTypeDefs, dslByFile);

        if (!errors.isEmpty()) {
            throw new ModuleTransformationError(errors);
        }

        return new AuthorizationModel()
                .schemaVersion(schemaVersion)
                .typeDefinitions(typeDefs)
                .conditions(conditions);
    }

    private void mergeExtensions(
            List<TypeDefinition> typeDefs,
            Map<String, List<TypeDefinition>> extendedTypeDefs,
            Map<String, ModuleDsl> dslByFile) {
        extendedTypeDefs.forEach((filename, extensions) -> {
            var dsl = dslByFile.get(filename);

            for (var typeDef : extensions) {
                var original = typeDefs.stream()
                        .filter(t -> t.getType().equals(typeDef.getType()))
                        .findFirst()
                        .orElse(null);

                if (original == null) {
                    addError(
                            "extended type " + typeDef.getType() + " does not exist",
                            dsl,
                            dsl.getExtendedTypeLineNumber(typeDef.getType()),
                            typeDef.getType(),
                            filename);
                    continue;
                }

                if (original.getMetadata() == null) {
                    original.setMetadata(new Metadata());
                }

                var existingRelations = original.getRelations();
                if (existingRelations == null || existingRelations.isEmpty()) {
                    original.setRelations(typeDef.getRelations());
                    var relations = typeDef.getMetadata().getRelations();
                    original.getMetadata().setRelations(relations);
                    if (relations != null) {
                        relations.values().forEach(meta -> meta.setSourceInfo(new SourceInfo()._file(filename)));
                    }
                    continue;
                }

                typeDef.getRelations().forEach((name, relation) -> {
                    if (existingRelations.containsKey(name)) {
                        addError(
                                "relation " + name + " already exists on type " + typeDef.getType(),
                                dsl,
                                dsl.getRelationLineNumber(name),
                                name,
                                filename);
                        return;
                    }

                    var relationMeta = typeDef.getMetadata().getRelations().get(name);
                    relationMeta.setSourceInfo(new SourceInfo()._file(filename));
                    original.putRelationsItem(name, relation);
                    original.getMetadata().putRelationsItem(name, relationMeta);
                });
            }
        });
    }

    private void addError(String message, ModuleDsl dsl, int lineIndex, String symbol, String file) {
        var position = dsl.resolve(lineIndex, symbol);
        addError(message, position[0], position[1], file);
    }

    private void addError(String message, StartEnd line, StartEnd column, String file) {
        errors.add(new ModuleTransformationSingleError(new ErrorProperties(line, column, message), file));
    }
}
