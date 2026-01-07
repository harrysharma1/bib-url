package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootPersistentFlagExist(t *testing.T) {
	flags := rootCmd.PersistentFlags()
	copy := flags.Lookup("copy")
	braces := flags.Lookup("braces")
	save := flags.Lookup("save")
	assert.NotNil(t, copy, "copy flag should exist")
	assert.NotNil(t, braces, "braces flag should exist")
	assert.NotNil(t, save, "save flag should exist")
}
