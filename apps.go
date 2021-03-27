package main

import (
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/sahlannasution/xnews-xapiens-backend/config"
	"github.com/sahlannasution/xnews-xapiens-backend/middlewares"
	"github.com/sahlannasution/xnews-xapiens-backend/migrator"
	"github.com/sahlannasution/xnews-xapiens-backend/resolvers"
	"github.com/sahlannasution/xnews-xapiens-backend/seeder"

	"github.com/gin-gonic/gin"
)

func main() {
	dbPG := config.Connection() // db Connection
	StrDB := middlewares.StrDB{DB: dbPG}
	ResolverDB := resolvers.ResolverDB{DB: dbPG}
	migrator.Migrations(dbPG) // migrate tables
	seeder.SeederUser(dbPG)   // seed User Data

	route := gin.Default()

	// route.POST("/", func(c *gin.Context) {
	// 	// fmt.Println("Welcome to X-News!")
	// 	c.JSON(http.StatusOK, "Welcome to X-News!") // Send Response
	// })
	/* User Register */
	route.POST("/register", ResolverDB.Register)
	/* User Signin */
	route.POST("/login", StrDB.MiddleWare().LoginHandler)
	route.NoRoute(StrDB.MiddleWare().MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// Define route
	// route.POST("/", StrDB.MiddleWare().MiddlewareFunc(), func(c *gin.Context) {
	// 	// Struvt Query
	// 	type Query struct {
	// 		Query string `json:"query"`
	// 	}

	// 	var query Query

	// 	c.Bind(&query) // Get query params

	// 	result := routes.ExecuteQuery(query.Query, schema.Schema) // Run Query
	// 	c.JSON(http.StatusOK, result)                             // Send Response
	// })
	route.Run()
}
