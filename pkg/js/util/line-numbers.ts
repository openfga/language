export const getConditionLineNumber = (conditionName: string, lines?: string[], skipIndex?: number) => {
  if (!skipIndex) {
    skipIndex = 0;
  }
  if (!lines) {
    return undefined;
  }
  return (
    lines.slice(skipIndex).findIndex((line: string) => line.trim().startsWith(`condition ${conditionName}`)) + skipIndex
  );
};

export const getTypeLineNumber = (typeName: string, lines?: string[], skipIndex?: number, extension = false) => {
  if (!skipIndex) {
    skipIndex = 0;
  }
  if (!lines) {
    return undefined;
  }
  return (
    lines
      .slice(skipIndex)
      .findIndex((line: string) => line.trim().match(`^${extension ? "extend " : ""}type ${typeName}$`)) + skipIndex
  );
};

export const getRelationLineNumber = (relation: string, lines?: string[], skipIndex?: number) => {
  if (!skipIndex) {
    skipIndex = 0;
  }
  if (!lines) {
    return undefined;
  }
  return (
    lines
      .slice(skipIndex)
      .findIndex((line: string) => line.trim().replace(/ {2,}/g, " ").match(`^define ${relation}\\s*:`)) + skipIndex
  );
};
