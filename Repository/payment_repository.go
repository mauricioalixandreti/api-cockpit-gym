package repository

import (
	Models "API-COCKPIT-GYM/Models"
)

const paymentPath = "/payments"

func GetPaymentsByMonth(gymID string, month string) ([]Models.Payment, error) {

	var payments []Models.Payment

	// Buscar daddos no Firebase.

	return payments, nil

}

func GetPaymentsByMonthAndStatus(gymID string, month string, status string) ([]Models.Payment, error) {

	var payments []Models.Payment

	// Buscar Mês e Status do pagamento.

	return payments, nil

}

func CreatePayment(payment *Models.Payment) error {

	//Salvar dados no Firebase.

	return nil

}

func UpdatePayment(paymentID string, payment *Models.Payment) error {

	// Atualizar dados no Firebase.

	return nil

}

func DeletePayment(paymentID string, gymID string) error {

	// Excluir dados no Firebase.

	return nil

}
