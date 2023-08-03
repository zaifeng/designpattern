package strategy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonSave(t *testing.T) {
	storage := new(FileStorage)
	storage.SetStrategy(new(JsonSave))
	assert.NoError(t, storage.Save("test.json", []byte("Hello Json")))
}

func TestFileSave(t *testing.T) {
	storage := new(FileStorage)
	storage.SetStrategy(new(FileSave))
	assert.NoError(t, storage.Save("test.txt", []byte("Hello Text")))
}
