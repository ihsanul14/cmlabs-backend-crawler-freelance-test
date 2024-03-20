package usecase

import (
	mocks "cmlabs-backend-crawler-freelance-test/framework/mocks/entity"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCrawl(t *testing.T) {
	ctx := context.Background()
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	data := []string{"http://google.com"}
	mockHttp := mocks.NewMockIHttp(mockController)
	mockMemory := mocks.NewMockIMemory(mockController)
	u := NewUsecase(mockHttp, mockMemory)

	t.Run("Success", func(t *testing.T) {
		err := u.Crawl(ctx, data)
		assert.Nil(t, err)
	})
}
