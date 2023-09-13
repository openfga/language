import OpenFGALexer from "../gen/OpenFGALexer";
import OpenFGAParser, { TypeNameContext, RelationNameContext, MainContext, TypeDefContext } from "../gen/OpenFGAParser";
import OpenFGAParserListener from "../gen/OpenFGAParserListener";
import OpenFGAVisitor from "../gen/OpenFGAParserVisitor";
import { CharStream, CommonTokenStream, ErrorListener, InputStream, ParseTree, ParseTreeListener, ParseTreeWalker, Parser, ParserRuleContext, RecognitionException, Recognizer, TerminalNode, Token } from "antlr4";

interface SymbolMap {
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

  visitTerminal(node: TerminalNode): void {
    if (!(node.parentCtx instanceof MainContext)) return;
  }
}

class OpenFgaSuggestionVisitor extends OpenFGAVisitor<void> {

}


class SuggestionErrorListiner<T> extends ErrorListener<T> {
  public suggestions: Set<(string|null)> = new Set();

  syntaxError(
    recognizer: Recognizer<T>,
    offendingSymbol: T,
    line: number,
    column: number,
    msg: string, e: RecognitionException | undefined): void {
    const parser = (recognizer as Parser);

    // For some tokens we'll need to check the error text to compare for matches to return
    // console.log((offendingSymbol as Token).text)

    // ExpectedTokens need to be converted to strings
    const expecedTokens = parser.getExpectedTokens();
    for(let inter of expecedTokens.intervals) {
      let s = OpenFGAParser.literalNames[inter.start];
      if(s) {
        this.suggestions.add(s.replace(/\'/g, ''));
      } else {
        this.suggestions.add(OpenFGAParser.symbolicNames[inter.start]);
      }
      
    }
  }
}

export function generateSymbols(dsl: string): SymbolMap {
  const is = new InputStream(dsl);
  const lexer = new OpenFGALexer(is as CharStream);
  lexer.removeErrorListeners();
  const stream = new CommonTokenStream(lexer);

  // Create the Parser
  const parser = new OpenFGAParser(stream);
  parser.removeErrorListeners();
  const parserContext = parser.main();
  // Finally parse the expression
  const visitor = new OpenFgaDslVisitor();

  visitor.visit(parserContext)
  
  return visitor.suggestions;
}

export function getSuggestions(dsl: string, line: number, column: number): Set<(string|null)> {

  // Generate all symbols - need to rethink usage.
  // const symbolMap = generateSymbols(dsl);

  const trimmedDsl = dsl.substring(0, findOffset(dsl, line, column));

  const is = new InputStream(trimmedDsl);
  const lexer = new OpenFGALexer(is as CharStream);
  const listener = new SuggestionErrorListiner();

  lexer.removeErrorListeners()
  lexer.addErrorListener(listener)
  const stream = new CommonTokenStream(lexer);

  // Create the Parser
  const parser = new OpenFGAParser(stream);
  parser.removeErrorListeners()
  parser.addErrorListener(listener)
  const parserContext = parser.main();
  // Finally parse the expression
  const visitor = new OpenFgaSuggestionVisitor();
  visitor.visit(parserContext)

  return listener.suggestions;
}

function findOffset(fileText: string, line: number, column: number): number {
  // we count our current line and column position
  let currentCol = 0;
  let currentLine = 1;
  let offset = 0;

  for (let ch of fileText.split('')) {
     // see if we found where we wanted to go to
     if (currentLine === line && currentCol === column) {
        return offset;
     }

     // line break - increment the line counter and reset the column
     if (ch === "\n") {
        currentLine++;
        currentCol = 0;
     } else {
         currentCol++;
     }

     offset++;
   }
   return -1; //not found
}