package main

import (
	"context"
	"os"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	oapimiddleware "github.com/oapi-codegen/echo-middleware"

	"private-llm-backend/internal/api"
	"private-llm-backend/internal/app/application/chat"
	"private-llm-backend/internal/app/controller"
	"private-llm-backend/internal/auth"
	"private-llm-backend/internal/config"
	"private-llm-backend/internal/database"
	"private-llm-backend/internal/docs"
	"private-llm-backend/internal/sentrysetup"
	"private-llm-backend/pkg/client/openai"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	// Load the configuration and secrets
	c, _, err := config.Load(context.Background(), "./configs/", env)
	if err != nil {
		log.Fatalf("failed to load config: %+v", err)
	}

	swagger, err := api.GetSwagger()
	if err != nil {
		panic(err)
	}

	// Allow all origins for swagger UI
	swagger.Servers = nil

	e := echo.New()
	e.HTTPErrorHandler = api.HTTPErrorHandler

	// Add middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Install Sentry
	err = sentrysetup.UseSentry(e, c)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	// Add swagger pages
	docs.UseSwagger(e, swagger)

	// Connect to the database
	dbConfig := database.NewMySQLConfig(*c.MySqlDsn)
	_, err = database.OpenConnection(dbConfig)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	// Initialize providers
	jwtProvider, err := auth.NewJWTProvider(*c.JWTSecret)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	// Initialize clients
	openaiClient, err := openai.NewClientWithApiKey(*c.OpenAIAPIKey)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	// Initialize the repository

	// Initialize the service
	chatService := chat.NewChatService(*c.AssistantID, openaiClient)

	// Initialize the adapter & facade

	// Initialize the authenticator
	jwtAuthenticator := auth.NewJWTAuthenticator(jwtProvider)
	apiKeyAuthenticator := auth.NewAPIKeyAuthenticator(*c.ClientKey)

	g := e.Group("")

	// Add CORS
	g.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:8080", "http://localhost:9090", "http://127.0.0.1:8080", "http://127.0.0.1:9090",
			"http://54.180.24.204:9090",
			"https://main.d2fs2yvrsy3li2.amplifyapp.com",
		},
		AllowCredentials: true,
		AllowHeaders: []string{
			echo.HeaderContentType, echo.HeaderAuthorization, echo.HeaderAccept, echo.HeaderOrigin, echo.HeaderCacheControl,
			"X-Api-Key",
		},
		ExposeHeaders: []string{
			echo.HeaderContentType, echo.HeaderContentLength, echo.HeaderContentEncoding, echo.HeaderContentDisposition,
		},
	}))

	// Add the API handlers
	g.Use(api.ContextMiddleware)
	g.Use(
		oapimiddleware.OapiRequestValidatorWithOptions(
			swagger,
			&oapimiddleware.Options{
				Options: openapi3filter.Options{
					AuthenticationFunc: api.NewAuthentication(map[string]api.Authenticator{
						"BearerAuth": jwtAuthenticator,
						"ApiKeyAuth": apiKeyAuthenticator,
					}),
				},
			},
		),
	)
	g.Use(sentrysetup.SetSentryExtraContextMiddleware)
	g.Use(api.NoCacheResponseMiddleware)
	handler := api.NewStrictHandler(
		controller.NewAPI(
			chatService,
		),
		nil,
	)
	api.RegisterHandlers(g, handler)

	// Add index page
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Private LLM API")
	})

	// Start the server
	e.Logger.Fatal(
		e.Start(c.ListenAddress(9090)),
	)
}
