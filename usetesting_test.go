package usetesting

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testCases := []struct {
		dir     string
		options map[string]string
	}{
		{dir: "oschdir/basic"},
		{dir: "oschdir/dot"},
		{dir: "oschdir/nottestfiles"},
		{dir: "oschdir/disable", options: map[string]string{"oschdir": "false"}},

		{dir: "contextbackground/basic"},
		{dir: "contextbackground/dot"},
		{dir: "contextbackground/nottestfiles"},
		{dir: "contextbackground/disable", options: map[string]string{"contextbackground": "false"}},

		{dir: "contexttodo/basic"},
		{dir: "contexttodo/dot"},
		{dir: "contexttodo/nottestfiles"},
		{dir: "contexttodo/disable", options: map[string]string{"contexttodo": "false"}},
	}

	for _, test := range testCases {
		t.Run(test.dir, func(t *testing.T) {
			newAnalyzer := NewAnalyzer()

			for k, v := range test.options {
				err := newAnalyzer.Flags.Set(k, v)
				if err != nil {
					t.Fatal(err)
				}
			}

			analysistest.Run(t, analysistest.TestData(), newAnalyzer, test.dir)
		})
	}
}
