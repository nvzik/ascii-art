package asciiartfs

func CheckNewLines(s []string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != "" {
			return false
		}
	}
	return true
}
