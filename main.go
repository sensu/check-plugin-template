package main

import (
	"fmt"
	"log"

	"github.com/sensu-community/sensu-plugin-sdk/sensu"
	"github.com/sensu/sensu-go/types"
)

// Plugin ...
type Plugin struct {
	sensu.PluginConfig
	Example string
}

// Options ...
type Options struct {
	Example sensu.PluginConfigOption
}

var (
	plugin = Plugin{
		PluginConfig: sensu.PluginConfig{
			Name:     "{{ .Name }}",
			Short:    "{{ .Description }}",
			Timeout:  10,
			Keyspace: "sensu.io/plugins/{{ .Name }}/config",
		},
	}

	checkOptions = Options{
		Example: sensu.PluginConfigOption{
			Path:      "example",
			Env:       "CHECK_EXAMPLE",
			Argument:  "example",
			Shorthand: "e",
			Default:   "",
			Usage:     "An example configuration option",
			Value:     &plugin.Example,
		},
	}

	options = []*sensu.PluginConfigOption{
		&checkOptions.Example,
	}
)

func main() {
	check := sensu.NewGoCheck(&plugin.PluginConfig, options, checkArgs, executeCheck, false)
	check.Execute()
}

func checkArgs(event *types.Event) (int, error) {
	if len(plugin.Example) == 0 {
		return 0, fmt.Errorf("--example or CHECK_EXAMPLE environment variable is required")
	}
	return 0, nil
}

func executeCheck(event *types.Event) (int, error) {
	log.Println("executing check with --example", checkOptions.Example)
	return 0, nil
}
