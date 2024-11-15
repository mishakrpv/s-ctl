package cmd

import "github.com/s-ctl/ams/pkg/config/static"

// AMSCmdConfiguration wraps the static configuration and extra parameters.
type AMSCmdConfiguration struct {
	static.Configuration `export:"true"`
	// ConfigFile is the path to the configuration file.
	ConfigFile string `description:"Configuration file to use. If specified all other flags are ignored." export:"true"`
}

func NewAWSCmdConfiguration() *AMSCmdConfiguration {
	return &AMSCmdConfiguration{}
}
