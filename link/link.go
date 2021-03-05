package link

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

// Link represents structure of the Link header as defined in RFC 5988.
type Link struct {
	URI    string
	Rel    string
	Params map[string]string
}

const (
	// FirstRel represents "first" value
	FirstRel = "first"
	// LastRel represents "last" value
	LastRel = "last"
	// NextRel represents "next" value
	NextRel = "next"
	// PrevRel represents "prev" value
	PrevRel = "prev"
)

// Serialize generates HTTP link header as defined in RFC 5988.
// It always generates relation-type (e.g., rel="next") params next to the link-value.
// Meanwhile, other link-param values are sorted alphabetically.
// For example, Link: <https://api.example.com/scenarios?page=10>; rel="next"; title="Next scenarios"; total="1000"
func Serialize(links []Link) (string, error) {
	var b strings.Builder

	for i, link := range links {
		if strings.TrimSpace(link.URI) == "" {
			return "", errors.New("URI cannot be blank")
		}

		if i == 0 {
			b.WriteString("Link: ")
		}
		if i > 0 { // append comma using index as opposing to joins
			b.WriteString(", ")
		}

		// Standard link header
		fmt.Fprintf(&b, `<%s>; rel="%s"`, link.URI, link.Rel)

		// Following with other params, alphabetically sorted
		keys := make([]string, 0, len(link.Params))
		for k := range link.Params {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			b.WriteString("; ")
			fmt.Fprintf(&b, `%s="%s"`, k, link.Params[k])
		}
	}

	return b.String(), nil
}
