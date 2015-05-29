package main

import (
	"github.com/mitchellh/packer/packer/plugin"
	"github.com/travis-ci/packer-builder-vmware-vsphere/builder/vmware/vsphere"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(vsphere.Builder))
	server.Serve()
}
