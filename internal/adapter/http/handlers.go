package customHTTP

import "github.com/iki-rumondor/init-golang-service/internal/adapter/http/response"

var (
	badRequestError = &response.Error{
		Code:    400,
		Message: "Input Yang Anda Masukkan Tidak Valid",
	}
)

type Handlers struct {
	AuthHandler        *AuthHandler
	ProdiHandler       *ProdiHandler
	InstrumenHandler   *InstrumenHandler
	UtilHandler        *UtilHandler
	AssTypeHandler     *AssTypeHandler
	AssQuestionHandler *AssQuestionHandler
	ResponseHandler    *ResponseHandler
}
