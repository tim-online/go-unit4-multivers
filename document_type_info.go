package multivers

type DocumentTypeInfo struct {
	DocumentType       int          `json:"documentType"`
	EmailSenderAddress string       `json:"emailSenderAddress"`
	MailTemplate       MailTemplate `json:"mailTemplate"`
}
