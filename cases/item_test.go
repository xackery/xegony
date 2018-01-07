package cases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemInitialize(t *testing.T) {
	var err error
	itemRepo := &ItemRepository{}
	err = itemRepo.Initialize(nil)
	assert.EqualError(t, err, "Invalid storage type")
}

func TestItemGet(t *testing.T) {
	itemRepo := &ItemRepository{}
	_, err := itemRepo.Get(0)
	assert.EqualError(t, err, "Invalid Item ID")
}

func TestItemCreate(t *testing.T) {
	itemRepo := &ItemRepository{}
	err := itemRepo.Create(nil)
	assert.EqualError(t, err, "Empty item")
}
