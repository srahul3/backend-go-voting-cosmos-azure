package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/srahul3/govoting/cosmos"
)

// getAlbums responds with the list of all albums as JSON.
func getVoteData(c *gin.Context) {
	var v = cosmos.List()

	ignoreCORS(c)
	c.IndentedJSON(http.StatusOK, v)
}

// getAlbums responds with the list of all albums as JSON.
func vote(c *gin.Context) {
	teamid := c.Param("teamid")
	cosmos.VoteUp(teamid)

	ignoreCORS(c)

	m := make(map[string]string)
	m["status"] = "success"
	c.IndentedJSON(http.StatusOK, m)
}

// getAlbums responds with the list of all albums as JSON.
func getReset(c *gin.Context) {
	v1 := &(cosmos.VoteCandiate{ID: "1", Name: "Liverpool F.C.", LogoUrl: "https://kgo.googleusercontent.com/profile_vrt_raw_bytes_1587515361_10542.jpg", Votes: 0})
	v1.CreateIfDoesntExist()
	v2 := &cosmos.VoteCandiate{ID: "2", Name: "FC Barcelona", LogoUrl: "https://lh3.googleusercontent.com/OQZi4ckWAs7UrOlZEPefXZgJOcdJuSM5FSH9zqD5rMg6c2MOaxcKpV5IMrb1Tju98fWyNmcI33E4RGb0uC09Ej4W", Votes: 0}
	v2.CreateIfDoesntExist()
	v3 := &cosmos.VoteCandiate{ID: "3", Name: "Manchester United F.C.", LogoUrl: "https://upload.wikimedia.org/wikipedia/en/thumb/7/7a/Manchester_United_FC_crest.svg/640px-Manchester_United_FC_crest.svg.png", Votes: 0}
	v3.CreateIfDoesntExist()

	arr := []cosmos.VoteCandiate{*v1, *v2, *v3}
	c.IndentedJSON(http.StatusOK, arr)
}

func ignoreCORS(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.Default()

	router.GET("/voting", getVoteData)
	router.GET("/bootstrap", getReset)

	// react js doesnt support patch well hence using PUT
	router.PUT("/vote/:teamid", vote)

	router.Use(CORSMiddleware())
	router.Use(cors.Default())
	router.Run("localhost:8080")
}
