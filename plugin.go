//go:build linux && plugin
// +build linux,plugin

//go:generate go generate ./core

package main

import "C"

import (
	"github.com/nooperpudd/csi-blockdevices/provider"
	"github.com/nooperpudd/csi-blockdevices/service"
)

////////////////////////////////////////////////////////////////////////////////
//                              Go Plug-in                                    //
////////////////////////////////////////////////////////////////////////////////

// SupportedVersions is a space-delimited list of supported CSI versions.
var SupportedVersions = service.SupportedVersions

// ServiceProviders is an exported symbol that provides a host program
// with a map of the service provider names and constructors.
var ServiceProviders = map[string]func() interface{}{
	service.Name: func() interface{} { return provider.New() },
}
