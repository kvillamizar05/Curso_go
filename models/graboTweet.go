package models

import "time"

/*GraboTweet es el formato que tendra cada tweet*/
type GraboTweet struct {
	UserID  string    `bson:"userid" json:"userid,omitempty`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty`
}
