package account

import amqp "github.com/rabbitmq/amqp091-go"

type IAccount interface {
	AbleToReport() (bool, *string)
}

type Account struct {
	UserId string
}

func InitAccount(UserId string, mq *amqp.Connection, ch *amqp.Channel) (*Account, error) {
	// if can't establish database connection = return nil, err
	return &Account{
		UserId: UserId,
	}, nil
}

// Check if user can make report
func (a *Account) AbleToReport() (bool, *string) {

	msg := "Monthly reports limit reached"
	return false, &msg
}
