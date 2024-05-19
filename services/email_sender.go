package services

import (
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/sirupsen/logrus"
)

var yourDomain string = "sandboxfbf511bc28be47d3a0485b8f895fece9.mailgun.org"

var privateAPIKey string = "b85f0e157ce6d87d4bbab28b3d0f6537-32a0fef1-c4adcfcf"

func Sender() {
	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)

	//	s.GetSubscriptions()
	rate, _ := GetExchangeRate()

	sender := "ksenia.agag@gmail.com"
	subject := "Actual UAH to USD exchange rate"
	body := fmt.Sprintf("%f", rate)
	recipient := "ksenia.agag@gmail.com"

	message := mg.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resp, id, _ := mg.Send(ctx, message)

	logrus.Errorf("ID: %s Resp: %s\n", id, resp)
}

//func (s *SubscriptionService) GetEmail() {
//	s.repo.GetByEmail()
//}
