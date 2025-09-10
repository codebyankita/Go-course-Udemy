// package mongodb

// import (
// 	"context"
// 	"grpcapi/pkg/utils"

// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func CreateMongoClient() (*mongo.Client, error) {
// 	ctx := context.Background()
// 	// client, err := mongo.Connect(ctx, options.Client().ApplyURI("username:password@mongodb://localhost:27017"))
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
// 	if err != nil {
// 		return nil, utils.ErrorHandler(err, "Unable to connect to database")
// 	}

// 	err = client.Ping(ctx, nil)
// 	if err != nil {
// 		return nil, utils.ErrorHandler(err, "Unable to ping database")
// 	}

// 	// log.Println("Connected to MongoDB")
// 	return client, nil
// }
package mongodb

import (
	"context"
	"fmt"
	"grpcapi/pkg/utils"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client   *mongo.Client
	Database *mongo.Database
)

// CreateMongoClient initializes and connects MongoDB client
func CreateMongoClient() (*mongo.Client, error) {
	ctx := context.Background()

	// Get URI and DB from environment
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = "mongodb://127.0.0.1:27017" // fallback default
	}
	dbName := os.Getenv("MONGO_DB")
	if dbName == "" {
		dbName = "grpc_api" // fallback default
	}

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, utils.ErrorHandler(err, "Unable to connect to database")
	}

	// Ping DB to test connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, utils.ErrorHandler(err, "Unable to ping database")
	}

	fmt.Println("✅ Connected to MongoDB at", uri, " DB:", dbName)

	// Assign global vars
	Client = client
	Database = client.Database(dbName)

	return client, nil
}
