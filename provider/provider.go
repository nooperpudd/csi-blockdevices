package provider

import (
	"github.com/rexray/gocsi"

	"github.com/nooperpudd/csi-blockdevices/service"
)

// New returns a new Mock Storage Plug-in Provider.
func New() gocsi.StoragePluginProvider {
	svc := service.New()
	return &gocsi.StoragePlugin{
		Controller: svc,
		Identity:   svc,
		Node:       svc,

		// BeforeServe allows the SP to participate in the startup
		// sequence. This function is invoked directly before the
		// gRPC server is created, giving the callback the ability to
		// modify the SP's interceptors, server options, or prevent the
		// server from starting by returning a non-nil error.
		BeforeServe: svc.BeforeServe,

		EnvVars: []string{
			// Enable serial volume access.
			gocsi.EnvVarSerialVolAccess + "=true",

			// Enable spec validation
			gocsi.EnvVarSpecReqValidation + "=true",

			// Provide the list of versions supported by this SP. The
			// specified versions will be:
			//     * Returned by GetSupportedVersions
			//     * Used to validate the Version field of incoming RPCs
			gocsi.EnvVarSupportedVersions + "=" + service.SupportedVersions,
		},
	}
}
