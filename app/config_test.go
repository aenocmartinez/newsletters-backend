package app

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestReadConfigApp(t *testing.T) {
	assert := assert.New(t)
	envFilePath := filepath.Join(os.Getenv("HOME"), "/gallery-env/.env")
	err := godotenv.Load(envFilePath)
	assert.Equal(err, nil)

	assert.NotEmpty(AppName())
	assert.NotEmpty(AppVersion())
	assert.NotEmpty(AppPath())
	assert.NotEmpty(NumItemsPerPage())
	assert.NotEmpty(NumItemsRecirculation())
	assert.NotEmpty(AwsUrlS3())
	assert.NotEmpty(AwsBucketS3())
	assert.NotEmpty(Sentry())
}
