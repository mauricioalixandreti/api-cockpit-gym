package payment

import (
	payments "API-COCKPIT-GYM/Service/payment"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPaymentsByMonth(c *gin.Context) {

	// Obtém os parâmetros da URL
	gymID := c.Param("gym_id")
	month := c.Param("month")

	// Chama o a validação do Service
	payments, err := payments.GetPaymentsByMonth(gymID, month)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	// Retorna a lista de pagamentos
	c.JSON(http.StatusOK, gin.H{

		"success": true,
		"message": "Pagamentos do mês retornados com sucesso",
		"count":   len(payments),
		"data":    payments,
	})
}

func GetPaymentsByMonthAndStatus(c *gin.Context) {

	gymID := c.Param("gym_id")
	month := c.Param("month")
	status := c.Param("status")

	// Chama o service.
	payments, err := payments.GetPaymentsByMonthAndStatus(gymID, month, status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	// Retorna o Mês e o Status do pagamento.
	c.JSON(http.StatusOK, gin.H{

		"success": true,
		"message": "Pagamentos do mês e status retornados com sucesso",
		"count":   len(payments),
		"status":  status,
		"month":   month,
		"gym_id":  gymID,
		"data":    payments,
	})

}
