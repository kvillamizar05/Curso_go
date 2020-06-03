package middlew

import (
	"net/http"

	"github.com/kvillamizar05/Curso_go/bd"
)

/*ChequeoBD es el middlewaresirve para revisar la conexion xon la base de datos*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexión perdida con la base de datos.", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
