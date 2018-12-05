package notification

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"log"
)

type AWSSESNotifier struct {
}

// Send SMS to a specific number with a specific SMS
// Need to set AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY to send the message using your AWS account
func (d AWSSESNotifier) Send(n VerificationNotification) error {
	sess := session.Must(session.NewSession())
	svc := sns.New(sess)
	attrs := map[string]*sns.MessageAttributeValue{}
	attrs["AWS.SNS.SMS.SenderID"] = &sns.MessageAttributeValue{
		DataType:    aws.String("String"),
		StringValue: aws.String(n.From),
	}

	attrs["AWS.SNS.SMS.SMSType"] = &sns.MessageAttributeValue{
		DataType:    aws.String("String"),
		StringValue: aws.String("Transactional"),
	}

	p := "+" + n.CountryCode + n.PhoneNumber

	params := &sns.PublishInput{
		Message:           aws.String(n.Message),
		PhoneNumber:       aws.String(p),
		MessageAttributes: attrs,
	}

	_, err := svc.Publish(params)
	if err != nil {
		return err
	}

	log.Printf("SMS was sent message with code: %s using AWS SES\n", n)

	return nil
}
