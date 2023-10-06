package builder

import (
    "gin-starter/app"
    "gin-starter/config"
    bookRepo "gin-starter/modules/book/v1/repository"
    bookService "gin-starter/modules/book/v1/service"
    "gin-starter/sdk/gcs"
    "github.com/gin-gonic/gin"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/gomodule/redigo/redis"
    "gorm.io/gorm"
)

// BuildBookHandler builds book handler
// starting from handler down to repository or tool.
func BuildBookHandler(cfg config.Config, router *gin.Engine, db *gorm.DB, redisPool *redis.Pool, awsSession *session.Session) {
    // Repository
    br := bookRepo.NewBookRepository(db)
    // Add other book-related repositories as needed.

    cloudStorage := gcs.NewGoogleCloudStorage(cfg)

    // Service
    bc := bookService.NewBookCreator(cfg, br)
    bf := bookService.NewBookFinder(cfg, br)
    bd := bookService.NewBookDeleter(cfg, br) // Pass the appropriate dependency here.
    bu := bookService.NewBookUpdater(cfg, br) // Pass the appropriate dependency here.
    
    // Handler
    app.BookFinderHTTPHandler(cfg, router, bf)
    app.BookCreatorHTTPHandler(cfg, router, bc, bf, cloudStorage)
    app.BookDeleterHTTPHandler(cfg, router, bd, cloudStorage)
    app.BookUpdaterHTTPHandler(cfg, router, bu, bf, cloudStorage)
    // Ensure you have the appropriate service and handler functions for updating and deleting books.
}
