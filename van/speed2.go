var i interface{}
fmt.Println(unsafe.Sizeof(i)) // 8  // HL
var s struct{}
fmt.Println(unsafe.Sizeof(s)) // 0  // HL

var sarr [1000000000]struct{}
var iarr [100000000]interface{}
fmt.Println(unsafe.Sizeof(sarr)) // 0   // HL
fmt.Println(unsafe.Sizeof(iarr)) // 800000000 // HL
