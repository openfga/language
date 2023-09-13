import OpenFGALexer from "../gen/OpenFGALexer";
import OpenFGAParser, { TypeNameContext, RelationNameContext, MainContext } from "../gen/OpenFGAParser";
import OpenFGAVisitor from "../gen/OpenFGAParserVisitor";
import * as antlr from "antlr4";

export interface SymbolMap {
  typeNames: Record<string, Set<string>>;
  restrictions: Set<string>;
  literals: Set<string>;
  operators: Set<string>;
}

class OpenFgaDslVisitor extends OpenFGAVisitor<void> {
  public suggestions: SymbolMap;
  private currentType: string | undefined;
  private currentRelation: string | undefined;

  constructor() {
    super();
    this.suggestions = {
      typeNames: {},
      restrictions: new Set(),
      literals: new Set(["model", "schema", "type", "relations", "define"]),
      operators: new Set(["and", "or", "but not", "from"]),
    };
    this.currentType = undefined;
    this.currentRelation = undefined;
  }

  visitTypeName = (ctx: TypeNameContext): void => {
    this.currentType = ctx.getText();
    this.suggestions.typeNames[this.currentType] = new Set();
    this.suggestions.restrictions.add(this.currentType).add(this.currentType + ":*");
  };

  visitRelationName = (ctx: RelationNameContext): void => {
    // Should never exit early
    if (!this.currentType) return;
    this.currentRelation = ctx.getText();

    this.suggestions.typeNames[this.currentType].add(this.currentRelation);

    this.suggestions.restrictions.add(this.currentType + "#" + this.currentRelation);
  };

  visitTerminal(node: antlr.TerminalNode): void {
    if (!(node.parentCtx instanceof MainContext)) return;
  }
}

export function generateSymbols(dsl: string): SymbolMap {
  const is = new antlr.InputStream(dsl);
  const lexer = new OpenFGALexer(is as antlr.CharStream);
  const stream = new antlr.CommonTokenStream(lexer);

  // Create the Parser
  const parser = new OpenFGAParser(stream);
  const parserContext = parser.main();
  // Finally parse the expression
  const visitor = new OpenFgaDslVisitor();
  // @ts-ignore
  visitor.visit(parserContext);
  return visitor.suggestions;
}
