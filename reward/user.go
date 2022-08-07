package reward

type User struct {
	email        string `validate:"required,email"`
	point        int    `validate:"required,gte=0"`
	registeredAt string `validate:"required,rfc3339"`
	updateAt     string `validate:"required,rfc3339"`
}
