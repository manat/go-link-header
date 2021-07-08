# link-header

This project implements serialization of HTTP Link headers as defined in [RFC 5988](https://tools.ietf.org/html/rfc5988). However, it does not claim to fully implement the spec, but more focusing on supporting the major usecases.

### Usage

#### Serialize Example

```go
package main

import (
	"fmt"

	"github.com/manat/go-link-header/link"
)

func main() {
	links := []link.Link{
		{
			URI: "https://api.example.com/scenarios?page=3",
			Rel: link.NextRel,
		},
		{
			URI: "https://api.example.com/scenarios?page=1",
			Rel: link.PrevRel,
		},
		{
			URI: "https://api.example.com/scenarios?page=50",
			Rel: link.LastRel,
			Params: map[string]string{
				"title": "This is the last page",
				"total": "5000",
			},
		},
	}

	fmt.Println(link.Serialize(links))
	//<https://api.example.com/scenarios?page=3>; rel="next", <https://api.example.com/scenarios?page=1>; rel="prev", <https://api.example.com/scenarios?page=50>; rel="last"; title="This is the last page"; total="5000"
}
```

### Reference
* https://github.com/tent/http-link-go/blob/master/link.go
* https://github.com/peterhellberg/link
