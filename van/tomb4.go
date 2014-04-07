task3.Run(vr, args)

time.Sleep(1 * time.Second)
log.Fatal(task3.Stop()) // HL
// internally calls task3.Tombstone.Kill(nil)

