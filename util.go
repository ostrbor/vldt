package vldt

func short(v string) string {
	if len(v) > MaxStringLength {
		return v[:MaxStringLength] + "..."
	}
	return v
}

func printableASCII(s string) bool {
	for _, r := range s {
		if r < 32 || r > 126 {
			return false
		}
	}
	return true
}
