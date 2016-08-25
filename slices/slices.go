package slices

// ContainsString reports whether the needle is contained in the haystack or not.
func ContainsString(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}
