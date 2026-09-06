package bd

import (
	"context"
	"fmt"

	"github.com/nicos92/eltuiter/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var MongoCN *mongo.Client
var DatabaseName string

func ContarBD(ctx context.Context) error {
	user := ctx.Value(models.Key("user")).(string)
	password := ctx.Value(models.Key("password")).(string)
	host := ctx.Value(models.Key("host")).(string)
	connStr := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, password, host)

	var clienteOptions = options.Client().ApplyURI(connStr)
	client, err := mongo.Connect(ctx, clienteOptions)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Conexión exitosa con la BD")
	MongoCN = client
	DatabaseName := ctx.Value(models.Key("database")).(string)
	return nil
}

func BaseConectada() bool {
	err := MongoCN.Ping(context.TODO(), nil)

	return err == nil
}
