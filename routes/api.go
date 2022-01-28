package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/controllers"
	"github.com/rullyafrizal/Final-Project-BDS-Sanbercode-Golang-Batch-31/middlewares"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	v1 := r.Group("/api/v1")
	{
		// Auth
		v1.POST("/auth/login", controllers.Login)
		v1.POST("/auth/register", controllers.Register)

		middlewaredAuthRoute := v1.Group("/auth")
		{
			middlewaredAuthRoute.Use(middlewares.JwtAuthMiddleware())
			middlewaredAuthRoute.GET("/me", controllers.Me)
			middlewaredAuthRoute.PUT("/update-password", controllers.UpdatePassword)
		}

		// Roles
		rolesRoute := v1.Group("/roles")
		{
			rolesRoute.Use(middlewares.JwtAuthMiddleware(), middlewares.AdminMiddleware())
			rolesRoute.GET("", controllers.IndexRoles)
			rolesRoute.POST("", controllers.StoreRole)
			rolesRoute.GET("/:id", controllers.ShowRole)
			rolesRoute.PUT("/:id", controllers.UpdateRole)
			rolesRoute.DELETE("/:id", controllers.DestroyRole)
		}

		// Users
		v1.GET("/users/:id", controllers.ShowUser)

		jwtAdminUsersRoute := v1.Group("/users")
		{
			jwtAdminUsersRoute.Use(middlewares.JwtAuthMiddleware(), middlewares.AdminMiddleware())
			jwtAdminUsersRoute.GET("", controllers.IndexUsers)
			jwtAdminUsersRoute.POST("", controllers.StoreUser)
		}

		middlewaredUsersRoute := v1.Group("/users")
		{
			middlewaredUsersRoute.Use(middlewares.JwtAuthMiddleware())
			middlewaredUsersRoute.PUT("/:id", controllers.UpdateUser)
			middlewaredUsersRoute.DELETE("/:id", controllers.DestroyUser)
		}

		// Posts
		v1.GET("/posts", controllers.IndexPosts)
		v1.GET("/posts/:id", controllers.ShowPost)

		middlewaredPostsRoute := v1.Group("/posts")
		{
			middlewaredPostsRoute.Use(middlewares.JwtAuthMiddleware())
			middlewaredPostsRoute.GET("/my", controllers.IndexMyPosts)
			middlewaredPostsRoute.GET("/:id/my", controllers.ShowMyPost)
			middlewaredPostsRoute.POST("", controllers.StorePost)
			middlewaredPostsRoute.PUT("/:id", controllers.UpdatePost)
			middlewaredPostsRoute.DELETE("/:id", controllers.DestroyPost)
			middlewaredPostsRoute.PATCH("/:id/publish", controllers.PublishPost)

			// Votes
			middlewaredPostsRoute.GET("/:id/votes/up", controllers.Upvote)
			middlewaredPostsRoute.GET("/:id/votes/down", controllers.Downvote)
		}

		// Review
		middlewaredReviewsRoute := middlewaredPostsRoute
		{
			middlewaredReviewsRoute.POST("/:id/reviews", controllers.StoreReview)
			middlewaredReviewsRoute.PUT("/:id/reviews/:review_id", controllers.UpdateReview)
			middlewaredReviewsRoute.DELETE("/:id/reviews/:review_id", controllers.DestroyReview)
		}

	}

	r.Run(":8080")
}
