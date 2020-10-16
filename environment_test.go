package environment

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvironment(t *testing.T) {
	// default environment is always production
	_ = os.Setenv(environmentVariableName, "")
	bootstrap()
	assert.Equal(t, Production, GetEnvironment())

	// override to development
	OverrideEnvironment(Development)
	assert.Equal(t, Development, GetEnvironment())
}
