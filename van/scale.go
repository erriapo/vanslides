type Scale uint16 // HL

const (
    // -scale width:_
    WidthOnly Scale = 1 << iota // HL

    // -scale _:height
    HeightOnly

    // -scale width:height
    WidthHeight
)
