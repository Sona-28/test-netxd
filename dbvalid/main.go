package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const ConnectionString = "mongodb+srv://sonas:sona1234@cluster0.buvinnz.mongodb.net/admin?retryWrites=true&replicaSet=atlas-zmtpvi-shard-0&readPreference=primary&srvServiceName=mongodb&connectTimeoutMS=10000&authSource=admin&authMechanism=SCRAM-SHA-1"

type Sample struct{
	Paragraph string `json:"paragraph" bson:"paragraph"`
}

func main() {
	engine := gin.Default()
	engine.GET("/", handler)
	engine.Run(":8080")
}

func handler(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
    if tokenString == "" {
		fmt.Println("error1")
        return
    }
    tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
    secretKey := []byte("your-256-bit-secret")
    claims, err := validateAndDecodeToken(tokenString, secretKey)
    if err != nil {
		fmt.Println("error2")

        return
    }
    paragraph := claims["paragraph"].(string)
    fmt.Printf("Authenticated as: %s\n", paragraph)
	sample := Sample{
		Paragraph: paragraph,
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoConnection := options.Client().ApplyURI(ConnectionString)
	mongoClient, err := mongo.Connect(ctx, mongoConnection)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println("Database Connected")
	collection := mongoClient.Database("k6demo").Collection("token")
	result, err := collection.InsertOne(ctx, sample)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, result.InsertedID)
}

func validateAndDecodeToken(tokenString string, secretKey []byte) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("invalid signing method")
        }
        return secretKey, nil
    })

    if err != nil {
        return nil, err
    }

    // Verify if the token is valid
    if !token.Valid {
        return nil, fmt.Errorf("invalid token")
    }

    // Extract the claims from the token
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		
        return claims, nil
    }

    return nil, fmt.Errorf("failed to extract claims")
}
