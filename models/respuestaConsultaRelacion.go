package models

/*RespuestaConsultaRelacion tiene el true o false que optiene al consultar la relacion entre dos usuarios*/
type RespuestaConsultaRelacion struct {
	Status bool `json:"status"`
}
