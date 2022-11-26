package user

type UserRegistrationFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

func FormatUser(user User) UserRegistrationFormatter {
	formattedUser := UserRegistrationFormatter{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      user.PasswordHash,
	}

	return formattedUser
}
