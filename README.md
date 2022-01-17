# Go API client for United States Patent and Trademark Office (USPTO)

Interact with the API of the USPTO

## Status

Alpha Version

## Standards

### Grants

At the moment there are parsers that are tested with examples of the following versions of the USPTO XML format:

- [ ]  (JAN 2022 - Present) Patent Grant Full Text Data/XML Version 4.6 ICE
- [ ] (JAN 1976 – DEC 2001) Patent Grant Full Text Data/APS – Green Book
- [ ] (JAN 2001 – DEC 2001) Patent Grant Full Text Data/SGML Version 2.4
- [x] (JAN 2002 – DEC 2004) Patent Grant Full Text Data/XML Version 2.5
- [ ] (JAN 2005 – DEC 2005) Patent Grant Full Text Data/XML Version 4.0 ICE
- [ ] (JAN 2006 – DEC 2006) Patent Grant Full Text Data/XML Version 4.1 ICE
- [ ] (JAN 2007 – DEC 2012) Patent Grant Full Text Data/XML Version 4.2 ICE
- [ ] (JAN 2013 – DEC 2013) Patent Grant Full Text Data/XML Version 4.3 ICE
- [ ] (JAN 2013 – DEC 2014) Patent Grant Full Text Data/XML Version 4.4 ICE
- [x] (JAN 2015 - Present) Patent Grant Full Text Data/XML Version 4.5 ICE

[See dtd standards](dtds)

### Applications

- [ ] (JAN 2022 - Present) Patent Application Data/XML Version 4.6 ICE
- [x] (JAN 2015 - DEC 2021) Patent Application Data/XML Version 4.4 ICE
- [ ] (JAN 2013 – DEC 2014) Patent Application Data/XML Version 4.3 ICE
- [ ] (JAN 2007 – DEC 2012) Patent Application Data/XML Version 4.2 ICE
- [ ] (JAN 2006 – DEC 2006) Patent Application Data/XML Version 4.1 ICE 
- [ ] (JAN 2005 – DEC 2005) Patent Application Data/XML Version 4.0 ICE
- [ ] (JAN 2002 – DEC 2004) Patent Application Data/XML Version 1.6  
- [x] (MAR 15, 2001 – DEC 2001) Patent Application Data/XML Version 1.5

[See dtd standards](dtds)

## Installation

Add the package to your project via the following command:

```shell
go get github.com/max-planck-innovation-competition/go-uspto
```

## Usage

Get a list of all the dates where the USPTO has published new patent grants:
```go
import "github.com/max-planck-innovation-competition/go-uspto/pkg/uspto"

...

// get all download links for bulk xml patent archives between the following dates
loc, _ := time.LoadLocation("Europe/Berlin")
start := time.Date(2021, 9, 1, 0, 0, 0, 0, loc)
end := time.Date(2021, 10, 1, 0, 0, 0, 0, loc)
res, err := uspto.GetPatentXmlBulkFileList(start, end)
```

Download a bulk zip file from the USPTO:
```go
exportFilePath, err := uspto.DownloadBulkFile("https://bulkdata.uspto.gov/data/patent/grant/redbook/fulltext/2021/ipg210907.zip", "./test-data")
```

Process the zip file
```go
err := uspto.ProcessBulkFile("./test-data/pg020101.zip", "./test-data/pg020101/xml")
```

Process the xml file
```go
patDoc, err := uspto.ProcessXMLFileSimple("./test-data/2-5-b1-patent.xml")
```

