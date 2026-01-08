package cmd

import (
	"bibcli/models"
	"io"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"golang.design/x/clipboard"
)

// Mock for test with persistent flags e.g. --braces, --copy, --save
type PersistentFlagsMock struct {
	BracesArgs     []string
	BracesContains []string
	CopyArgs       []string
	CopyContains   []string
	SaveArgs       []string
	SaveContains   []string
	TempDirPath    string
}

// Helper: initialise struct with sub command and it's expected flags
func initCommands(t *testing.T) []struct {
	cmd           *cobra.Command
	expectedFlags map[string]string
} {
	t.Helper()
	retVal := []struct {
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
	return retVal
}

// Helper: initialise list of mocks for each subcommand to run dynamic tests
func initPersistentFlagMocks(t *testing.T) []PersistentFlagsMock {
	t.Helper()
	retVal := []PersistentFlagsMock{
		initArticleMock(t),
		initBookMock(t),
		initBookletMock(t),
		initInbookMock(t),
		initIncollectionMock(t),
		initInproceedingsMock(t),
		initManualMock(t),
	}
	return retVal
}

// Helper: resetting global state variables of persistent flags
func resetGlobalState(t *testing.T) {
	t.Helper()
	copy = false
	save = ""
	braces = false
	article = models.Article{}
	book = models.Book{}
	booklet = models.Booklet{}
	inbook = models.Inbook{}
	incollection = models.Incollection{}
	inproceedings = models.Inproceedings{}
	manual = models.Manual{}
	mastersthesis = models.Mastersthesis{}
	misc = models.Misc{}
	phdthesis = models.Phdthesis{}
	proceedings = models.Proceedings{}
	techreport = models.Techreport{}
	unpublished = models.Unpublished{}
}

// Helper: capturing stdout to test printing
// from: https://stackoverflow.com/questions/10473800/in-go-how-do-i-capture-stdout-of-a-function-into-a-string
func captureStdout(t *testing.T, fn func()) string {
	t.Helper()

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fn()
	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout

	return string(out)
}

// Helper: assertion for save function
func assertSaveFile(t *testing.T, path string, out string) {
	t.Helper()
	data, err := os.ReadFile(path)
	assert.Nil(t, err, "failed to open file at tmpdir: %s", path)
	assert.Equal(t, string(data), out, "saved file does not match command output string")

}

// Helper: assertion of flags specific for each subcommand
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

// Helper: assertion of flags as shorthand for each subcommand
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

// Check each subcommands flag/shorthand exists
func TestNonInheritedFlags(t *testing.T) {
	tests := initCommands(t)
	for _, test := range tests {
		assertNonInheritedFlags(t, test.cmd, test.expectedFlags)
		assertNonInheritedFlagsShorthand(t, test.cmd, test.expectedFlags)
	}
}

// Check each persistent flag exists within each subcommand
func TestInheritedPersistentFlags(t *testing.T) {
	tests := initCommands(t)
	for _, test := range tests {
		flags := test.cmd.InheritedFlags()
		assert.NotNil(t, flags.Lookup("copy"), "copy should persist to the child command %s", test.cmd.Name())
		assert.NotNil(t, flags.Lookup("braces"), "braces should persist to the child command %s", test.cmd.Name())
		assert.NotNil(t, flags.Lookup("save"), "save should persist to the child command %s", test.cmd.Name())
	}
}

// Testing functionality of `--braces` persistent flag
func TestDynamicPersistentFlagsBraces(t *testing.T) {
	resetGlobalState(t)
	mocks := initPersistentFlagMocks(t)
	for _, mock := range mocks {
		resetGlobalState(t)
		cmd := rootCmd
		cmd.SetArgs(mock.BracesArgs)
		out := captureStdout(t, func() {
			err := cmd.Execute()
			assert.Nil(t, err, "should not err check command flags in test")
		})
		for _, expected := range mock.BracesContains {
			assert.Contains(t, out, expected)
		}
	}
	resetGlobalState(t)
}

// Testing functionality of `--copy` persistent flag
func TestDynamicPersistentFlagsCopy(t *testing.T) {
	resetGlobalState(t)
	mocks := initPersistentFlagMocks(t)
	for _, mock := range mocks {
		resetGlobalState(t)
		cmd := rootCmd
		cmd.SetArgs(mock.CopyArgs)
		out := captureStdout(t, func() {
			err := cmd.Execute()
			assert.Nil(t, err, "should not err check command flags in test")
		})
		for _, expected := range mock.CopyContains {
			assert.Contains(t, out, expected)
		}
		hasClipboard := clipboard.Read(clipboard.FmtText)
		assert.Contains(t, out, string(hasClipboard), "clipboard should copy bibtex only not stdout")
	}
	resetGlobalState(t)
}

// Testing functionality of `--save` persistent flag
func TestDynamicPersistentFlagsSave(t *testing.T) {
	resetGlobalState(t)
	mocks := initPersistentFlagMocks(t)
	for _, mock := range mocks {
		resetGlobalState(t)
		cmd := rootCmd
		cmd.SetArgs(mock.SaveArgs)
		out := captureStdout(t, func() {
			err := cmd.Execute()
			assert.Nil(t, err, "should not err check command flags in test")
		})
		for _, expected := range mock.SaveContains {
			assert.Contains(t, out, expected)
		}
		assertSaveFile(t, mock.TempDirPath, out)
	}
	resetGlobalState(t)
}
