package msgDomain

type IMsg struct {
	ERR_NOT_FOUND            string
	ERR_ID_ALREADY_EXISTS    string
	ERR_EMAIL_ALREADY_EXISTS string
	ERR_SAVING_IN_DATABASE   string
}

var Msg IMsg

func New() {
	Msg = IMsg{
		ERR_NOT_FOUND:            "ERR_NOT_FOUND",
		ERR_ID_ALREADY_EXISTS:    "ERR_ID_ALREADY_EXISTS",
		ERR_EMAIL_ALREADY_EXISTS: "ERR_EMAIL_ALREADY_EXISTS",
		ERR_SAVING_IN_DATABASE:   "ERR_SAVING_IN_DATABASE",
	}

}
