# Storage Service GRPC Client

Example (using local file)
```go
package main

import (
	"context"
	"os"

	storageservice "monorepo/sdk/storage-service"

	"pkg.agungdwiprasetyo.com/candi/candihelper"
)

func main() {

	file, err := os.Open("[path-to-your-file]")
	if err != nil {
		panic(err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}

	storageService, _ := storageservice.NewStorageServiceGRPC("[storage-service host]", "Basic xxxxx", 50*candihelper.MByte)
	storageService.Upload(context.Background(), file, storageservice.Header{
		Size:   fileInfo.Size(),
		Folder: "", Filename: "file.ext",
	})
}
```
