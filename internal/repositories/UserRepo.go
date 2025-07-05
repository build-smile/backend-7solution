package repositories

import (
	"context"
	"fmt"
	"github.com/build-smile/backend-7solution/infrastructure"
	"github.com/build-smile/backend-7solution/internal/core/port"
	"github.com/build-smile/backend-7solution/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type UserRepo struct {
	m *mongo.Database
}

func NewUserRepo() port.UserRepo {
	return &UserRepo{
		m: infrastructure.MongoDB,
	}
}

func (r *UserRepo) DeleteUser(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(infrastructure.CFG.MongoDB.ExecuteTimeoutMilli)*time.Millisecond)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("invalid user ID format: %s", id))
	}
	col := r.m.Collection("user")

	filter := bson.M{"_id": objID}

	result, err := col.DeleteOne(ctx, filter)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("failed to delete user: %v", err))
	}

	if result.DeletedCount == 0 {
		return utils.NewCustomError(http.StatusNotFound, fmt.Sprintf("user with ID %s not found", id))
	}

	return nil
}

func (r *UserRepo) UpdateUser(ctx context.Context, id string, name, email string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(infrastructure.CFG.MongoDB.ExecuteTimeoutMilli)*time.Millisecond)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("invalid user ID format: %s", id))
	}
	col := r.m.Collection("user")

	filter := bson.M{"_id": objID}

	updateFields := bson.M{}
	if name != "" {
		updateFields["name"] = name
	}
	if email != "" {
		updateFields["email"] = email
	}

	if len(updateFields) == 0 {
		return utils.NewCustomError(http.StatusBadRequest, "no fields to update provided")
	}

	update := bson.M{"$set": updateFields}

	result, err := col.UpdateOne(ctx, filter, update)
	if err != nil {
		return utils.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("failed to update user: %v", err))
	}

	if result.MatchedCount == 0 {
		return utils.NewCustomError(http.StatusNotFound, fmt.Sprintf("user with ID %s not found", id))
	}

	return nil
}

func (r *UserRepo) GetUser(req port.User, ctx context.Context) (*port.User, error) {
	var user port.User

	// Set timeout only if context doesn't already have one
	ctx, cancel := context.WithTimeout(ctx, time.Duration(infrastructure.CFG.MongoDB.ExecuteTimeoutMilli)*time.Millisecond)
	defer cancel()

	col := r.m.Collection("user")
	err := col.FindOne(ctx, bson.M{"name": req.Name}).Decode(&user)
	return &user, err
}

func (r *UserRepo) GetUsers(ctx context.Context) ([]port.User, error) {
	// Set timeout only if context doesn't already have one
	ctx, cancel := context.WithTimeout(ctx, time.Duration(infrastructure.CFG.MongoDB.ExecuteTimeoutMilli)*time.Millisecond)
	defer cancel()

	col := r.m.Collection("user")

	cursor, err := col.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to find users: %w", err)
	}
	defer cursor.Close(ctx)

	var users []port.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, fmt.Errorf("failed to decode users: %w", err)
	}

	return users, nil
}

func (r *UserRepo) CreateUser(req port.User, ctx context.Context) error {
	// Create user collection if it doesn't exist
	c, cancel := context.WithTimeout(ctx, time.Duration(infrastructure.CFG.MongoDB.ExecuteTimeoutMilli)*time.Millisecond)
	defer cancel()
	coll := r.m.Collection("user")
	_, err := coll.InsertOne(c, req)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares a plain password with a hashed one
func (r *UserRepo) CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}
