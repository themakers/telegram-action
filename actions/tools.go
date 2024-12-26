package actions

import (
	"encoding/json"
	"os"
	"slices"
	"strings"
)

func Input(name string) string {
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ToUpper(name)
	name = "INPUT_" + name
	return strings.TrimSpace(os.Getenv(name))
}

func Env() []string {
	var env = os.Environ()
	slices.SortFunc(env, func(a, b string) int {
		return strings.Compare(strings.ToLower(a), strings.ToLower(b))
	})
	return env
}

func DumpEnv() string {
	if data, err := json.MarshalIndent(Env(), "", "  "); err != nil {
		panic(err)
	} else {
		return string(data)
	}
}
