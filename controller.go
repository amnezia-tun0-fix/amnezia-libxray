package libXray

type DialerController interface {
	ProtectFd(int) bool
}

// UidFilterController is implemented on the Android side to back the "Strict
// Split Tunneling" feature. For each new connection entering the tunnel, Allow
// receives the 5-tuple (src = the originating app endpoint, dst = destination)
// and returns whether the connection is permitted — typically by resolving the
// owning app UID via getConnectionOwnerUid and applying the split-tunnel
// policy. It is called on the tunnel's forwarder goroutine, so implementations
// should cache UID lookups and be safe for concurrent use.
type UidFilterController interface {
	Allow(network string, srcIp string, srcPort int, dstIp string, dstPort int) bool
}
