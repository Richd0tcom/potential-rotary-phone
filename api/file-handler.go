package api

import (
	"encoding/json"
	"errors"
	"fmt"
	// "io"
	"os"

	"github.com/Richd0tcom/potential-rotary-phone/utils"
	"github.com/gofiber/fiber/v2"
)

type Video struct {
	Name string `json:"name"`
	BytesWritten int `json:"bytes_written"`
}

func NewStore() Video {
	return Video{
		
	}
}

var Vs Video = NewStore()

var dbName string = "./db.json"

var str = make(map[string]int)




func HandlePreUpload(ctx *fiber.Ctx) error {

	id:= utils.RandomString()

	fmt.Println(id)

	return ctx.JSON(map[string]interface{}{
		"file-id":id,
	})
	
}

func HandleUpload(ctx *fiber.Ctx) error {

	stream := ctx.Body()

	// stream, err:= io.ReadAll(ctx.Request().BodyStream())

	// if err != nil {
	// 	return err
	// }

	fileName := ctx.Get("file-id")
	extension:= ctx.Get("extension")

	if fileName == "" {
		return fiber.ErrBadRequest
	}

	if extension == "" {
		extension = "mp4"
	}

	_, err := os.Stat("./" + fileName+ "."+extension)

	dets := make(map[string]Video)

	if err != nil {
		// when file is a new file and has never been created before
		if errors.Is(err, os.ErrNotExist) {
			tmpfile, err := os.Create("./" + fileName+ "."+extension)


			if err != nil {
				return err
			}

			defer tmpfile.Close()

			n, _:= tmpfile.Write(stream)

			_, err = os.Stat(dbName)

			if errors.Is(err, os.ErrNotExist) {
				db, err := os.Create(dbName)



				if err != nil {
					return err
				}

				defer db.Close()

				


				//write the database stuff here

				dets[fileName] = Video{
						Name: fileName,
						BytesWritten: n,}

						b, err := json.Marshal(dets)

						str[fileName] = n
						if err == nil {
							_,_ = db.Write(b)
						}
			}

			dbJson, err:= os.ReadFile(dbName)

			if err != nil {
				return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
			}

			json.Unmarshal(dbJson, &dets)

			dets[fileName] = Video{
				Name: fileName,
				BytesWritten: n,}
			
				b, err := json.Marshal(dets)

						str[fileName] = n
						if err == nil {
							os.WriteFile(dbName, b, 0777)
						}

			
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
	n, _ :=file.Write(stream)

	fmt.Println(n)

	//set properties
	


	dbJson, err:= os.ReadFile(dbName)

	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}
	 
	json.Unmarshal(dbJson, &dets)

	str[fileName] = str[fileName] + n
	
	fmt.Println(str[fileName])

	vid:= dets[fileName]

	// fmt.Println("Bytes written: ", vid)

	vid.BytesWritten = vid.BytesWritten + n

	fmt.Println("Dbug: ", vid.BytesWritten)
	dets[fileName] = vid

	b, err := json.Marshal(dets)

	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}

	os.WriteFile(dbName, b, 0777)

	return ctx.SendString("done writing existing")
}

 
func ServeVideoData(ctx *fiber.Ctx) error {

	fileName := ctx.Get("file-id")
	extension:= ctx.Get("extension")

	dets := make(map[string]Video)


	if fileName == "" {
		return fiber.ErrBadRequest
	}

	if extension == "" {
		extension = "mp4"
	}

	dbJson, err:= os.ReadFile(dbName)

	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}
	
	json.Unmarshal(dbJson, &dets)

	vid:= dets[fileName]

	return ctx.JSON(vid)
	
}

func RedirectToDocs(ctx *fiber.Ctx) error {

	return ctx.Redirect("https://github.com/Richd0tcom/potential-rotary-phone/blob/main/README.md")
}