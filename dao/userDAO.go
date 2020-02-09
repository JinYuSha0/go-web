package dao

import (
	"context"

	"go-web/models"
	"go-web/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserDAO interface {
	Register(account string, passwordEncrypted string) error
	IsExists(account string) (bool, error)
}

type userDAO struct {
	client     mongo.Client
	collection mongo.Collection
}

func NewUserDAO() UserDAO {
	config := utils.GetConfig()

	client := Connect()
	collection := client.Database(config.Mongo.Database).Collection("user")

	// 创建索引
	indexModel := []mongo.IndexModel{
		{
			Keys: bson.M{
				"account": 1,
			},
			Options: options.Index().SetUnique(true),
		},
	}
	collection.Indexes().CreateMany(
		context.Background(),
		indexModel,
	)

	return &userDAO{
		client:     *client,
		collection: *collection,
	}
}

func (d *userDAO) Register(account string, passwordEncrypted string) error {
	s := models.UserRegister{account, passwordEncrypted}
	_, err := d.collection.InsertOne(context.TODO(), s)
	return err
}

func (d *userDAO) IsExists(account string) (isExists bool, err error) {
	isExists = false

	count, err := d.collection.CountDocuments(context.TODO(), bson.M{
		"account": account,
	})

	if err != nil {
		return
	}

	if count > 0 {
		isExists = true
	}

	return
}
