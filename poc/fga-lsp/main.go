// fgalsp is a proof-of-concept language server for the OpenFGA DSL: stdio
// JSON-RPC, full-document sync, diagnostics from graph-based validation.
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/openfga/language/pkg/go/transformer"
	"github.com/openfga/language/pkg/go/validation"
)

type request struct {
	ID     *json.RawMessage `json:"id,omitempty"`
	Method string           `json:"method"`
	Params json.RawMessage  `json:"params,omitempty"`
}

type position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

type lspRange struct {
	Start position `json:"start"`
	End   position `json:"end"`
}

type diagnostic struct {
	Range    lspRange `json:"range"`
	Severity int      `json:"severity"`
	Code     string   `json:"code,omitempty"`
	Source   string   `json:"source"`
	Message  string   `json:"message"`
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	for {
		msg, err := readFrame(in)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "read:", err)

			return
		}

		var req request
		if err := json.Unmarshal(msg, &req); err != nil {
			continue
		}

		switch req.Method {
		case "initialize":
			reply(out, req.ID, map[string]any{
				"capabilities": map[string]any{
					"textDocumentSync": 1, // full
				},
				"serverInfo": map[string]any{"name": "fga-lsp-poc"},
			})
		case "textDocument/didOpen":
			var p struct {
				TextDocument struct {
					URI  string `json:"uri"`
					Text string `json:"text"`
				} `json:"textDocument"`
			}
			if json.Unmarshal(req.Params, &p) == nil {
				publish(out, p.TextDocument.URI, diagnose(p.TextDocument.Text))
			}
		case "textDocument/didChange":
			var p struct {
				TextDocument struct {
					URI string `json:"uri"`
				} `json:"textDocument"`
				ContentChanges []struct {
					Text string `json:"text"`
				} `json:"contentChanges"`
			}
			if json.Unmarshal(req.Params, &p) == nil && len(p.ContentChanges) > 0 {
				publish(out, p.TextDocument.URI, diagnose(p.ContentChanges[len(p.ContentChanges)-1].Text))
			}
		case "shutdown":
			reply(out, req.ID, nil)
		case "exit":
			return
		}
	}
}

// diagnose turns a DSL into LSP diagnostics: transformer syntax errors or
// graph-validation findings. Both count positions from zero, like LSP.
func diagnose(dsl string) []diagnostic {
	diagnostics := []diagnostic{}

	model, err := transformer.TransformDSLToProto(dsl)
	if err != nil {
		return append(diagnostics, diagnostic{
			Range:    lspRange{Start: position{0, 0}, End: position{0, 1}},
			Severity: 1,
			Code:     "syntax",
			Source:   "openfga",
			Message:  err.Error(),
		})
	}

	for _, finding := range validation.ExtractAllAs[*validation.Finding](
		validation.ValidateDSLWithGraph(model, dsl)) {
		d := diagnostic{
			Severity: 1,
			Code:     string(finding.Metadata.Kind),
			Source:   "openfga",
			Message:  finding.Message,
		}
		if finding.Line != nil && finding.Column != nil {
			d.Range = lspRange{
				Start: position{finding.Line.Start, finding.Column.Start},
				End:   position{finding.Line.End, finding.Column.End},
			}
		}

		diagnostics = append(diagnostics, d)
	}

	return diagnostics
}

func publish(out *bufio.Writer, uri string, diagnostics []diagnostic) {
	notify(out, "textDocument/publishDiagnostics", map[string]any{
		"uri":         uri,
		"diagnostics": diagnostics,
	})
}

func readFrame(in *bufio.Reader) ([]byte, error) {
	length := 0

	for {
		line, err := in.ReadString('\n')
		if err != nil {
			return nil, err
		}
		if line == "\r\n" {
			break
		}

		fmt.Sscanf(line, "Content-Length: %d", &length)
	}

	body := make([]byte, length)
	_, err := io.ReadFull(in, body)

	return body, err
}

func send(out *bufio.Writer, payload any) {
	body, _ := json.Marshal(payload)
	fmt.Fprintf(out, "Content-Length: %d\r\n\r\n%s", len(body), body)
	out.Flush()
}

func reply(out *bufio.Writer, id *json.RawMessage, result any) {
	send(out, map[string]any{"jsonrpc": "2.0", "id": id, "result": result})
}

func notify(out *bufio.Writer, method string, params any) {
	send(out, map[string]any{"jsonrpc": "2.0", "method": method, "params": params})
}
