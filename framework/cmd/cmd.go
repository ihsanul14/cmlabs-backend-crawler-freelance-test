package cmd

import (
	"cmlabs-backend-crawler-freelance-test/application"
	"cmlabs-backend-crawler-freelance-test/entity/http"
	"cmlabs-backend-crawler-freelance-test/entity/memory"
	"cmlabs-backend-crawler-freelance-test/framework/router"
	"cmlabs-backend-crawler-freelance-test/usecase"
)

func Run() {
	entityHttp := http.NewHttp()
	entityMemory := memory.NewMemory()
	uc := usecase.NewUsecase(entityHttp, entityMemory)
	app := application.NewApplication(uc)
	router := router.NewRouter(app)
	router.Run()
}
