package bd

import (
	"context"
	"time"

	"github.com/kvillamizar05/Curso_go/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario sirve para verficar que no se vuelva a crear un usuario existente*/
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("Curso_twittor")
	col := db.Collection("users")

	condicion := bson.M{"email": email}
	var resultado models.Usuario
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID
}
