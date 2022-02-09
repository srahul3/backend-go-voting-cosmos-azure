package main

import (
    "net/http"

    "github.com/gin-gonic/gin"

    //"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"

    //"backend-go-voting-cosmos-azure/cosmos"
    //"./db"
    //"github.com/srahul3/govoting/cosmos"
    "github.com/srahul3/govoting/cosmos"
    //"github.com/srahul3/govoting/model"
    
)

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

// getAlbums responds with the list of all albums as JSON.
func getVoteData(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

// getAlbums responds with the list of all albums as JSON.
func getReset(c *gin.Context) {
    //var voting = [] *cosmos.VoteCandiate {
        //{ID: "1", Name: "Liverpool F.C.", LogoUrl: "https://upload.wikimedia.org/wikipedia/en/thumb/0/0c/Liverpool_FC.svg/640px-Liverpool_FC.svg.png", Votes: 0},
        //{id: "1", name: "FC Barcelona", logoUrl: "https://upload.wikimedia.org/wikipedia/en/thumb/0/0c/Liverpool_FC.svg/640px-Liverpool_FC.svg.png", votes: 0},
        //{id: "1", name: "Manchester United F.C.", logoUrl: "https://upload.wikimedia.org/wikipedia/en/thumb/0/0c/Liverpool_FC.svg/640px-Liverpool_FC.svg.png", votes: 0},
    //}

    v := &cosmos.VoteCandiate{ID: "1", Name: "Liverpool F.C.", LogoUrl: "https://upload.wikimedia.org/wikipedia/en/thumb/0/0c/Liverpool_FC.svg/640px-Liverpool_FC.svg.png", Votes: 0}
    v.Create()
    v1 := &cosmos.VoteCandiate{ID: "2", Name: "FC Barcelona", LogoUrl: "https://lh3.googleusercontent.com/OQZi4ckWAs7UrOlZEPefXZgJOcdJuSM5FSH9zqD5rMg6c2MOaxcKpV5IMrb1Tju98fWyNmcI33E4RGb0uC09Ej4W", Votes: 0}
    v1.Create()
    v2 := &cosmos.VoteCandiate{ID: "3", Name: "Manchester United F.C.", LogoUrl: "https://upload.wikimedia.org/wikipedia/en/thumb/7/7a/Manchester_United_FC_crest.svg/640px-Manchester_United_FC_crest.svg.png", Votes: 0}
    v2.Create()

    c.IndentedJSON(http.StatusOK, albums)
}


func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/vote", getVoteData)
    router.GET("/reset", getReset)
    router.Run("localhost:8080")
}