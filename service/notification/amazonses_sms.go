package notification

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

// Type of SMS delivery mode.
type Type string

const (
	// Promotional are non-critical messages, such as marketing messages.
	// Amazon SNS optimizes the message delivery to incur the lowest cost.
	Promotional Type = "Promotional"

	// Transactional messages are critical messages that support
	// customer transactions, such as one-time passcodes for multi-factor authentication.
	// Amazon SNS optimizes the message delivery to achieve the highest reliability.
	Transactional = "Transactional"
)

// Send SMS to a specific number with a specific SMS
// Need to set AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY to send the message using your AWS account
func (n *VerifyNotification) Send() error {
	sess := session.Must(session.NewSession())
	svc := sns.New(sess)
	attrs := map[string]*sns.MessageAttributeValue{}
	attrs["AWS.SNS.SMS.SenderID"] = &sns.MessageAttributeValue{
		DataType:    aws.String("String"),
		StringValue: aws.String(n.From),
	}

	kind := Promotional
	attrs["AWS.SNS.SMS.SMSType"] = &sns.MessageAttributeValue{
		DataType:    aws.String("String"),
		StringValue: aws.String(string(kind)),
	}

	params := &sns.PublishInput{
		Message:           aws.String(n.Message),
		PhoneNumber:       aws.String(n.PhoneNumber),
		MessageAttributes: attrs,
	}

	_, err := svc.Publish(params)
	if err != nil {
		return err
	}

	return nil
}
