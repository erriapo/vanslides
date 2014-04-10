var (
    task2    *io.Muxer
       vr    *io.VideoReader
)

pWg := task2.Run(vr, args)
pWg.Wait() // HL
if err := task2.Error(); err != nil { // HL
    syscall.Exit(126)
}
