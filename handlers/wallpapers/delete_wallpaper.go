package wallpapers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func DeleteWallpaper(w http.ResponseWriter, r *http.Request) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://tushardbanduser:Tushartxt11223344@cluster0.at7gz.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	collection := client.Database("myFirstDatabase").Collection("wallpapers")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	vars := mux.Vars(r)
	id := vars["id"]
	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
		"count":  strconv.Itoa(int(result.DeletedCount)),
	})
}
