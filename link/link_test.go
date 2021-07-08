package link

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerating(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		links []Link
		exp   string
		err   error
	}{
		{
			"ok with single link",
			[]Link{
				{
					URI: "https://api.example.com/scenarios?page=3",
					Rel: NextRel,
				},
			},
			`<https://api.example.com/scenarios?page=3>; rel="next"`,
			nil,
		},
		{
			"ok with multiple links",
			[]Link{
				{
					URI: "https://api.example.com/scenarios?page=3",
					Rel: NextRel,
				},
				{
					URI: "https://api.example.com/scenarios?page=1",
					Rel: PrevRel,
				},
				{
					URI: "https://api.example.com/scenarios?page=50",
					Rel: LastRel,
				},
			},
			`<https://api.example.com/scenarios?page=3>; rel="next", <https://api.example.com/scenarios?page=1>; rel="prev", <https://api.example.com/scenarios?page=50>; rel="last"`,
			nil,
		},
		{
			"ok with multiple links and params",
			[]Link{
				{
					URI: "https://api.example.com/scenarios?page=10",
					Rel: NextRel,
					Params: map[string]string{
						"title": "Next scenarios are waiting",
						"total": "1000",
					},
				},
				{
					URI: "https://api.example.com/scenarios?page=1",
					Rel: FirstRel,
					Params: map[string]string{
						"hreflang": "th-TH",
						"rev":      "canonical",
						"title":    "The begin of everything",
						"total":    "1000",
					},
				},
			},
			`<https://api.example.com/scenarios?page=10>; rel="next"; title="Next scenarios are waiting"; total="1000", ` +
				`<https://api.example.com/scenarios?page=1>; rel="first"; hreflang="th-TH"; rev="canonical"; title="The begin of everything"; total="1000"`,
			nil,
		},
		{
			"err when link URI is blank",
			[]Link{
				{
					URI: " ",
					Rel: NextRel,
				},
			},
			"",
			errors.New("URI cannot be blank"),
		},
		{
			"err when link URI is missing",
			[]Link{
				{
					Rel: NextRel,
				},
			},
			"",
			errors.New("URI cannot be blank"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			result, err := Serialize(test.links)

			assert.Equal(t, test.err, err)
			assert.Equal(t, test.exp, result)
		})
	}
}
