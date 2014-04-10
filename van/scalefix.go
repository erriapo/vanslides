func (s Scale) rangeCheck() error {  // HL
    switch s {
    case   WidthOnly: fallthrough
    case  HeightOnly: fallthrough
    case WidthHeight: return nil      
    default: return fmt.Errorf("Invalid Scale %d", s)
    }
}

func Rescale(s Scale) error {  // HL
    if err := s.rangeCheck(); err != nil { // HL
        return err
    }
    // do other work
    return nil
}

