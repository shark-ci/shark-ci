package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Identity struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty"`
	ServiceName string               `bson:"serviceName,omitempty"`
	Username    string               `bson:"username,omitempty"`
	Token       OAuth2Token          `bson:"token,omitempty"`
	Repos       []primitive.ObjectID `bson:"repos,omitempty"`
}

// This structure must be compatible with https://pkg.go.dev/golang.org/x/oauth2#Token
type OAuth2Token struct {
	AccessToken  string     `json:"access_token" bson:"accessToken"`
	TokenType    string     `json:"token_type,omitempty" bson:"tokenType,omitempty"`
	RefreshToken string     `json:"refresh_token,omitempty" bson:"refreshToken,omitempty"`
	Expiry       *time.Time `json:"expiry,omitempty" bson:"expiry,omitempty"`
}

func GetOrCreateIdentity(ctx context.Context, identity *Identity, user *User) (*Identity, error) {
	filter := bson.D{
		{Key: "serviceName", Value: identity.ServiceName},
		{Key: "username", Value: identity.Username},
	}
	opts := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	err := Identities.FindOneAndUpdate(ctx, filter, bson.D{{Key: "$set", Value: identity}}, opts).Decode(identity)
	if err != nil {
		return nil, err
	}

	if user == nil {
		user, err = CreateUser(ctx)
		// TODO: Maybe delete created identity if user cannot be created
		if err != nil {
			return nil, err
		}
	}

	filter = bson.D{
		{Key: "_id", Value: user.ID},
		{Key: "identities", Value: bson.D{
			{Key: "$ne", Value: identity.ID},
		}},
	}
	update := bson.D{
		{Key: "$push", Value: bson.D{
			{Key: "identities", Value: identity.ID},
		}},
	}
	_, err = Users.UpdateOne(ctx, filter, update)
	return identity, err
}

func GetIdentityByUser(ctx context.Context, user *User, service string) (*Identity, error) {
	var identity Identity
	filter := bson.D{
		{Key: "_id", Value: bson.D{
			{Key: "$in", Value: user.Identities},
		}},
		{Key: "serviceName", Value: service},
	}
	err := Identities.FindOne(ctx, filter).Decode(&identity)
	if err != nil {
		return nil, err
	}

	return &identity, nil
}

func GetIdentityByUsername(ctx context.Context, username string, service string) (*Identity, error) {
	var identity Identity
	filter := bson.D{
		{Key: "username", Value: username},
		{Key: "serviceName", Value: service},
	}
	err := Identities.FindOne(ctx, filter).Decode(&identity)
	if err != nil {
		return nil, err
	}

	return &identity, nil
}