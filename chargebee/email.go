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

func sendEmail(clientEmail string, referralID string) error {
	body := brevo.SendSmtpEmail{
		Subject:    cfg.ReferralEmailSubject,
		TemplateId: int64(240),
		Sender: &brevo.SendSmtpEmailSender{
			Name:  cfg.ReferralEmailName,
			Email: cfg.ReferralEmailAddress,
		},
		To: []brevo.SendSmtpEmailTo{
			{
				Name:  "",
				Email: clientEmail,
			},
		},
		ReplyTo: &brevo.SendSmtpEmailReplyTo{
			Name:  cfg.ReferralEmailName,
			Email: cfg.ReferralEmailAddress,
		},
		Params: map[string]interface{}{
			"REFERRALID": referralID,
			"subject":    cfg.ReferralEmailSubject,
		},
	}
	_, _, err := brevoApiClient.TransactionalEmailsApi.SendTransacEmail(context.Background(), body)
	if err != nil {
		return tracerr.Wrap(err)
	}
	return nil
}
