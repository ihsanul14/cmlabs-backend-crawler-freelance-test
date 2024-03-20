package application

import (
	"cmlabs-backend-crawler-freelance-test/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IApplication interface {
	Crawl(*gin.Context)
}

type Application struct {
	Usecase usecase.IUsecase
}

type Request struct {
	Url []string
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewApplication(uc usecase.IUsecase) IApplication {
	return &Application{
		Usecase: uc,
	}
}

// @Summary Crawl
// @Description Handle Crawling Website
// @Accept json
// @Produce json
// @Param request body Request true "Sample request payload"
// @Success 200 {object} Response
// @Router /api/image/convert [post]
func (a *Application) Crawl(c *gin.Context) {
	var req *Request
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	err = a.Usecase.Crawl(c, req.Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Code:    500,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Code:    200,
		Message: "your request is being processed, see the log for the details",
	})
}
