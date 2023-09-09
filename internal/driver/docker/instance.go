package docker

import "github.com/srepio/cli/internal/metadata"

type Container struct {
	Id         string
	Name       string
	Image      string
	Ports      []metadata.Port
	Volumes    []metadata.Volume
	Privileged bool
}
