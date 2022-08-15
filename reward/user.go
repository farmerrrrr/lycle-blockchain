package reward

type User struct {
	Email        string `validate:"required,email"`
	Point        int    `validate:"required,gte=0"`
	RegisteredAt string `validate:"required,rfc3339"`
	UpdatedAt     string `validate:"required,rfc3339"`
}
