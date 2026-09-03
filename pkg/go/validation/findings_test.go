package validation

import (
	"errors"
	"fmt"
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

func TestJoinFindings(t *testing.T) {
	t.Parallel()

	t.Run("nil when there is nothing to report", func(t *testing.T) {
		t.Parallel()

		require.NoError(t, joinFindings())
		require.NoError(t, joinFindings(nil, nil))
	})

	t.Run("drops nil findings", func(t *testing.T) {
		t.Parallel()

		err := joinFindings(nil, &Finding{Message: "boom"}, nil)
		require.Error(t, err)

		found := ExtractAllAs[*Finding](err)
		require.Len(t, found, 1)
		assert.Equal(t, "boom", found[0].Message)
	})

	t.Run("keeps the order it is given", func(t *testing.T) {
		t.Parallel()

		err := joinFindings(&Finding{Message: "first"}, &Finding{Message: "second"})

		found := ExtractAllAs[*Finding](err)
		require.Len(t, found, 2)
		assert.Equal(t, "first", found[0].Message)
		assert.Equal(t, "second", found[1].Message)
	})
}

func TestExtractAllAs(t *testing.T) {
	t.Parallel()

	t.Run("nil error yields nothing", func(t *testing.T) {
		t.Parallel()

		assert.Empty(t, ExtractAllAs[*Finding](nil))
	})

	t.Run("recovers findings from a nested tree in order", func(t *testing.T) {
		t.Parallel()

		// A join of joins, the shape validate builds from its phases.
		left := joinFindings(&Finding{Message: "1"}, &Finding{Message: "2"})
		right := joinFindings(&Finding{Message: "3"}, &Finding{Message: "4"})

		found := ExtractAllAs[*Finding](errors.Join(left, right))
		require.Len(t, found, 4)
		assert.Equal(t, []string{"1", "2", "3", "4"}, []string{
			found[0].Message, found[1].Message, found[2].Message, found[3].Message,
		})
	})

	t.Run("terminates on a leaf that is neither the target nor a wrapper", func(t *testing.T) {
		t.Parallel()

		// errors.New has neither an Unwrap() []error nor an Unwrap() error, so
		// the walk has nothing to descend into and ends its branch there.
		found := ExtractAllAs[*Finding](errors.Join(&Finding{Message: "boom"}, errors.New("unrelated")))
		require.Len(t, found, 1)
		assert.Equal(t, "boom", found[0].Message)
	})

	t.Run("descends through a single-error wrapper", func(t *testing.T) {
		t.Parallel()

		found := ExtractAllAs[*Finding](fmt.Errorf("context: %w", &Finding{Message: "wrapped"}))
		require.Len(t, found, 1)
		assert.Equal(t, "wrapped", found[0].Message)
	})
}

// TestFindingErrorsInterop pins that the standard errors helpers still reach a
// finding, so a caller that only wants the first is not forced through
// ExtractAllAs.
func TestFindingErrorsInterop(t *testing.T) {
	t.Parallel()

	first := &Finding{Message: "first", Metadata: Metadata{Kind: InvalidName}}
	second := &Finding{Message: "second", Metadata: Metadata{Kind: DuplicatedError}}
	err := joinFindings(first, second)

	t.Run("errors.As reaches the first finding", func(t *testing.T) {
		t.Parallel()

		var finding *Finding
		require.ErrorAs(t, err, &finding)
		assert.Same(t, first, finding)
	})

	t.Run("errors.Is matches each finding", func(t *testing.T) {
		t.Parallel()

		require.ErrorIs(t, err, error(first))
		require.ErrorIs(t, err, error(second))
	})
}
