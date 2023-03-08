package configDomain

type Config struct {
	PORT       int
	JWT_SECRET string
	JWT_EXP    int
	DB_CONFIG  string

	SMTP_HOST  string
	SMTP_PORT  int
	SMTP_USER  string
	SMTP_PASS  string
	SMTP_FROM  string
}
