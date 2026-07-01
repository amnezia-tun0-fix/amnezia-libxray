package libXray

import (
	"github.com/amnezia-vpn/amnezia-libxray/nodep"
	"github.com/amnezia-vpn/amnezia-tun2socks/v2/filter"
)

// uidFilterAdapter bridges an Android-supplied UidFilterController to the
// tun2socks filter.PacketFilter mechanism.
type uidFilterAdapter struct {
	controller UidFilterController
}

func (a *uidFilterAdapter) Allow(network, srcIP string, srcPort int, dstIP string, dstPort int) bool {
	return a.controller.Allow(network, srcIP, srcPort, dstIP, dstPort)
}

// RegisterUidFilter enables Strict Split Tunneling: every new connection is
// gated through the controller before it enters the tunnel. Call this before
// StartTun2Socks, and only when the user toggle is ON.
func RegisterUidFilter(controller UidFilterController) string {
	filter.Set(&uidFilterAdapter{controller: controller})
	return nodep.WrapError(nil)
}

// UnregisterUidFilter disables Strict Split Tunneling, restoring the legacy
// allow-all behavior with no per-connection overhead.
func UnregisterUidFilter() string {
	filter.Set(nil)
	return nodep.WrapError(nil)
}
