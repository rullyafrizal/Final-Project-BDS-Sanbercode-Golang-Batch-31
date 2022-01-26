package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/controllers"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	grp1 := r.Group("/api/v1")
	{
		// Auth
		grp1.POST("/auth/login", controllers.Login)

		// Roles
		grp1.GET("/roles", controllers.IndexRoles)
		grp1.POST("/roles", controllers.StoreRole)
		grp1.GET("/roles/:id", controllers.ShowRole)
		grp1.PUT("/roles/:id", controllers.UpdateRole)
		grp1.DELETE("/roles/:id", controllers.DestroyRole)

		// Users
		grp1.GET("/users", controllers.IndexUsers)
		grp1.POST("/users", controllers.StoreUser)
		grp1.GET("/users/:id", controllers.ShowUser)
		grp1.PUT("/users/:id", controllers.UpdateUser)
		grp1.DELETE("/users/:id", controllers.DestroyUser)

		// Posts
		grp1.GET("/posts", controllers.IndexPosts)
		grp1.POST("/posts", controllers.StorePost)
		grp1.GET("/posts/:id", controllers.ShowPost)
		grp1.PUT("/posts/:id", controllers.UpdatePost)
		grp1.DELETE("/posts/:id", controllers.DestroyPost)
		grp1.PATCH("/posts/:id/publish", controllers.PublishPost)
	}

	r.Run(":8080")
}
