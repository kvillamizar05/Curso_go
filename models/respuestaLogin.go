package models

/*RespuestaLogin tiene el token para el login*/
type RespuestaLogin struct {
	Token string `json:"token,omitempty"`
}
