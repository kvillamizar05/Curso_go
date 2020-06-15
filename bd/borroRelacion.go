package bd

import (
	"context"
	"time"

	"github.com/kvillamizar05/Curso_go/models"
)

/*BorroRelacion elimina el follow entre usuarios*/
func BorroRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("Curso_twittor")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
