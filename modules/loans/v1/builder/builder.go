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
    lr := loanRepo.NewLoanRepository(db)

    cloudStorage := gcs.NewGoogleCloudStorage(cfg)

    // Service
    lc := loanService.NewLoanCreator(cfg, lr)
    lf := loanService.NewLoanFinder(cfg, lr) 
    lu := loanService.NewLoanUpdater(cfg, lr)

    // Handler
    app.LoanFinderHTTPHandler(cfg, router, lf)
    app.LoanCreatorHTTPHandler(cfg, router, lc, cloudStorage)
    app.LoanUpdaterHTTPHandler(cfg, router, lu, lf, cloudStorage)
}
