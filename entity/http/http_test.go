package http

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCall(t *testing.T) {
	memory := NewHttp()
	ctx := context.Background()
	res, err := memory.Call(ctx, "https://google.com")
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
