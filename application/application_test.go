package application

import (
	"bytes"
	mocks "cmlabs-backend-crawler-freelance-test/framework/mocks/usecase"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	errorMessage = "error"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestNewApplication(t *testing.T) {
	NewApplication(nil)
}

func TestCrawl(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUsecase := mocks.NewMockIUsecase(mockController)
	app := &Application{
		Usecase: mockUsecase,
	}
	assert.NotNil(t, app)

	actionHttpMethod := "POST"
	actionHttpUrl := "/api/image/convert"

	router := initRouter()
	router.POST(actionHttpUrl, app.Crawl)

	t.Run("200", func(t *testing.T) {
		mockUsecase.EXPECT().Crawl(gomock.Any(), gomock.Any()).Return(nil)
		body := `{"url": ["https://google.com"]}`
		res := httptest.NewRequest(actionHttpMethod, actionHttpUrl, bytes.NewBuffer([]byte(body)))
		res.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, res)

		var response Response
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.Nil(t, err)
		assert.Equal(t, 200, response.Code)
	})

	t.Run("400", func(t *testing.T) {
		res := httptest.NewRequest(actionHttpMethod, actionHttpUrl, nil)
		res.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, res)

		var response Response
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.Nil(t, err)
		assert.Equal(t, 400, response.Code)
	})
}
