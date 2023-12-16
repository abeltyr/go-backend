// main.go
package main

import (
	graphqlGenerated "adtec/backend/src/graphql/generated"
	graphqlResolver "adtec/backend/src/graphql/resolver"
	"adtec/backend/src/middleware"
	"flag"
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/adaptor/v2"
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
		// playground := playground.Handler("GraphQL playground", "/query")

		// playground(c.Context())
		return nil
	})
	config := graphqlGenerated.Config{
		Resolvers: &graphqlResolver.Resolver{},
	}

	api.All("/graphql", func(c *fiber.Ctx) error {

		srv := handler.NewDefaultServer(graphqlGenerated.NewExecutableSchema(config))

		// srv.AroundFields()
		// srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		// 	rc := graphql.GetOperationContext(ctx)
		// 	rc.DisableIntrospection = true
		// 	return next(ctx)
		// })

		gqlHandler := adaptor.HTTPHandler(srv)
		gqlHandler(c)
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
