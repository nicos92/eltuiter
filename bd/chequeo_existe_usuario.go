package bd

import (
	"context"

	"github.com/nicos92/eltuiter/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)

	col := db.Collection("usuario")
	condition := bson.M{"email": email}
	var result models.Usuario

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
