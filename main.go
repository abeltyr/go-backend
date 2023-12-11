// main.go
package main

import (
	"adtec_backend/src/graphql/generated"
	graph "adtec_backend/src/graphql/resolver"
	"adtec_backend/src/middleware"
	"context"
	"flag"
	"log"
	"os"

	"github.com/arsmn/fastgql/graphql"
	"github.com/arsmn/fastgql/graphql/handler"
	"github.com/arsmn/fastgql/graphql/playground"
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

var app *fiber.App

// init the Fiber Server
func init() {
	log.Printf("Fiber cold start")

	godotenv.Load(".env")
	app = fiber.New(fiber.Config{
		BodyLimit:   20 * 1024 * 1024, // this is the default limit of MB
		ProxyHeader: "X-Real-IP",
	})
	app.Use(logger.New())
	app.Use(cors.New())

	api := app.Group("/api", middleware.Auth)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(
			"Everything is up and running",
		)
	})

	app.Get("/playground", func(c *fiber.Ctx) error {
		playground := playground.Handler("GraphQL playground", "/query")

		playground(c.Context())
		return nil
	})

	api.All("/graphql", func(c *fiber.Ctx) error {
		config := generated.Config{
			Resolvers: &graph.Resolver{},
		}

		defer func() {
		}()

		// user := c.Locals("user")
		// accessToken := c.Locals("accessToken")
		srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))
		srv.AroundFields(
			func(ctx context.Context, next graphql.Resolver) (res interface{}, err error) {
				// ctx = context.WithValue(ctx, model.ConfigKey("user"), user)
				// ctx = context.WithValue(ctx, model.ConfigKey("accessToken"), accessToken)
				// ctx = context.WithValue(ctx, model.ConfigKey("client"), client)
				// ctx = context.WithValue(ctx, model.ConfigKey("clientCtx"), clientCtx)
				// ctx = context.WithValue(ctx, model.ConfigKey("usageClient"), usageClient)
				// ctx = context.WithValue(ctx, model.ConfigKey("usageClientCtx"), usageClientCtx)
				return next(ctx)
			},
		)
		srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
			rc := graphql.GetOperationContext(ctx)
			rc.DisableIntrospection = true
			return next(ctx)
		})

		gqlHandler := srv.Handler()
		gqlHandler(c.Context())
		return nil
	})
}

var (
	addr = flag.String("addr :", os.Getenv("PORT"), "")
)

func main() {

	sentryUrl := os.Getenv("SENTRY")
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              sentryUrl,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	flag.Parse()

	if *addr == "" {
		*addr = ":9000"
	}
	err = app.Listen(*addr)

	if err != nil {
		log.Fatalln(err.Error())
	} else {
		log.Printf("Fiber cold start")
	}
}
