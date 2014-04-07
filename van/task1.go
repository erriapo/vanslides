var (
    task1    *io.FrameGenerator
       vr    *io.VideoReader
)

reply := task1.Run(vr, args)
if err := <- reply; err != nil { // HL
    syscall.Exit(126)
}

