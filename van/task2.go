var (
    task2    *io.Muxer
       vr    *io.VideoReader
)

wg := task2.Run(vr, args)
wg.Wait() // HL
if err := task2.Error(); err != nil { // HL
    syscall.Exit(126)
}
