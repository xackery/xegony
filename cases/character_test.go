package cases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharacterInitialize(t *testing.T) {
	var err error
	charRepo := &CharacterRepository{}
	err = charRepo.Initialize(nil)
	assert.EqualError(t, err, "Invalid storage type")
}

func TestCharacterGet(t *testing.T) {
	charRepo := &CharacterRepository{}
	_, err := charRepo.Get(0)
	assert.EqualError(t, err, "Invalid Character ID")
}
