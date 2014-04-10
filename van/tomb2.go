defer g.Tombstone.Done() // HL

for {
    select {
    case <- g.Tombstone.Dying(): // HL
        return 
    default:
        // do some more work
    }
}
