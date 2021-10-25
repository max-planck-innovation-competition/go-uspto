# Go API client for United States Patent and Trademark Office (USPTO)

Interact with the API of the USPTO

## Status

Alpha Version

## Installation

Add the package to your project via the following command:

```shell
go get github.com/max-planck-innovation-competition/go-uspto
```

## Usage

```go
import "github.com/max-planck-innovation-competition/go-uspto/pkg/uspto" 

...

// get all download links for bulk xml patent archives 
loc, _ := time.LoadLocation("Europe/Berlin")
start := time.Date(2021, 9, 1, 0, 0, 0, 0, loc)
end := time.Date(2021, 10, 1, 0, 0, 0, 0, loc)
res, err := GetPatentXmlBulkFileList(start, end)
```