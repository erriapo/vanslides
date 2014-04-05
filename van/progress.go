type Status struct {
    frame int32
    drop_frames int32
    progress string // HL
}

type MyHandler struct {
    notifyQueue *chan Status // HL
}
