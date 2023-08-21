package app

import (
	"os"
	"pulzo/src/shared/infraestructure/util"
)

func AppName() string {
	return os.Getenv("APP_NAME")
}

func AppVersion() string {
	return os.Getenv("APP_VERSION")
}

func AppPath() string {
	return AppName() + "/" + AppVersion()
}

func NumItemsPerPage() int {
	numItemsPerPage := os.Getenv("PAGINATOR_NUM_ITEMS")
	return util.StringToInt(numItemsPerPage)
}

func NumItemsRecirculation() int {
	numItemsRecirculation := os.Getenv("RECIRCULATION_NUM_ITEMS")
	return util.StringToInt(numItemsRecirculation)
}

func AwsUrlS3() string {
	return os.Getenv("AWS_URL_UPLOAD")
}

func AwsBucketS3() string {
	return os.Getenv("AWS_BUCKET")
}

func Sentry() string {
	return os.Getenv("SENTRY_DNS")
}
