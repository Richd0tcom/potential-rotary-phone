# potential-rotary-phone


A github repo for the HNG chrom extension task (BE)

## API

There are 3 main endpoints available fo users

- POST `/api/pre-upload` 
  This endpoint returns a unique identifier for the file about to be uploaded.

  response:
  ```json
    {
        "file-id": "RSVqPyVGABDF1bMx5H7MgQQM"
    }
  ```

- POST `/api/upload`
  Using the unique id gotten in the previous step, the file to be upload is streamed in byte-sized chunks (ArrayBuffer) to this endpoint.
  The id is passed in the requset header along with the extension of the file.

  E.g

  ```javascript
  let input = document.getElementById('file-holder');
        let button = document.getElementById('file-button');


        button.addEventListener("click", (event) => {

            if (input.files.length == 0) {
                return
            }
            // event.preventDefault()
            // return
            let uploadFile = input.files[0];
            let reader = new FileReader()

            reader.addEventListener('loadend', async (e) => {
                
                //The array buffer containing the file in question
                let readData = reader.result
                
                const CHUNK_SIZE = 10000;
                
                let times = Math.ceil(uploadFile.size / CHUNK_SIZE)

                for (let index = 0; index < times; index++) {
                    let beginning = index * CHUNK_SIZE
                    let end  = (index +1) * CHUNK_SIZE
                    let chunk = readData.slice(beginning, end)

                    // console.log(chunk, "The chunk")

                   await fetch("http://localhost:1738/api/upload", {
                        method: "POST",
                        body:chunk,
                        headers:{
                           "content-type": "application/octet-stream",
                           "file-id": "RSVqPyVGABDF1bMx5H7MgQQM", //unique identifier gottnfrom previous step
                           "extension":  "mp4" //extension of file
                        }
                    })
                    
                }

            });

            reader.readAsArrayBuffer(uploadFile)

            
        })
  ```

response 
```txt
"done writing chunk"
```

- GET `/api/video-details`
  This endpoint returns the video details fo the file specified by the uniqu identfier. Pass the file id in the header as in the previous step.

  


