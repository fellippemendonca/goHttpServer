package auth

func Check(token string) bool {
	if token != "admin" {
		return false
	}
	return true
}
