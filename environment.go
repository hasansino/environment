package environment

import (
	"os"
	"strings"
)

const environmentVariableName = "ENVIRONMENT"

const (
	environmentVariableValueProd     = "production"
	environmentVariableValueAvangard = "avangard"
	environmentVariableValueStaging  = "staging"
	environmentVariableValueDev      = "development"
)

const (
	Production  Environment = "production"  // self explanatory
	Avangard    Environment = "avangard"    // small portion of production servers
	Staging     Environment = "staging"     // `pre-production`, QA playground
	Development Environment = "development" // local instance
)

var currentEnvironment Environment

type Environment string

func (e Environment) String() string {
	return string(e)
}

func init() {
	bootstrap()
}

// bootstrap reads OS'es environment variable and initializes package
func bootstrap() {
	switch strings.ToLower(os.Getenv(environmentVariableName)) {
	case environmentVariableValueProd:
		currentEnvironment = Production
	case environmentVariableValueAvangard:
		currentEnvironment = Avangard
	case environmentVariableValueStaging:
		currentEnvironment = Staging
	case environmentVariableValueDev:
		currentEnvironment = Development
	default:
		currentEnvironment = Production
	}
}

// OverrideEnvironment overrides default detected environment
// Useful in unit tests, do not use in other cases
func OverrideEnvironment(e Environment) {
	currentEnvironment = e
}

// GetEnvironment returns current environment where service is executed
func GetEnvironment() Environment {
	return currentEnvironment
}

// IsProduction returns true if service operates in production environment
func IsProduction() bool {
	return GetEnvironment() == Production
}

// IsAvangard returns true if service operates in avangard environment
func IsAvangard() bool {
	return GetEnvironment() == Avangard
}

// IsStaging returns true if service operates in staging environment
func IsStaging() bool {
	return GetEnvironment() == Staging
}

// IsDevelopment returns true if service operates in development environment
func IsDevelopment() bool {
	return GetEnvironment() == Development
}
