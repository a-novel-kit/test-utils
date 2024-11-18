package testutils_test

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/require"

	testutils "github.com/a-novel-kit/test-utils"
	testutilsmocks "github.com/a-novel-kit/test-utils/mocks"
)

func TestUnsafe(t *testing.T) {
	src := testutilsmocks.NewStructWithPrivateField("foo")

	privateField, err := testutils.ReadPrivateField[testutilsmocks.StructWithPrivateField, string](src, "privateField")
	require.NoError(t, err)
	require.Equal(t, "foo", privateField)

	err = testutils.AssignPrivateField[testutilsmocks.StructWithPrivateField, string](src, "privateField", "bar")
	require.NoError(t, err)
	require.Equal(t, "bar", src.GetPrivateField())

	t.Run("FieldNotFound", func(t *testing.T) {
		_, err = testutils.ReadPrivateField[testutilsmocks.StructWithPrivateField, string](src, "missingField")
		require.ErrorIs(t, err, testutils.ErrFieldNotFound)

		err = testutils.AssignPrivateField[testutilsmocks.StructWithPrivateField, string](src, "missingField", "bar")
		require.ErrorIs(t, err, testutils.ErrFieldNotFound)
	})

	t.Run("NonStructPtr", func(t *testing.T) {
		_, err = testutils.ReadPrivateField[map[string]string, int](lo.ToPtr(map[string]string{}), "privateField")
		require.ErrorIs(t, err, testutils.ErrNonStructPtr)

		err = testutils.AssignPrivateField[map[string]string, int](lo.ToPtr(map[string]string{}), "privateField", 42)
		require.ErrorIs(t, err, testutils.ErrNonStructPtr)
	})
}
