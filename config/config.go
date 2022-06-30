package config

import "os"

var (
	Dsn                = os.Getenv("DB_DSN")
	AWSAccessKeyID     = os.Getenv("AWS_ACCESS_KEY_ID")
	AWSSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	AWSRegion          = os.Getenv("AWS_REGION")
	S3Bucket           = os.Getenv("S3_BUCKET")
)
