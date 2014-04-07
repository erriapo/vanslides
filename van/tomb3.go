var task3    *io.GifWriter
task3.Run(vr, args)
if err := task3.Tombstone.Wait(); err != nil { // HL
    syscall.Exit(126)
}

