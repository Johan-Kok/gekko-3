// +build go1.13

package multierr

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type errGreatSadness struct{ id int }

func (errGreatSadness) Error() string {
	return "great sadness"
}

type errUnprecedentedFailure struct{ id int }

func (errUnprecedentedFailure) Error() string {
	return "unprecedented failure"
}

func (e errUnprecedentedFailure) Unwrap() error {
	return errRootCause{e.id}
}

type errRootCause struct{ i int }

func (errRootCause) Error() string {
	return "root cause"
}

func TestErrorsWrapping(t *testing.T) {
	err := Append(
		errGreatSadness{42},
		errUnprecedentedFailure{43},
	)

	t.Run("left", func(t *testing.T) {
		t.Run("As", func(t *testing.T) {
			var got errGreatSadness
			require.True(t, errors.As(err, &got))
			assert.Equal(t, 42, got.id)
		})

		t.Run("Is", func(t *testing.T) {
			assert.False(t, errors.Is(err, errGreatSadness{41}))
			assert.True(t, errors.Is(err, errGreatSadness{42}))
		})
	})

	t.Run("right", func(t *testing.T) {
		t.Run("As", func(t *testing.T) {
			var got errUnprecedentedFailure
			require.True(t, errors.As(err, &got))
			assert.Equal(t, 43, got.id)
		})

		t.Run("Is", func(t *testing.T) {
			assert.False(t, errors.Is(err, errUnprecedentedFailure{42}))
			assert.True(t, errors.Is(err, errUnprecedentedFailure{43}))
		})
	})

	t.Run("top-level", func(t *testing.T) {
		t.Run("As", func(t *testing.T) {
			var got interface{ Errors() []error }
			require.True(t, errors.As(err, &got))
			assert.Len(t, got.Errors(), 2)
		})

		t.Run("Is", func(t *testing.T) {
			assert.True(t, errors.Is(err, err))
		})
	})

	t.Run("root cause", func(t *testing.T) {
		t.Run("As", func(t *testing.T) {
			var got errRootCause
			require.True(t, errors.As(err, &got))
			assert.Equal(t, 43, got.i)
		})

		t.Run("Is", func(t *testing.T) {
			assert.False(t, errors.Is(err, errRootCause{42}))
			assert.True(t, errors.Is(err, errRootCause{43}))
		})
	})

	t.Run("mismatch", func(t *testing.T) {
		t.Run("As", func(t *testing.T) {
			var got *os.PathError
			assert.False(t, errors.As(err, &got))
		})

		t.Run("Is", func(t *testing.T) {
			assert.False(t, errors.Is(err, errors.New("great sadness")))
		})
	})
}

func TestErrorsWrappingSameType(t *testing.T) {
	err := Combine(
		errGreatSadness{1},
		errGreatSadness{2},
		errGreatSadness{3},
	)

	t.Run("As returns first", func(t *testing.T) {
		var got errGreatSadness
		require.True(t, errors.As(err, &got))
		assert.Equal(t, 1, got.id)
	})

	t.Run("Is matches all", func(t *testing.T) {
		assert.True(t, errors.Is(err, errGreatSadness{1}))
		assert.True(t, errors.Is(err, errGreatSadness{2}))
		assert.True(t, errors.Is(err, errGreatSadness{3}))
	})
}
