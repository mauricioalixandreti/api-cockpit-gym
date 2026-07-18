package Models

type User struct {
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Phone    string `bson:"phone"`
	Password string `bson:"password"`
}
