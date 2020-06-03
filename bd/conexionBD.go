package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexion a la base de datos*/
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb://Kevin_0512:mern1234@kevincluster05-shard-00-00-mvi62.mongodb.net:27017,kevincluster05-shard-00-01-mvi62.mongodb.net:27017,kevincluster05-shard-00-02-mvi62.mongodb.net:27017/test?replicaSet=KevinCluster05-shard-0&ssl=true&authSource=admin")

/*ConectarBD es la funcion que permite conectar a la base de datos de mongoDb*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexion exitosa con la BD")
	return client
}

/*ChequeoConnection es el ping a la base de datos*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
