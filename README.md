# Go API client for United States Patent and Trademark Office (USPTO)

Interact with the API of the USPTO

## Status

Alpha Version

## Standards

* (JAN 1976 – DEC 2001) Patent Grant Full Text Data/APS – Green Book
* (JAN 2001 – DEC 2001) Patent Grant Full Text Data/SGML Version 2.4
* (JAN 2002 – DEC 2004) Patent Grant Full Text Data/XML Version 2.5 
* (JAN 2005 – DEC 2005) Patent Grant Full Text Data/XML Version 4.0 ICE
* (JAN 2006 – DEC 2006) Patent Grant Full Text Data/XML Version 4.1 ICE
* (JAN 2007 – DEC 2012) Patent Grant Full Text Data/XML Version 4.2 ICE
* (JAN 2013 – DEC 2013) Patent Grant Full Text Data/XML Version 4.3 ICE
* (JAN 2013 – DEC 2014) Patent Grant Full Text Data/XML Version 4.4 ICE
* (JAN 2015 - Present) Patent Grant Full Text Data/XML Version 4.5 ICE

[See dtd standards](dtds)

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