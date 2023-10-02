package api

import (
	"errors"
	"fmt"
	"io"

	// "io"
	"os"

	"github.com/Richd0tcom/potential-rotary-phone/utils"
	"github.com/gofiber/fiber/v2"
)

type Store struct {
	Video map[string]interface{}
}

func NewStore() Store {
	return Store{}
}

var Vs Store = NewStore()

func HandlePreUpload(ctx *fiber.Ctx) error {

	id:= utils.RandomString()

	fmt.Println(id)

	return ctx.JSON(map[string]interface{}{
		"file-id":id,
	})
	
}

func HandleUpload(ctx *fiber.Ctx) error {

	// stream := ctx.Body()

	stream, err:= io.ReadAll(ctx.Request().BodyStream())

	if err != nil {
		return err
	}

	fileName := ctx.Get("file-id")
	extension:= ctx.Get("extension")

	if extension == "" {
		extension = "mp4"
	}

	_, err = os.Stat("./" + fileName+ "."+extension)

	if err != nil {
		// when file is a new file and has never been created before
		if errors.Is(err, os.ErrNotExist) {
			tmpfile, err := os.Create("./" + fileName+ "."+extension)

			if err != nil {
				return err
			}

			defer tmpfile.Close()

			tmpfile.Write(stream)
			
			return ctx.SendString("done writing")
	
		}

		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}

	//If the file exists
	file, err:= os.OpenFile(fileName+"."+extension, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)

	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}

	// Append to file, if it exists already
	file.Write(stream)

	return ctx.SendString("done writing existing")
}

 
func ServeVideoData(ctx *fiber.Ctx) error {

	return fiber.ErrBadGateway
}