package lazy

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		getter := New(func() int { return 1 })
		assert.Equal(t, 1, getter())
	})

	t.Run("OnlyOnce", func(t *testing.T) {
		var i int
		getter := New(func() int {
			i++
			return i
		})
		assert.Equal(t, 1, getter())
		assert.Equal(t, 1, getter())
		assert.Equal(t, 1, getter())
	})
}

func TestNewErrorable(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		getter := NewErrorable(func() (int, error) { return 1, nil })
		v, err := getter()
		require.NoError(t, err)
		assert.Equal(t, 1, v)
	})

	t.Run("OnlyOnce", func(t *testing.T) {
		var i int
		getter := NewErrorable(func() (int, error) {
			i++
			return i, nil
		})
		v, err := getter()
		assert.Equal(t, 1, v)
		assert.NoError(t, err)
		v, err = getter()
		assert.Equal(t, 1, v)
		assert.NoError(t, err)
		v, err = getter()
		assert.Equal(t, 1, v)
		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		var i int
		getter := NewErrorable(func() (int, error) {
			i++
			return i, assert.AnError
		})
		v, err := getter()
		assert.Equal(t, 1, v)
		assert.Error(t, err)
		v, err = getter()
		assert.Equal(t, 1, v)
		assert.Error(t, err)
		v, err = getter()
		assert.Equal(t, 1, v)
		assert.Error(t, err)
	})
}
