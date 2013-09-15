package text

import (
	"io/ioutil"
	"reflect"
	"testing"

	"launchpad.net/goyaml"
)

type conformanceSuite struct {
	Tests struct {
		Hashtags []struct {
			Description string
			Text        string
			Expected    []string
		}
		HashtagIndices []struct {
			Description string
			Text        string
			Expected    []struct {
				Hashtag string
				Indices []int
			}
		} `yaml:"hashtags_with_indices"`

		URLs []struct {
			Description string
			Text        string
			Expected    []string
		}
	}
}

var conformance conformanceSuite

func init() {
	data, err := ioutil.ReadFile("conformance/extract.yml")
	if err != nil {
		panic(err)
	}
	if err := goyaml.Unmarshal(data, &conformance); err != nil {
		panic(err)
	}
}

func TestExtractHashtags(t *testing.T) {
	for _, test := range conformance.Tests.Hashtags {
		res := ExtractHashtags(test.Text)
		if len(test.Expected) == 0 && len(res) == 0 {
			continue
		}
		if !reflect.DeepEqual(res, test.Expected) {
			t.Errorf("%s: got %v, want %v", test.Description, res, test.Expected)
		}
	}
}

func TestExtractHashtagIndices(t *testing.T) {
	for _, test := range conformance.Tests.HashtagIndices {
		res := ExtractHashtagIndices(test.Text)
		if len(test.Expected) == 0 && len(res) == 0 {
			continue
		}
		if len(test.Expected) != len(res) {
			t.Errorf("%s: want %s, got %s", test.Description, res, test.Expected)
		}
		for i, expected := range test.Expected {
			hashtag := test.Text[res[i][0]:res[i][1]]
			ex := unicodeSlice(test.Text, expected.Indices[0], expected.Indices[1])
			if hashtag != ex || hashtag != expected.Hashtag {
				t.Errorf("%s: [%d] want %s %s, got %s", test.Description, i, ex, expected.Hashtag, hashtag)
			}
		}
	}
}

var skipURLTests = map[int]bool{
	32: true, 33: true, 34: true, // CJK surrounded without protocol
	65: true, 66: true, // special t.co extraction
}

func TestExtractURLs(t *testing.T) {
	for i, test := range conformance.Tests.URLs {
		if skipURLTests[i] {
			continue
		}
		res := ExtractURLs(test.Text)
		if len(test.Expected) == 0 && len(res) == 0 {
			continue
		}
		if !reflect.DeepEqual(res, test.Expected) {
			t.Errorf("%s: [%d] got %#v, want %#v", test.Description, i, res, test.Expected)
		}
	}
}

func unicodeSlice(s string, j, k int) string {
	res := make([]rune, 0, k-j-1)
	var i int
	for _, r := range s {
		i++
		if i-1 <= j {
			continue
		}
		if i > k {
			break
		}
		res = append(res, r)
	}
	return string(res)
}
