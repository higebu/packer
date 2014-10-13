package iso

import (
	vmwcommon "github.com/mitchellh/packer/builder/vmware/common"
)

type RemoteDriver interface {
	vmwcommon.Driver

	// UploadISO uploads a local ISO to the remote side and returns the
	// new path that should be used in the VMX along with an error if it
	// exists.
	UploadISO(string, string, string) (string, error)

	// Adds a VM to inventory specified by the path to the VMX given.
	Register(string) error

	// Destroies a VM
	Destroy() error

	// Checks if the VM is destroied.
	IsDestroied() (bool, error)
}
