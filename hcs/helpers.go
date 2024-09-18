//go:build windows

package hcsshim

import (
	"github.com/Microsoft/hcsshim/internal/hcs/schema1"
)

// ContainerProperties holds the properties for a container and the processes running in that container
type ContainerProperties = schema1.ContainerProperties

// NetworkStats holds the network statistics for a container
type NetworkStats = schema1.NetworkStats
