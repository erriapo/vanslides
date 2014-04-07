done := make(chan struct{}) // HL
generator := even(done)
for j := 0; j < 5; j++ {
  i, ok := <- generator
  log(i, ok)
}

//done <- struct{}{}  
close(done)  // HL

i, ok := <- generator
log(i, ok)
