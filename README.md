# gofiber_request

Este proyecto utiliza el framework [Fiber](https://gofiber.io/) para manejar solicitudes HTTP en Go.

## Instalación

Para instalar las dependencias del proyecto, ejecuta:

```bash
go get github.com/nelsonp17/gofiber_request
```

## Uso

A continuación se muestra un ejemplo básico de cómo utilizar este proyecto:

```go
import (
	"github.com/nelsonp17/gofiber_request"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func TestingRequest(c *fiber.Ctx) error {
    // campos y reglas
	Request := request.Request{
		Fields: map[string]string{
			"email":    "required|email",
			"age":      "required|integer",
			"phone":    "required|phone",
			"password": "required|password",
			"name":     "required|min:3|max:10",
			"price":    "required|price",
			"url":      "required|url",
			"float":    "required|float",
			"date":     "required|date",
			"datetime": "required|datetime",
			"time":     "required|time",
			"optional": "optional",
		},
	}
    // inicia la extracción de datos
	err := Request.Start(c) 
	if err != nil {
		fmt.Println("Error", err)
		return c.Status(400).JSON(err.Error())
	}

    // Válida los campos
	validated := Request.Validated()
	if validated == false {
		fmt.Println("Error Request", Request.Errors)
		return c.Status(400).JSON(Request.Errors)
	}

    // retorna los datos capturados
	return c.JSON(Request.Form)
}
```

## Contribuciones

Las contribuciones son bienvenidas. Por favor, abre un issue o un pull request para discutir cualquier cambio que te gustaría hacer.

## Licencia

Este proyecto está bajo la Licencia MIT. Consulta el archivo `LICENSE` para más detalles.