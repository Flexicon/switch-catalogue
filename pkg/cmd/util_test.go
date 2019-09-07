package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMessageWithUnderline(t *testing.T) {
	actual := messageWithUnderline("Test")
	expected := "Test\n====\n"

	assert.Equal(t, expected, actual)
}

func TestMessageWithUnderlineEmptyString(t *testing.T) {
	actual := messageWithUnderline("")
	expected := ""

	assert.Equal(t, expected, actual)
}

func TestDryRunBanner(t *testing.T) {
	actual := dryRunBanner()
	expected := "--- Dry Run ---"

	assert.Equal(t, expected, actual)
}
