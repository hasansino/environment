package environment

import (
	"os"
	"testing"
)

func TestEnvironment(t *testing.T) {
	// default environment is always production
	_ = os.Setenv(environmentVariableName, "")
	bootstrap()
	if Production != GetEnvironment() {
		t.Error("GetEnvironment() returned invalid value")
	}

	// override to development
	OverrideEnvironment(Development)
	if Development != GetEnvironment() {
		t.Error("GetEnvironment() returned invalid value")
	}
}
