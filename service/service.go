package service

import (
	"context"
	"net"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/rexray/gocsi"
	csictx "github.com/rexray/gocsi/context"

	"github.com/nooperpudd/csi-blockdevices/core"
)

const (
	// Name is the name of the CSI plugin
	Name = "com.thecodeteam.blockdevices"

	// SupportedVersions is a list of supported CSI versions
	SupportedVersions = "0.1.0"

	defaultDevDir  = "/dev/disk/csi-blockdevices"
	defaultPrivDir = "/dev/disk/csi-bd-private"
)

// Manifest is the SP's manifest.
var Manifest = map[string]string{
	"url":    "https://github.com/rexray/csi-blockdevices",
	"semver": core.SemVer,
	"commit": core.CommitSha32,
	"formed": core.CommitTime.Format(time.RFC1123),
}

// Service is the CSI Network File System (NFS) service provider.
type Service interface {
	csi.ControllerServer
	csi.IdentityServer
	csi.NodeServer
	BeforeServe(context.Context, *gocsi.StoragePlugin, net.Listener) error
}

type service struct {
	DevDir  string
	privDir string
}

// New returns a new Service
func New() Service {

	return &service{}
}

func (s *service) BeforeServe(
	ctx context.Context, sp *gocsi.StoragePlugin, lis net.Listener) error {

	defer func() {
		fields := map[string]interface{}{
			"privatedir": s.privDir,
			"devicedir":  s.DevDir,
		}

		log.WithFields(fields).Infof("configured %s", Name)
	}()

	// remove the dir X_CSI_PRIVATE_MOUNT_DIR
	//if pd, ok := csictx.LookupEnv(ctx, gocsi.EnvVarPrivateMountDir); ok {
	//	s.privDir = pd
	//}
	if s.privDir == "" {
		s.privDir = defaultPrivDir
	}

	if pd, ok := csictx.LookupEnv(ctx, EnvBlockDevDir); ok {
		s.DevDir = pd
	}
	if s.DevDir == "" {
		s.DevDir = defaultDevDir
	}

	return nil
}
