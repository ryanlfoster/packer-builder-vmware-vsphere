package vsphere

import (
	vmwcommon "github.com/mitchellh/packer/builder/vmware/common"
	"github.com/mitchellh/packer/common"
	"github.com/mitchellh/packer/helper/config"
)

type Config struct {
	common.PackerConfig      `mapstructure:",squash"`
	vmwcommon.DriverConfig   `mapstructure:",squash"`
	vmwcommon.OutputConfig   `mapstructure:",squash"`
	vmwcommon.RunConfig      `mapstructure:",squash"`
	vmwcommon.ShutdownConfig `mapstructure:",squash"`
	vmwcommon.SSHConfig      `mapstructure:",squash"`
	vmwcommon.ToolsConfig    `mapstructure:",squash"`
	vmwcommon.VMXConfig      `mapstructure:",squash"`

	VSphereURL string `mapstructure:"vsphere_url"`
}

func NewConfig(raws ...interface{}) (*Config, []string, error) {
	c := &Config{}
	err := config.Decode(c, &config.DecodeOpts{
		Interpolate: true,
	}, raws...)

	if err != nil {
		return nil, nil, err
	}

	warnings := []string{}

	return c, warnings, nil
}
