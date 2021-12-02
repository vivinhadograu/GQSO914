package main
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gofiber/fiber/v2"
)
func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Bem-vindo(a) a calculadora API.")
	})

	registerHandlers(app)

	app.Listen(":8000")
}
func registerHandlers(app *fiber.App) {
	http.HandleFunc("/soma/", somaHandler)
	http.HandleFunc("/sub/", subHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil))
}