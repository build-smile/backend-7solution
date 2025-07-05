package port

import (
	"context"
	"time"
)

type UserRepo interface {
	GetUser(req User, ctx context.Context) (*User, error)
	GetUsers(ctx context.Context) ([]User, error)
	CreateUser(req User, ctx context.Context) error
	UpdateUser(ctx context.Context, id string, name, email string) error
	DeleteUser(ctx context.Context, id string) error
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) error
}

type User struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	Name      string    `bson:"name" json:"name"`
	Email     string    `bson:"email" json:"email"`
	Password  string    `bson:"password" json:"-"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}
