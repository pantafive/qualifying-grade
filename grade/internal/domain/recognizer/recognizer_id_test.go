package recognizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecognizerIdEqualsTrue(t *testing.T) {
	id := NewRecognizerId(3)
	otherId := NewRecognizerId(3)
	assert.True(t, id.Equals(otherId))
}

func TestRecognizerIdEqualsFalse(t *testing.T) {
	id := NewRecognizerId(3)
	otherId := NewRecognizerId(4)
	assert.False(t, id.Equals(otherId))
}
