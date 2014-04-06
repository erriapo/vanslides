type Status struct {
    frame int32
    drop_frames int32
    progress string // HL
}

// Has a send-only channel
type MyHandler struct {
    pings chan<- Status // HL
}
