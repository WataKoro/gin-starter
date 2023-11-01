package builder

import (
    "gin-starter/app"
    "gin-starter/config"
    bookRepo "gin-starter/modules/master/v1/repository"
    bookService "gin-starter/modules/master/v1/service"
    "gin-starter/sdk/gcs"
    "github.com/gin-gonic/gin"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/gomodule/redigo/redis"
    "gorm.io/gorm"
)

// BuildBookHandler builds book handler
func BuildBookHandler(cfg config.Config, router *gin.Engine, db *gorm.DB, redisPool *redis.Pool, awsSession *session.Session) {
    // Repository
    bookRepository := bookRepo.NewBookRepository(db)

    cloudStorage := gcs.NewGoogleCloudStorage(cfg)

    // Service
    bookCreator := bookService.NewBookCreator(cfg, bookRepository)
    bookFinder := bookService.NewBookFinder(cfg, bookRepository)
    bookDeleter := bookService.NewBookDeleter(cfg, bookRepository)
    bookUpdater := bookService.NewBookUpdater(cfg, bookRepository)
    
    // Handler
    app.BookFinderHTTPHandler(cfg, router, bookFinder)
    app.BookCreatorHTTPHandler(cfg, router, bookCreator, bookFinder, cloudStorage)
    app.BookDeleterHTTPHandler(cfg, router, bookDeleter, cloudStorage)
    app.BookUpdaterHTTPHandler(cfg, router, bookUpdater, bookFinder, cloudStorage)
}
