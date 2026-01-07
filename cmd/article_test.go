package cmd

import (
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestArticleCmdNonInheritedFlags(t *testing.T) {
	haveFlags := make(map[string]*pflag.Flag)
	articleCmd.NonInheritedFlags().VisitAll(func(f *pflag.Flag) {
		haveFlags[f.Name] = f
	})
	expectedFlags := []string{
		"author", "doi", "journal", "key", "month",
		"note", "number", "pages", "title", "volume", "year",
	}

	for _, flag := range expectedFlags {
		if assert.Contains(t, haveFlags, flag, "%s flag is missing", flag) {
			assert.Equal(t, flag, haveFlags[flag].Name)
		}
	}
}

func TestArticleCmdNonInheritedFlagsShorthand(t *testing.T) {
	haveFlagsShorthand := make(map[string]*pflag.Flag)
	articleCmd.NonInheritedFlags().VisitAll(func(f *pflag.Flag) {
		haveFlagsShorthand[f.Name] = f
	})
	expected := map[string]string{
		"author":  "a",
		"doi":     "d",
		"journal": "j",
		"key":     "k",
		"month":   "m",
		"note":    "n",
		"number":  "i",
		"pages":   "p",
		"title":   "t",
		"volume":  "v",
		"year":    "y",
	}

	for name, shorthand := range expected {
		if assert.Contains(t, haveFlagsShorthand, name, "%s flag missing", name) {
			assert.Equal(t, shorthand, haveFlagsShorthand[name].Shorthand)
		}
	}
}
