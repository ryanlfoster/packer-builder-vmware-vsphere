package vsphere

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/mitchellh/multistep"
	vmwcommon "github.com/mitchellh/packer/builder/vmware/common"
	"github.com/mitchellh/packer/common"
	"github.com/mitchellh/packer/packer"
	"github.com/vmware/govmomi"
)

type Builder struct {
	config *Config
	runner multistep.Runner
}

func (b *Builder) Prepare(raws ...interface{}) ([]string, error) {
	c, warnings, err := NewConfig(raws...)
	if err != nil {
		return warnings, err
	}
	b.config = c

	return warnings, nil
}

func (b *Builder) Run(ui packer.Ui, hook packer.Hook, cache packer.Cache) (packer.Artifact, error) {
	parsedURL, err := url.Parse(b.config.VSphereURL)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse vSphere URL: %s", err)
	}
	client, err := govmomi.NewClient(nil, parsedURL, true)
	if err != nil {
		return nil, fmt.Errorf("Failed creating vSphere client: %s", err)
	}

	state := &multistep.BasicStateBag{}
	state.Put("config", b.config)
	state.Put("client", client)
	state.Put("hook", hook)
	state.Put("ui", ui)

	// TODO set up all steps
	steps := []multistep.Step{}

	if b.config.PackerDebug {
		b.runner = &multistep.DebugRunner{
			Steps:   steps,
			PauseFn: common.MultistepDebugFn(ui),
		}
	} else {
		b.runner = &multistep.BasicRunner{Steps: steps}
	}
	b.runner.Run(state)

	if rawErr, ok := state.GetOk("error"); ok {
		return nil, rawErr.(error)
	}

	if _, ok := state.GetOk(multistep.StateCancelled); ok {
		return nil, errors.New("Build was cancelled.")
	}

	if _, ok := state.GetOk(multistep.StateHalted); ok {
		return nil, errors.New("Build was halted.")
	}

	return vmwcommon.NewLocalArtifact(".")
}

func (b *Builder) Cancel() {
}
