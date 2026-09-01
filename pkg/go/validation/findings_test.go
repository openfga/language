package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindingError(t *testing.T) {
	t.Parallel()

	t.Run("with position", func(t *testing.T) {
		t.Parallel()

		finding := &Finding{
			Message: "the relation `viewer` does not exist.",
			Line:    &Range{Start: 4, End: 4},
			Column:  &Range{Start: 12, End: 18},
		}

		assert.Equal(t, "validation error at line=4, column=12: the relation `viewer` does not exist.", finding.Error())
	})

	t.Run("without position", func(t *testing.T) {
		t.Parallel()

		finding := &Finding{Message: "schema version required"}

		assert.Equal(t, "validation error: schema version required", finding.Error())
	})
}

func TestFindingsError(t *testing.T) {
	t.Parallel()

	t.Run("empty", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, "no validation errors", Findings{}.Error())
	})

	t.Run("one finding", func(t *testing.T) {
		t.Parallel()

		findings := Findings{{Message: "first"}}

		assert.Equal(t, "1 error occurred:\n\t* validation error: first\n\n", findings.Error())
	})

	t.Run("two findings pluralize", func(t *testing.T) {
		t.Parallel()

		findings := Findings{{Message: "first"}, {Message: "second"}}

		assert.Equal(t, "2 errors occurred:\n\t* validation error: first\n\t* validation error: second\n\n",
			findings.Error())
	})
}

func TestFindingsErr(t *testing.T) {
	t.Parallel()

	t.Run("nil for no findings", func(t *testing.T) {
		t.Parallel()

		require.NoError(t, Findings(nil).Err())
		require.NoError(t, Findings{}.Err())
	})

	t.Run("the collection itself otherwise", func(t *testing.T) {
		t.Parallel()

		findings := Findings{{Message: "boom"}}
		err := findings.Err()

		require.Error(t, err)

		var recovered Findings
		require.ErrorAs(t, err, &recovered)
		assert.Len(t, recovered, 1)
	})
}

func TestFindingsUnwrap(t *testing.T) {
	t.Parallel()

	first := &Finding{Message: "first", Metadata: Metadata{Kind: InvalidName}}
	second := &Finding{Message: "second", Metadata: Metadata{Kind: DuplicatedError}}
	err := Findings{first, second}.Err()

	// errors.As walks Unwrap() []error and stops at the first finding.
	var finding *Finding
	require.ErrorAs(t, err, &finding)
	assert.Same(t, first, finding)

	require.ErrorIs(t, err, error(first))
	require.ErrorIs(t, err, error(second))
}

func TestFindingsAdd(t *testing.T) {
	t.Parallel()

	var findings Findings

	findings = findings.add(nil)
	assert.Empty(t, findings, "a nil finding is nothing found")

	findings = findings.add(&Finding{Message: "found"})
	assert.Len(t, findings, 1)
}

func TestFindingIn(t *testing.T) {
	t.Parallel()

	t.Run("stamps file and module", func(t *testing.T) {
		t.Parallel()

		finding := (&Finding{}).in("core.fga", "core")

		assert.Equal(t, "core.fga", finding.File)
		assert.Equal(t, "core", finding.Metadata.Module)
	})

	t.Run("nil finding stays nil", func(t *testing.T) {
		t.Parallel()

		var finding *Finding
		assert.Nil(t, finding.in("core.fga", "core"))
	})
}

// TestFindingsAsError pins the boundary contract: a validation error is always
// a Findings, and errors.As is the documented way back to the findings.
func TestFindingsAsError(t *testing.T) {
	t.Parallel()

	err := Findings{{Message: "boom", Metadata: Metadata{Kind: InvalidName, Symbol: "x"}}}.Err()

	var findings Findings
	require.ErrorAs(t, err, &findings)
	assert.Equal(t, InvalidName, findings[0].Metadata.Kind)
}
