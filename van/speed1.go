var speeds = map[string]interface{}{  // HL
    "veryfast": nil,
    "faster":   nil,
    "placebo":  nil,
    "slower":   nil,
    "veryslow": nil,
}

// set membership check
if _, ok := speeds[arg]; !ok {  // HL
    // report error
}
