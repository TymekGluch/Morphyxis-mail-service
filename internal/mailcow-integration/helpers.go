package mailcowIntegration

import (
	"bytes"
	"fmt"
	"mime/quotedprintable"
	"strings"
	"time"
)

func createSmtpMessage(from string, to string, subject string, body string) []byte {
	hostForMessageID := "localhost"
	parts := strings.Split(from, "@")
	if len(parts) == 2 && parts[1] != "" {
		hostForMessageID = parts[1]
	}

	messageID := fmt.Sprintf("<%d.%s>", time.Now().UnixNano(), hostForMessageID)
	date := time.Now().Format(time.RFC1123Z)

	var encodedBody bytes.Buffer
	writer := quotedprintable.NewWriter(&encodedBody)
	_, _ = writer.Write([]byte(body))
	_ = writer.Close()

	return []byte("From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Date: " + date + "\r\n" +
		"Message-ID: " + messageID + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"Content-Transfer-Encoding: quoted-printable\r\n" +
		"\r\n" +
		encodedBody.String(),
	)
}
