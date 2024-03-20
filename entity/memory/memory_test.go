package memory

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	memory := NewMemory()
	ctx := context.Background()
	req := SaveRequest{
		Body:     "",
		Domain:   "test.com",
		FileName: "test.com",
	}
	err := memory.Save(ctx, req)
	assert.Nil(t, err)
	err = os.RemoveAll("framework")
	assert.Nil(t, err)
}
