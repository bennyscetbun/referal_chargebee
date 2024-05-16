package chargebee

import (
	"context"

	brevo "github.com/getbrevo/brevo-go/lib"
	"github.com/ztrue/tracerr"
)

var brevoApiClient *brevo.APIClient

func setupBrevoClient(apiKey string) error {
	cfg := brevo.NewConfiguration()
	//Configure API key authorization: api-key
	cfg.AddDefaultHeader("api-key", apiKey)

	brevoApiClient = brevo.NewAPIClient(cfg)
	return nil
}

func sendEmail(clientEmail string, emailConf ConfigEmail, params map[string]interface{}) error {
	body := brevo.SendSmtpEmail{
		Subject:    emailConf.Subject,
		TemplateId: int64(emailConf.TemplateID),
		Sender: &brevo.SendSmtpEmailSender{
			Name:  emailConf.Name,
			Email: emailConf.Address,
		},
		To: []brevo.SendSmtpEmailTo{
			{
				Name:  "",
				Email: clientEmail,
			},
		},
		ReplyTo: &brevo.SendSmtpEmailReplyTo{
			Name:  emailConf.Name,
			Email: emailConf.Address,
		},
		Params: params,
	}
	_, _, err := brevoApiClient.TransactionalEmailsApi.SendTransacEmail(context.Background(), body)
	if err != nil {
		return tracerr.Wrap(err)
	}
	return nil
}

func sendReferalEmail(clientEmail string, referralID string) error {
	return sendEmail(clientEmail, cfg.ReferralEmail, map[string]interface{}{"REFERRALID": referralID})
}

func sendCreditAddedEmail(clientEmail string) error {
	return sendEmail(clientEmail, cfg.CreditAddedEmail, map[string]interface{}{})
}
