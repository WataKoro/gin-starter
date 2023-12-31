package app

import (
	"gin-starter/common/interfaces"
	"gin-starter/config"
	"gin-starter/middleware"
	"gin-starter/response"

	// Auth
	authhandlerv1 "gin-starter/modules/auth/v1/handler"
	authservicev1 "gin-starter/modules/auth/v1/service"
	
	// Master
	masterhandlerv1 "gin-starter/modules/master/v1/handler"
	masterservicev1 "gin-starter/modules/master/v1/service"
	
	// Notification
	notificationhandlerv1 "gin-starter/modules/notification/v1/handler"
	notificationservicev1 "gin-starter/modules/notification/v1/service"
	
	// User
	userhandlerv1 "gin-starter/modules/user/v1/handler"
	userservicev1 "gin-starter/modules/user/v1/service"
	
	// Loan
	loanhandlerv1 "gin-starter/modules/loans/v1/handler"
    loanservicev1 "gin-starter/modules/loans/v1/service"
	
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
)

// DeprecatedAPI is a handler for deprecated APIs
func DeprecatedAPI(c *gin.Context) {
	c.JSON(http.StatusForbidden, response.ErrorAPIResponse(http.StatusForbidden, "this version of api is deprecated. please use another version."))
	c.Abort()
}

// DefaultHTTPHandler is a handler for default APIs
func DefaultHTTPHandler(cfg config.Config, router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, response.ErrorAPIResponse(http.StatusNotFound, "invalid route"))
		c.Abort()
	})
}

// AuthHTTPHandler is a handler for auth APIs
func AuthHTTPHandler(cfg config.Config, router *gin.Engine, auc authservicev1.AuthUseCase) {
	hnd := authhandlerv1.NewAuthHandler(auc)
	v1 := router.Group("/v1")
	{
		v1.POST("/user/login", hnd.Login)
		v1.POST("/cms/login", hnd.LoginCMS)
	}
}

// NotificationFinderHTTPHandler is a handler for notification APIs
func NotificationFinderHTTPHandler(cfg config.Config, router *gin.Engine, cf notificationservicev1.NotificationFinderUseCase, nu notificationservicev1.NotificationUpdaterUseCase) {
	hnd := notificationhandlerv1.NewNotificationFinderHandler(cf, nu)
	v1 := router.Group("/v1")

	v1.Use(middleware.Auth(cfg))
	{
		v1.GET("/user/notifications", hnd.GetNotification)
		v1.GET("/user/notification/count", hnd.CountUnreadNotifications)
	}
}

// NotificationCreatorHTTPHandler is a handler for notification APIs
func NotificationCreatorHTTPHandler(cfg config.Config, router *gin.Engine, cf notificationservicev1.NotificationCreatorUseCase) {
	hnd := notificationhandlerv1.NewNotificationCreatorHandler(cf)
	v1 := router.Group("/v1")
	{
		v1.POST("/cms/notification", hnd.CreateNotification)
	}
}

// NotificationUpdaterHTTPHandler is a handler for notification APIs
func NotificationUpdaterHTTPHandler(cfg config.Config, router *gin.Engine, cf notificationservicev1.NotificationUpdaterUseCase) {
	hnd := notificationhandlerv1.NewNotificationUpdaterHandler(cf)
	v1 := router.Group("/v1")

	v1.Use(middleware.Auth(cfg))
	{
		v1.PUT("/user/notification/set", hnd.RegisterUnregisterPlayerID)
		v1.PUT("/user/notification/read", hnd.UpdateReadNotification)
	}
}

// MasterCreatorHTTPHandler is a handler for master APIs
func MasterCreatorHTTPHandler(cfg config.Config, router *gin.Engine, mc masterservicev1.MasterCreatorUseCase, cloudStorage interfaces.CloudStorageUseCase) {
	_ = masterhandlerv1.NewMasterCreatorHandler(mc, cloudStorage)
	v1 := router.Group("/v1")

	v1.Use(middleware.Auth(cfg))
	v1.Use(middleware.Admin(cfg))
}

// MasterFinderHTTPHandler is a handler for master APIs
func MasterFinderHTTPHandler(cfg config.Config, router *gin.Engine, mf masterservicev1.MasterFinderUseCase) {
	hnd := masterhandlerv1.NewMasterFinderHandler(mf)
	v1 := router.Group("/v1")
	{
		v1.GET("/provinces", hnd.GetProvinces)
		v1.GET("/regencies/:province_id", hnd.GetRegenciesByProvinceID)
		v1.GET("/districts/:regency_id", hnd.GetDistrictsByRegencyID)
		v1.GET("/villages/:district_id", hnd.GetVillagesByDistrictID)
	}
}

// UserFinderHTTPHandler is a handler for user APIs
func UserFinderHTTPHandler(cfg config.Config, router *gin.Engine, cf userservicev1.UserFinderUseCase) {
	hnd := userhandlerv1.NewUserFinderHandler(cf)
	v1 := router.Group("/v1")
	{
		v1.GET("/user/forgot-password/profile/:token", hnd.GetUserByForgotPasswordToken)
	}

	v1.Use(middleware.Auth(cfg))
	{
		v1.GET("/user/profile", hnd.GetUserProfile)
	}

	v1.Use(middleware.Admin(cfg))
	{
		v1.GET("/cms/profile", hnd.GetAdminProfile)
		v1.GET("/cms/admin/list", hnd.GetAdminUsers)
		v1.GET("/cms/admin/detail/:id", hnd.GetAdminUserByID)
		v1.GET("/cms/user/list", hnd.GetUsers)
		v1.GET("/cms/user/detail/:id", hnd.GetUserByID)
		v1.GET("/cms/roles", hnd.GetUserRoles)
		v1.GET("/cms/role/:id", hnd.GetUserRoles)
		v1.GET("/cms/permission", hnd.GetPermissions)
		v1.GET("/cms/user/permission", hnd.GetUserPermissions)
	}
}

// UserCreatorHTTPHandler is a handler for user APIs
func UserCreatorHTTPHandler(cfg config.Config, router *gin.Engine, uc userservicev1.UserCreatorUseCase, uf userservicev1.UserFinderUseCase, cloudStorage interfaces.CloudStorageUseCase) {
	hnd := userhandlerv1.NewUserCreatorHandler(uc, cloudStorage)
	v1 := router.Group("/v1")

	{
		v1.POST("/user/register", hnd.RegisterUser)
	}

	v1.Use(middleware.Auth(cfg))
	v1.Use(middleware.Admin(cfg))
	{
		v1.POST("/cms/user", hnd.CreateUser)
		v1.POST("/cms/admin/user", hnd.CreateAdmin)
		v1.POST("/cms/permission", hnd.CreatePermission)
		v1.POST("/cms/role", hnd.CreateUserRole)
	}
}

// UserUpdaterHTTPHandler is a handler for user APIs
func UserUpdaterHTTPHandler(cfg config.Config, router *gin.Engine, uu userservicev1.UserUpdaterUseCase, uf userservicev1.UserFinderUseCase, cloudStorage interfaces.CloudStorageUseCase) {
	hnd := userhandlerv1.NewUserUpdaterHandler(uu, uf, cloudStorage)
	v1 := router.Group("/v1")
	{
		v1.PUT("/user/forgot-password/request", hnd.ForgotPasswordRequest)
		v1.PUT("/user/forgot-password", hnd.ForgotPassword)
	}

	v1.Use(middleware.Auth(cfg))
	{
		v1.PUT("/user/profile", hnd.UpdateUser)
		v1.PUT("/user/password", hnd.ChangePassword)
	}

	v1.Use(middleware.Admin(cfg))
	{
		v1.PUT("/user/profile/:id", hnd.UpdateUser)
		v1.PUT("/cms/admin/:id", hnd.UpdateAdmin)
		v1.PUT("/cms/user/activate/:id", hnd.ActivateDeactivateUser)
		v1.PUT("/cms/role/:id", hnd.UpdateRole)
		v1.PUT("/cms/permission/:id", hnd.UpdatePermission)
	}
}

// UserDeleterHTTPHandler is a handler for user APIs
func UserDeleterHTTPHandler(cfg config.Config, router *gin.Engine, ud userservicev1.UserDeleterUseCase, cloudStorage interfaces.CloudStorageUseCase) {
	hnd := userhandlerv1.NewUserDeleterHandler(ud, cloudStorage)
	v1 := router.Group("/v1")

	v1.Use(middleware.Auth(cfg))
	v1.Use(middleware.Admin(cfg))
	{
		v1.DELETE("/cms/user/:id", hnd.DeleteUsers)
		v1.DELETE("/cms/admin/:id", hnd.DeleteAdmin)
		v1.DELETE("/cms/role/:id", hnd.DeleteRole)
	}
}

// BookCreatorHTTPHandler is a handler for book APIs
func BookCreatorHTTPHandler(cfg config.Config, router *gin.Engine, uc masterservicev1.BookCreatorUseCase, uf masterservicev1.BookFinderUseCase, cloudStorage interfaces.CloudStorageUseCase) {
	hnd := masterhandlerv1.NewBookCreatorHandler(uc, cloudStorage)
	v1 := router.Group("/v1")

	v1.Use(middleware.Auth(cfg))
	v1.Use(middleware.Admin(cfg))
	{
		v1.POST("/book", hnd.CreateBook)
	}
}

// BookFinderHTTPHandler is a handler for book APIs
func BookFinderHTTPHandler(cfg config.Config, router *gin.Engine, cf masterservicev1.BookFinderUseCase) {
	hnd := masterhandlerv1.NewBookFinderHandler(cf)
	v1 := router.Group("/v1")
	{
		log.Println("masuk finder")
		v1.GET("/book/list", hnd.GetBooks)
	}
}

// BookDeleterHTTPHandler is a handler for book APIs
func BookDeleterHTTPHandler(cfg config.Config, router *gin.Engine, ud masterservicev1.BookDeleterUseCase, cloudStorage interfaces.CloudStorageUseCase) {
	hnd := masterhandlerv1.NewBookDeleterHandler(ud, cloudStorage)
	v1 := router.Group("/v1")

	v1.Use(middleware.Auth(cfg))
	v1.Use(middleware.Admin(cfg))
	{
		v1.DELETE("/book/delete/:id", hnd.DeleteBook)
	}
}

// BookUpdaterHTTPHandler is a handler for book APIs
func BookUpdaterHTTPHandler(cfg config.Config, router *gin.Engine, uu masterservicev1.BookUpdaterUseCase, uf masterservicev1.BookFinderUseCase, cloudStorage interfaces.CloudStorageUseCase) {
	hnd := masterhandlerv1.NewBookUpdaterHandler(uu, uf)
	v1 := router.Group("/v1")

	v1.Use(middleware.Auth(cfg))
	v1.Use(middleware.Admin(cfg))
	{
		v1.PUT("/book/update/:id", hnd.UpdateBook)
	}
}

// LoanCreatorHTTPHandler is a handler for loan APIs
func LoanCreatorHTTPHandler(cfg config.Config, router *gin.Engine, lc loanservicev1.LoanCreatorUseCase, cloudStorage interfaces.CloudStorageUseCase) {
	hnd := loanhandlerv1.NewLoanRequestHandler(lc, cloudStorage)
	v1 := router.Group("/v1")

	{
		v1.POST("/loan/request", hnd.CreateLoanRequest)
	}
}

func LoanUpdaterHTTPHandler(cfg config.Config, router *gin.Engine, lu loanservicev1.LoanUpdaterUseCase, lf loanservicev1.LoanFinderUseCase, cloudStorage interfaces.CloudStorageUseCase) {
	hnd := loanhandlerv1.NewLoanUpdaterHandler(lu, lf)
	v1 := router.Group("/v1")

	v1.Use(middleware.Auth(cfg))
	v1.Use(middleware.Admin(cfg))
	{
		v1.PUT("/loan/update/:id", hnd.UpdateLoan)
	}
}

// BookFinderHTTPHandler is a handler for book APIs
func LoanFinderHTTPHandler(cfg config.Config, router *gin.Engine, lf loanservicev1.LoanFinderUseCase) {
	hnd := loanhandlerv1.NewLoanFinderHandler(lf)
	v1 := router.Group("/v1")
	{
		v1.GET("/loan/list", hnd.GetLoanRequests)
		v1.GET("/loan/details/:id", hnd.GetLoanRequestByID)
	}
}