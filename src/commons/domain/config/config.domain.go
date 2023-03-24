package configDomain

type Config struct {
	PORT       int
	JWT_SECRET string
	JWT_EXP    int
	DB_CONFIG  string

	SMTP_HOST string
	SMTP_PORT int
	SMTP_USER string
	SMTP_PASS string
	SMTP_FROM string

	MONGO_DB string

	AWS_SECRET_ACCESS_KEY string
	AWS_ACCESS_KEY_ID     string
	AWS_BUCKET            string
	AWS_REGION            string
	AWS_ENDPOINT          string
	AWS_URL_PUBLIC        string
}
