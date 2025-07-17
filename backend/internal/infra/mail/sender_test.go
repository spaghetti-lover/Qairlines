package mail

import (
	"testing"

	"github.com/spaghetti-lover/qairlines/config"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping email test in short mode")
	}
	config, err := config.LoadConfig("../../..")
	require.NoError(t, err)

	sender := NewGmailSender(config.MailSenderName, config.MailSenderAddress, config.MailSenderPassword)

	subject := "Test Email"
	content := "<h1>This is a test email</h1>"
	to := []string{"22028063@vnu.edu.vn"}
	attachFiles := []string{"../../../Makefile"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err, "Failed to send email")
	t.Log("Email sent successfully")
}
