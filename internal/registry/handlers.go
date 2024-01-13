package registry

import (
	customHTTP "github.com/iki-rumondor/init-golang-service/internal/adapter/http"
	"github.com/iki-rumondor/init-golang-service/internal/application"
	"github.com/iki-rumondor/init-golang-service/internal/repository"
)

func GetHandlers(repo *repository.Repositories) *customHTTP.Handlers {
	auth_service := application.NewAuthService(repo)
	prodi_service := application.NewProdiService(repo)
	instrumen_service := application.NewInstrumenService(repo)
	util_service := application.NewUtilService(repo)

	return &customHTTP.Handlers{
		AuthHandler:      customHTTP.NewAuthHandler(auth_service),
		ProdiHandler:     customHTTP.NewProdiHandler(prodi_service),
		InstrumenHandler: customHTTP.NewInstrumenHandler(instrumen_service),
		UtilHandler:      customHTTP.NewUtilHandler(util_service),
	}
}
