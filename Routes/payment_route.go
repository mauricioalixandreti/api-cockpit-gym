package routes

import (
	controller "API-COCKPIT-GYM/Controller/payment"

	"github.com/gin-gonic/gin"
)

func PaymentsRoutes(api *gin.RouterGroup) {
	payments := api.Group("/payments")
	payments.GET("/gyms/:gym_id/payments_by_month/:month", controller.GetPaymentsByMonth)
	payments.GET("/payments_by_months/:month/gym/:gym_id/status/:status", controller.GetPaymentsByMonthAndStatus)
	payments.POST("/payments")
	payments.PUT("/payments/payment_id")
	payments.DELETE("/payments/payment_id")
}
