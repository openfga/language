import { parseDocument, LineCounter, Scalar, isSeq, isNode, YAMLSeq } from "yaml";
import { LinePos } from "yaml/dist/errors";
import { FGAModFileValidationError, FGAModFileValidationSingleError } from "../../errors";
/**
 * An `fga.mod` file represented as JSON
 */
export interface ModFile {
  /**
   * The schema version. This currently will only be 1.2
   */
  schema: string
  /**
   * The module name.
   */
  module: string
  /**
   * The individual files that make up the modular model.
   */
  contents: string[]
}

/**
 * Gets the line and column data from the `linePos` returned from parsing a yaml file.
 * @param linePos - The `linePos` property.
 * @returns Line and column data.
 */
function getLineAndColumnFromLinePos(
  linePos?: [LinePos] | [LinePos, LinePos]
): { line: { start: number; end: number; }, column: { start: number; end: number; }} {
  if (linePos === undefined) {
    return { line: { start: 1, end: 1 }, column: { start: 1, end: 1 } };
  }
  // eslint-disable-next-line prefer-const
  let [start, end] = linePos;
  if (end === undefined) {
    end = start;
  }
  return {
    line: {
      start: start.line ,
      end: end.line
    },
    column: {
      start: start.col,
      end: end.col
    }
  };
}

/**
 * Gets the line and column data for a node in a yaml doc.
 * @param node - The node from the yaml doc.
 * @param counter - The instance of LineCounter that was passed to `parseDocument`.
 * @returns Line and column data.
 */
function getLineAndColumnFromNode(
  node: unknown,
  counter: LineCounter
): { line: { start: number; end: number; }, column: { start: number; end: number; }} {
  if (!isNode(node) || !node.range) {
    return { line: { start: 1, end: 1 }, column: { start: 1, end: 1 } };
  }

  const start = counter.linePos(node.range[0]);
  const end = counter.linePos(node.range[1]);
  return {
    line: {
      start: start.line,
      end: end.line
    },
    column: {
      start: start.col ,
      end: end.col
    }
  };
}

/**
 * Transforms an `fga.mod` file into the JSON representation and validate the fields are correct.
 * @param {string} modFile - The `fga.mod` file
 * @returns {ModFile} The jSON representation of the `fga.mod` file.
 */
export const transformModFileToJSON = (modFile: string): ModFile => {
  const lineCounter = new LineCounter();
  const yamlDoc = parseDocument(modFile, {
    lineCounter,
    keepSourceTokens: true
  });

  const errors: FGAModFileValidationSingleError[] = [];
  // Copy over any general yaml parsing errors
  for (const error of yamlDoc.errors) {
    errors.push(new FGAModFileValidationSingleError({
      msg: error.message,
      ...getLineAndColumnFromLinePos(error.linePos)
    }));
  }

  if (!yamlDoc.has("schema")) {
    errors.push(new FGAModFileValidationSingleError({
      msg: "missing schema field",
      ...getLineAndColumnFromLinePos()
    }));
  } else {
    const schema = yamlDoc.get("schema");
    if (typeof schema !== "string") {
      const node = yamlDoc.getIn(["schema"], true);

      errors.push(new FGAModFileValidationSingleError({
        msg: `unexpected schema type, expected string got value ${schema}`,
        ...getLineAndColumnFromNode(node, lineCounter)
      }));
    } else if (schema !== "1.2") {
      const node = yamlDoc.getIn(["schema"], true);
      errors.push(new FGAModFileValidationSingleError({
        msg: "unsupported schema version, fga.mod only supported in version `1.2`",
        ...getLineAndColumnFromNode(node, lineCounter)
      }));
    }
  }

  if (!yamlDoc.has("module")) {
    errors.push(new FGAModFileValidationSingleError({
      msg: "missing module field",
      ...getLineAndColumnFromLinePos()
    }));
  } else if (typeof yamlDoc.get("module") !== "string") {
    const node = yamlDoc.getIn(["module"], true);
    errors.push(new FGAModFileValidationSingleError({
      msg: `unexpected module type, expected string got value ${yamlDoc.get("module")}`,
      ...getLineAndColumnFromNode(node, lineCounter)
    }));
  }

  if (!yamlDoc.has("contents")) {
    errors.push(new FGAModFileValidationSingleError({
      msg: "missing contents field",
      ...getLineAndColumnFromLinePos()
    }));
  } else if (!isSeq(yamlDoc.get("contents"))) {
    const node = yamlDoc.getIn(["contents"], true);
    const contents = yamlDoc.get("contents");
    errors.push(new FGAModFileValidationSingleError({
      msg: `unexpected contents type, expected list of strings got value ${contents}`,
      ...getLineAndColumnFromNode(node, lineCounter)
    }));
  } else {
    const contents = yamlDoc.get("contents") as YAMLSeq<Scalar>;
    for (const file of contents.items) {
      if (typeof file.value !== "string") {
        errors.push(new FGAModFileValidationSingleError({
          msg: `unexpected contents item type, expected string got value ${file.value}`,
          ...getLineAndColumnFromNode(file, lineCounter)
        }));
        continue;
      }

      if (!file.value.endsWith(".fga")) {
        errors.push(new FGAModFileValidationSingleError({
          msg: `contents items should use fga file extension, got ${file.value}`,
          ...getLineAndColumnFromNode(file, lineCounter)
        }));
      }
    }
  }

  if (errors.length) {
    throw new FGAModFileValidationError(errors);
  }

  return yamlDoc.toJSON();
};
