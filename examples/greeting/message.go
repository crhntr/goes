package greeting

type Envelope struct {
	Message string
}

func Reverse(msg string) string {
	n := 0
	rune := make([]rune, len(msg))
	for _, r := range msg {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	return string(rune)
}
