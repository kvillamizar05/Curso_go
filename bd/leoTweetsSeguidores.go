package bd

import (
	"context"
	"time"

	"github.com/kvillamizar05/Curso_go/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*LeoTweetsSeguidores sirve para leer todos los tweets de las cuentas que sigo*/
func LeoTweetsSeguidores(ID string, pagina int) ([]models.DevuelvoTweetSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("Curso_twittor")
	col := db.Collection("relacion")

	skip := (pagina - 1) * 20

	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid",
			"foreingField": "userid",
			"as":           "Tweet",
		}})
	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"tweet.fecha": -1}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, condiciones)

	var result []models.DevuelvoTweetSeguidores
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}

	return result, true

}
