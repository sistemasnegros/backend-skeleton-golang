package smtpDomain

type SendArgs struct {
	To       string
	Cc       string
	Subject  string
	Data     interface{}
	Template string
}

type EmailTemplateDefault struct {
	FullName      string
	Message       string
	URL           string
	ButtonMessage string
}
