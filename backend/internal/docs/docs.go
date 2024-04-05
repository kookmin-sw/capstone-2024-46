package docs

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

func byteHandler(b []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Write(b)
	}
}

func HandleSpec(swagger *openapi3.T) http.HandlerFunc {
	b, err := swagger.MarshalJSON()
	if err != nil {
		panic(err)
	}
	return byteHandler(b)
}

func UseSwagger(e *echo.Echo, swagger *openapi3.T) {
	e.GET("/docs/swagger.json", echo.WrapHandler(HandleSpec(swagger)))
	e.Static("/docs", "docs/swagger-ui")
}
