package cmd

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func assertNonInheritedFlags(
	t *testing.T,
	cmd *cobra.Command,
	expectedFlags map[string]string) {
	t.Helper()
	haveFlags := make(map[string]*pflag.Flag)
	cmd.NonInheritedFlags().VisitAll(func(f *pflag.Flag) {
		haveFlags[f.Name] = f
	})
	for flag := range expectedFlags {
		if assert.Contains(t, haveFlags, flag, "%s flag is missing", flag) {
			assert.Equal(t, flag, haveFlags[flag].Name)
		}
	}

}

func assertNonInheritedFlagsShorthand(
	t *testing.T,
	cmd *cobra.Command,
	expectedFlags map[string]string) {
	t.Helper()
	haveFlagsShorthand := make(map[string]*pflag.Flag)
	cmd.NonInheritedFlags().VisitAll(func(f *pflag.Flag) {
		haveFlagsShorthand[f.Name] = f
	})
	for flag, shorthand := range expectedFlags {
		if assert.Contains(t, haveFlagsShorthand, flag, "%s  (--%s) flag missing", shorthand, flag) {
			assert.Equal(t, shorthand, haveFlagsShorthand[flag].Shorthand)
		}
	}

}

func TestNonInheritedFlags(t *testing.T) {
	tests := []struct {
		cmd           *cobra.Command
		expectedFlags map[string]string
	}{
		{articleCmd, articleExpectedFlags},
		{bookCmd, bookExpectedFlags},
		{bookletCmd, bookletExpectedFlags},
		{inbookCmd, inbookExpectedFlags},
		{incollectionCmd, incollectionExpectedFlags},
		{inproceedingsCmd, inproceedingsExpectedFlags},
		{manualCmd, manualExpectedFlags},
		{mastersthesisCmd, mastersthesisExpectedFlags},
		{miscCmd, miscExpectedFlags},
		{phdthesisCmd, phdthesisExpectedFlags},
		{proceedingsCmd, proceedingsExpectedFlags},
		{techreportCmd, techreportExpectedFlags},
		{unpublishedCmd, unpublishedExpectedFlags}}
	for _, test := range tests {

		assertNonInheritedFlags(t, test.cmd, test.expectedFlags)
		assertNonInheritedFlagsShorthand(t, test.cmd, test.expectedFlags)
	}
}
