package routes

import (
	"net/http"

	authhandler "backend/internal/modules/auth/handler"
	authroutes "backend/internal/modules/auth/routes"
	authsvc "backend/internal/modules/auth/service"

	"backend/internal/modules/user/handler"
	userrepo "backend/internal/modules/user/repository"
	userroutes "backend/internal/modules/user/routes"
	usersvc "backend/internal/modules/user/service"

	violationhandler "backend/internal/modules/violation/handler"
	violationrepo "backend/internal/modules/violation/repository"
	violationroutes "backend/internal/modules/violation/routes"
	violationsvc "backend/internal/modules/violation/service"

	invoicehandler "backend/internal/modules/invoice/handler"
	invoicerepo "backend/internal/modules/invoice/repository"
	invoiceroutes "backend/internal/modules/invoice/routes"
	invoicesvc "backend/internal/modules/invoice/service"

	paymentshandler "backend/internal/modules/payment-transaction/handler"
	paymentsrepo "backend/internal/modules/payment-transaction/repository"
	paymentsroutes "backend/internal/modules/payment-transaction/routes"
	paymentssvc "backend/internal/modules/payment-transaction/service"

	uploadhandler "backend/internal/modules/upload/handler"
	uploaddroutes "backend/internal/modules/upload/routes"
	uploadsvc "backend/internal/modules/upload/service"

	"backend/pkg/response"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine, db *gorm.DB, jwtSecret string) {
	v1 := router.Group("/api")

	userRepository := userrepo.NewRepository(db)
	userService := usersvc.NewService(userRepository)
	userHandler := handler.NewHandler(userService)
	userroutes.Register(v1, userHandler, jwtSecret)

	violationRepository := violationrepo.NewRepository(db)
	violationService := violationsvc.NewService(violationRepository)
	violationHandler := violationhandler.NewHandler(violationService)
	violationroutes.Register(v1, violationHandler, jwtSecret)

	violationTypeRepository := violationrepo.NewViolationTypeRepository(db)
	violationTypeService := violationsvc.NewViolationTypeService(violationTypeRepository)
	violationTypeHandler := violationhandler.NewViolationTypeHandler(violationTypeService)
	violationroutes.RegisterViolationType(v1, violationTypeHandler, jwtSecret)

	fineRuleVersionRepository := violationrepo.NewFineRuleVersionRepository(db)
	fineRuleVersionService := violationsvc.NewFineRuleVersionService(fineRuleVersionRepository)
	fineRuleVersionHandler := violationhandler.NewFineRuleVersionHandler(fineRuleVersionService)
	violationroutes.RegisterFineRuleVersion(v1, fineRuleVersionHandler, jwtSecret)

	fineRuleDetailRepository := violationrepo.NewFineRuleDetailRepository(db)
	fineRuleDetailService := violationsvc.NewFineRuleDetailService(fineRuleDetailRepository)
	fineRuleDetailHandler := violationhandler.NewFineRuleDetailHandler(fineRuleDetailService)
	violationroutes.RegisterFineRuleDetail(v1, fineRuleDetailHandler, jwtSecret)

	invoiceRepository := invoicerepo.NewRepository(db)
	fineCalculator := violationsvc.NewFineCalculationService()
	invoiceService := invoicesvc.NewService(invoiceRepository, violationRepository, fineCalculator)
	invoiceHandler := invoicehandler.NewHandler(invoiceService)
	invoiceroutes.Register(v1, invoiceHandler, jwtSecret)

	paymentsRepo := paymentsrepo.NewRepository(db)
	paymentsService := paymentssvc.NewService(paymentsRepo)
	paymentsHandler := paymentshandler.NewHandler(paymentsService)
	paymentsroutes.Register(v1, paymentsHandler, jwtSecret)

	authService := authsvc.NewService(userRepository, jwtSecret)
	authHandler := authhandler.NewHandler(authService)
	authroutes.Register(v1, authHandler)

	uploadService := uploadsvc.NewService()
	uploadHandler := uploadhandler.NewHandler(uploadService)
	uploaddroutes.Register(v1, uploadHandler)

	router.Static("/uploads", "./uploads")

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.Success("server is running", nil))
	})
}

