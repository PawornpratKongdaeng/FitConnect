package domain

type Role string

const (
	RoleUser    Role = "user"
	RoleTrainer Role = "trainer"
	RoleAdmin   Role = "admin"
)

type User struct {
	ID        string `bson:"_id,omitempty" json:"id"`
	Name      string `bson:"name" json:"name"`
	Email     string `bson:"email" json:"email"`
	Password  string `bson:"password" json:"-"`
	Role      Role   `bson:"role" json:"role"`
	CreatedAt int64  `bson:"created_at" json:"createdAt"`
}
