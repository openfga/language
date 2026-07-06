export const getConditionLineNumber = (conditionName: string, lines?: string[], skipIndex?: number) => {
  if (!skipIndex || skipIndex < 0) {
    skipIndex = 0;
  }
  if (!lines) {
    return undefined;
  }
  const index = lines
    .slice(skipIndex)
    .findIndex((line: string) => line.trim().startsWith(`condition ${conditionName}`));
  return index === -1 ? -1 : index + skipIndex;
};

export const getTypeLineNumber = (typeName: string, lines?: string[], skipIndex?: number, extension = false) => {
  if (!skipIndex || skipIndex < 0) {
    skipIndex = 0;
  }
  if (!lines) {
    return undefined;
  }
  // Allow an optional trailing comment (e.g. `type page # module: ...`) after the type name.
  const index = lines
    .slice(skipIndex)
    .findIndex((line: string) => line.trim().match(`^${extension ? "extend " : ""}type ${typeName}\\s*(#.*)?$`));
  return index === -1 ? -1 : index + skipIndex;
};

export const getRelationLineNumber = (relation: string, lines?: string[], skipIndex?: number) => {
  if (!skipIndex || skipIndex < 0) {
    skipIndex = 0;
  }
  if (!lines) {
    return undefined;
  }
  const index = lines
    .slice(skipIndex)
    .findIndex((line: string) => line.trim().replace(/ {2,}/g, " ").match(`^define ${relation}\\s*:`));
  return index === -1 ? -1 : index + skipIndex;
};
