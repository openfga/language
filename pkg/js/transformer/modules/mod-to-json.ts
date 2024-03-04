import { parseDocument, LineCounter, Scalar, isSeq, isNode, YAMLSeq } from "yaml";
import { LinePos } from "yaml/dist/errors";
import { FGAModFileValidationError, FGAModFileValidationSingleError } from "../../errors";

/**
 * Represents a property contained in the `fga.mod` file, includes the value as well as line an
 * column information to allow improved error reporting.
 */
export interface ModFileProperty<T> {
  /**
   * The value of the property.
   */
  value: T

  /**
   * The start and end line number of the property.
   */
  line: {
    start: number;
    end: number;
  }

  /**
   * The start and end column number of the property.
   */
  column: {
    start: number;
    end: number;
  }
}

/**
 * An `fga.mod` file represented as JSON
 */
export interface ModFile {
  /**
   * The schema version. This currently will only be 1.2
   */
  schema: ModFileProperty<string>
  /**
   * The module name.
   */
  module: ModFileProperty<string>
  /**
   * The individual files that make up the modular model.
   */
  contents: ModFileProperty<ModFileProperty<string>[]>
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

  const parsedModFile: Partial<ModFile> = {};

  const schemaNode = yamlDoc.get("schema", true) as Scalar<string>;
  if (!schemaNode) {
    errors.push(new FGAModFileValidationSingleError({
      msg: "missing schema field",
      ...getLineAndColumnFromLinePos()
    }));
  } else if (typeof schemaNode.value !== "string") {
    errors.push(new FGAModFileValidationSingleError({
      msg: `unexpected schema type, expected string got value ${schemaNode.value}`,
      ...getLineAndColumnFromNode(schemaNode, lineCounter)
    }));
  } else if (schemaNode.value !== "1.2") {
    errors.push(new FGAModFileValidationSingleError({
      msg: "unsupported schema version, fga.mod only supported in version `1.2`",
      ...getLineAndColumnFromNode(schemaNode, lineCounter)
    }));
  } else {
    parsedModFile.schema = {
      value: schemaNode.value,
      ...getLineAndColumnFromNode(schemaNode, lineCounter)
    };
  }

  const moduleNode = yamlDoc.get("module", true) as Scalar<string>;
  if (!moduleNode) {
    errors.push(new FGAModFileValidationSingleError({
      msg: "missing module field",
      ...getLineAndColumnFromLinePos()
    }));
  } else if (typeof moduleNode.value !== "string") {
    errors.push(new FGAModFileValidationSingleError({
      msg: `unexpected module type, expected string got value ${moduleNode.value}`,
      ...getLineAndColumnFromNode(moduleNode, lineCounter)
    }));
  } else {
    parsedModFile.module = {
      value: moduleNode.value,
      ...getLineAndColumnFromNode(moduleNode, lineCounter)
    };
  }

  const contentsNode = yamlDoc.get("contents", true) as YAMLSeq<Scalar>;
  if (!contentsNode) {
    errors.push(new FGAModFileValidationSingleError({
      msg: "missing contents field",
      ...getLineAndColumnFromLinePos()
    }));
  } else if (!isSeq(contentsNode)) {
    const node = yamlDoc.get("contents", true);
    const contents = yamlDoc.get("contents");
    errors.push(new FGAModFileValidationSingleError({
      msg: `unexpected contents type, expected list of strings got value ${contents}`,
      ...getLineAndColumnFromNode(node, lineCounter)
    }));
  } else {
    const contents = yamlDoc.get("contents") as YAMLSeq<Scalar>;
    const contentsValue = [];
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
        continue;
      }

      contentsValue.push({
        value: file.value,
        ...getLineAndColumnFromNode(file, lineCounter)
      });
    }
    const node = yamlDoc.get("contents", true);
    parsedModFile.contents = {
      value: contentsValue,
      ...getLineAndColumnFromNode(node, lineCounter)
    };

  }

  if (errors.length) {
    throw new FGAModFileValidationError(errors);
  }

  return parsedModFile as ModFile;
};
