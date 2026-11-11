package main

func isValidPassword(password string) bool {
	// check if password length in range from 5 to 12 characters (inclusive)
	if len(password) < 5 || len(password) > 12{
		return false
	} 

	hasUppercase := false
	hasNumber := false
	for _,w := range password{
		// check if password has ONE uppercase word
		if w >= 'A' && w <= 'Z' {
			hasUppercase = true
		}
		// check if password has ONE number
		if w >= '0' && w <= '9' {
			hasNumber = true
		}
	}

	if hasUppercase && hasNumber {
		return true
	} else {
		return false
	} 
}
