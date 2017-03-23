package multivers

type MailTemplate struct {
	Messages              []Message `json:"messages"`
	Bcc                   string    `json:"bcc"`
	canChange             bool      `json:"canChange"`
	Cc                    string    `json:"cc"`
	DocumentType          int       `json:"documentType"`
	// TODO find out what MailLanguageTemplates is
	// MailLanguageTemplates []string  `json:"mailLanguageTemplates"`
	MailTemplateID	      string    `json:"mailTemplateId"`
	Name                  string    `json:"name"`
	To                    string    `json:"to"`
}
