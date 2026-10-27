// Fold runs of space, tab and form feed — the separators the lexer's WHITESPACE
// rule admits — so `define\towner:` is found like `define owner:`.
const normalizeWhitespace = (line: string) => line.trim().replace(/[ \t\f]+/g, " ");

export const getConditionLineNumber = (conditionName: string, lines?: string[], skipIndex?: number) => {
  if (!skipIndex || skipIndex < 0) {
    skipIndex = 0;
  }
  if (!lines) {
    return undefined;
  }
  // Require `(` after the name so a condition name that is a prefix of another
  // (e.g. `less` vs `less_than`) cannot match the wrong line.
  const conditionPrefix = `condition ${conditionName}`;
  const index = lines.slice(skipIndex).findIndex((line: string) => {
    const normalized = normalizeWhitespace(line);
    return normalized.startsWith(conditionPrefix) && /^\s*\(/.test(normalized.slice(conditionPrefix.length));
  });
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
  // The comment must be preceded by whitespace so a `#` glued to the name isn't treated as a comment.
  // Match the type name literally (it may contain regex metacharacters like `.`).
  const typePrefix = `${extension ? "extend " : ""}type ${typeName}`;
  const index = lines.slice(skipIndex).findIndex((line: string) => {
    const normalized = normalizeWhitespace(line);
    return normalized.startsWith(typePrefix) && /^(\s+#.*)?$/.test(normalized.slice(typePrefix.length));
  });
  return index === -1 ? -1 : index + skipIndex;
};

export const getRelationLineNumber = (relation: string, lines?: string[], skipIndex?: number) => {
  if (!skipIndex || skipIndex < 0) {
    skipIndex = 0;
  }
  if (!lines) {
    return undefined;
  }
  // Match the relation name literally (it may contain regex metacharacters like `.`).
  const relationPrefix = `define ${relation}`;
  const index = lines.slice(skipIndex).findIndex((line: string) => {
    const normalized = normalizeWhitespace(line);
    return normalized.startsWith(relationPrefix) && /^\s*:/.test(normalized.slice(relationPrefix.length));
  });
  return index === -1 ? -1 : index + skipIndex;
};
