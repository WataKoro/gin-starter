package builder

import (
    "gin-starter/app"
    "gin-starter/config"
    loanRepo "gin-starter/modules/loans/v1/repository"
    loanService "gin-starter/modules/loans/v1/service"
    "gin-starter/sdk/gcs"
    "github.com/gin-gonic/gin"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/gomodule/redigo/redis"
    "gorm.io/gorm"
)

// BuildLoanHandler builds loan handler
func BuildLoanHandler(cfg config.Config, router *gin.Engine, db *gorm.DB, redisPool *redis.Pool, awsSession *session.Session) {
    // Repository
    loanRepository := loanRepo.NewLoanRepository(db)

    cloudStorage := gcs.NewGoogleCloudStorage(cfg)

    // Service
    loanCreator := loanService.NewLoanCreator(cfg, loanRepository)
    loanFinder := loanService.NewLoanFinder(cfg, loanRepository)
    loanUpdater := loanService.NewLoanUpdater(cfg, loanRepository)

    // Handler
    app.LoanFinderHTTPHandler(cfg, router, loanFinder)
    app.LoanCreatorHTTPHandler(cfg, router, loanCreator, cloudStorage)
    app.LoanUpdaterHTTPHandler(cfg, router, loanUpdater, loanFinder, cloudStorage)
}
