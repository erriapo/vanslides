defer g.Tombstone.Done() // HL

select {
case <- g.Tombstone.Dying(): // HL
    return 
default:
    // continue on doing stuff
}

