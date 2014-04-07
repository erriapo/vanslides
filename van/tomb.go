type GifWriter struct {
    Tombstone  tomb.Tomb // HL
}

type Tomb struct {
    m      sync.Mutex
    dying  chan struct{}  // HL
    dead   chan struct{}  // HL
    reason error
}
