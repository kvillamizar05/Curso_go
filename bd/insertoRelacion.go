package bd

import (
	"context"
	"time"

	"github.com/kvillamizar05/Curso_go/models"
)

/*InsertoRelacion sirve para grabar la relacion entre usuarios followers*/
func InsertoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("Curso_twittor")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil

}
