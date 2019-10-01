package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"GetPoetry/getPoetry"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongodb(mUrl string) mongo.Client {
	clientOptions := options.Client().ApplyURI(mUrl)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return *client
}

func CloseMongodb(client mongo.Client) {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func main() {
	var url string
	client := ConnectToMongodb("mongodb://localhost:27017")

	for i := 1; i <= 685; i++ {
		url = fmt.Sprintf("https://www.gswen.cn/ancient/0/0/16/0/0/%d/", i)
		resp, _ := http.Get(url)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		getPoetry.GetPoetryToFile(string(body), "poetry.data")
		getPoetry.GetPoetryToMongo(string(body), client)
	}
	CloseMongodb(client)
}
