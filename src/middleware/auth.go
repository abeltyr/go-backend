package middleware

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jws"
)

var body []byte

func Auth(c *fiber.Ctx) error {
	// Set some security headers:
	c.Set("X-XSS-Protection", "1; mode=block")
	c.Set("X-Content-Type-Options", "nosniff")
	c.Set("X-Download-Options", "noopen")
	c.Set("Strict-Transport-Security", "max-age=5184000")
	c.Set("X-Frame-Options", "SAMEORIGIN")
	c.Set("X-DNS-Prefetch-Control", "off")

	// fetch the token and the needed details
	godotenv.Load(".env")
	token := c.Request().Header.Peek("Authorization")

	if token == nil {
		log.Println("no token")
		return errors.New("please provide authentication token")
	}

	// parse the jwk from the api request body
	set, err := jwk.Parse(body)
	if err != nil {
		log.Printf("failed to parse JWK: %s", err)
		return err

	}

	var rsaData = ""
	// verify the token against the public key
	payload, err := jws.Verify(token, jwa.RS256, rsaData)
	if err != nil {
		log.Printf("failed to verify message: %s", err)
	}

	log.Println(set, payload)

	c.Locals("accessToken", string(token))

	return c.Next()
}
