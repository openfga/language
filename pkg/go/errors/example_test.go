package errors_test

import (
	"errors"
	"fmt"

	fgaerrors "github.com/openfga/language/pkg/go/errors"
)

// Example_reasonAndScope shows how a caller recovers the sentinel with errors.Is
// and the scope with errors.As.
func Example_reasonAndScope() {
	// A validator raises the sentinel wrapped in the scope it was raised at.
	err := error(&fgaerrors.ErrRelation{
		ObjectType: "document",
		Relation:   "viewer",
		Cause:      fgaerrors.ErrNoEntrypoints,
	})

	// The sentinel, through however many layers of wrapping.
	if errors.Is(err, fgaerrors.ErrNoEntrypoints) {
		fmt.Println("reason: no entrypoints")
	}

	// The scope, with the fields that scope declares.
	var relationErr *fgaerrors.ErrRelation
	if errors.As(err, &relationErr) {
		fmt.Printf("scope: %s#%s\n", relationErr.ObjectType, relationErr.Relation)
	}

	// A different scope does not match.
	var conditionErr *fgaerrors.ErrCondition
	fmt.Println("condition scope:", errors.As(err, &conditionErr))

	// Output:
	// reason: no entrypoints
	// scope: document#viewer
	// condition scope: false
}

// ExampleSeverity_Blocks shows that only one severity makes a model invalid.
func ExampleSeverity_Blocks() {
	for _, severity := range []fgaerrors.Severity{
		fgaerrors.SeverityError,
		fgaerrors.SeverityWarning,
		fgaerrors.SeverityAdvisory,
	} {
		fmt.Printf("%s blocks: %t\n", severity, severity.Blocks())
	}

	// Output:
	// error blocks: true
	// warning blocks: false
	// advisory blocks: false
}
