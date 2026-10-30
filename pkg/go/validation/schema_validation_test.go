package validation

import (
	"testing"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateSchemaVersion(t *testing.T) {
	t.Parallel()

	model := func(version string) *openfgav1.AuthorizationModel {
		return &openfgav1.AuthorizationModel{SchemaVersion: version}
	}

	t.Run("supported versions yield nothing", func(t *testing.T) {
		t.Parallel()

		assert.Empty(t, ExtractAllAs[*Finding](validateSchemaVersion(model("1.1"), source{})))
		assert.Empty(t, ExtractAllAs[*Finding](validateSchemaVersion(model("1.2"), source{})))
	})

	t.Run("missing version is required at line zero", func(t *testing.T) {
		t.Parallel()

		findings := ExtractAllAs[*Finding](validateSchemaVersion(model(""), newSource("model\ntype user")))

		require.Len(t, findings, 1)
		assert.Equal(t, SchemaVersionRequired, findings[0].Metadata.Kind)
		assert.Equal(t, &Range{Start: 0, End: 0}, findings[0].Line)
	})

	t.Run("no source text means no position", func(t *testing.T) {
		t.Parallel()

		findings := ExtractAllAs[*Finding](validateSchemaVersion(model("1.3"), source{}))

		require.Len(t, findings, 1)
		assert.Nil(t, findings[0].Line)
		assert.Nil(t, findings[0].Column)
	})
}
