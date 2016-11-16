package ini

import (
	"os"
	"regexp"
)

// Variables in a ${NAME} form inside configuration file are
// expected to be treated as ENV vars.
var envVar = regexp.MustCompile(`\${([A-Za-z0-9._\-]+)}`)

// replaceEnvVars replaces environment variables in the received value,
// i.e. every ${SOME_VAR} is replaced by the corresponding environment
// variable's value.
func replaceEnvVars(s string) string {
	return envVar.ReplaceAllStringFunc(s, func(k string) string {
		return os.Getenv(envVar.ReplaceAllString(k, "$1"))
	})
}
