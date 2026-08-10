package payments

import (
	Models "API-COCKPIT-GYM/Models"
	repository "API-COCKPIT-GYM/Repository"
)

func GetPaymentsByMonth(gymID string, month string) ([]Models.Payment, error) {

	// Validar busca de dados no Firebase.

	return repository.GetPaymentsByMonth(gymID, month)

}

func GetPaymentsByMonthAndStatus(gymID string, month string, status string) ([]Models.Payment, error) {

	// Validar Mês e Status do pagamento.

	return repository.GetPaymentsByMonthAndStatus(gymID, month, status)

}

func CreatePayment(payment *Models.Payment) error {

	// Validar os dados no Firebase.

	return repository.CreatePayment(payment)

}

func UpdatePayment(paymentID string, payment *Models.Payment) error {

	// Validar atualizações no Firebase.

	return repository.UpdatePayment(paymentID, payment)

}

func DeletePayment(paymentID string, gymID string) error {

	// Validar a deleção dos dados no Firebase.

	return repository.DeletePayment(paymentID, gymID)

}
