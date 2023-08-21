package pulzoservices

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"pulzo/src/newsletter/domain"
	"regexp"
	"strings"

	"github.com/getsentry/sentry-go"
)

var projectDirName string = "redaccion-back"

func createFileLocal(newsletterName, extension string, content []byte) (filePath string, err error) {
	filePath = "./" + pathFile(newsletterName+"."+extension)

	f, err := os.Create(filePath)
	if err != nil {
		sentry.CaptureException(err)
		return "", err
	}

	defer f.Close()

	_, err = f.WriteString(string(content))
	if err != nil {
		sentry.CaptureException(err)
		return "", err
	}

	return filePath, err
}

func deleteFileLocal(filePath string) error {
	return os.Remove(filePath)
}

func UploadFile(newsletter domain.Newsletter, extension string, content []byte) (path string, code int, err error) {

	var bucketS3 string = os.Getenv("APP_BUCKET_S3") + "/" + newsletter.Section()
	var endpointPulzoAwsS3 string = os.Getenv("APP_ENDPOINT_S3")

	filePath, err := createFileLocal(newsletter.JsonFileName(), extension, content)
	if err != nil {
		return "", 500, err
	}

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	fileLocal, err := os.Open(filePath)
	if err != nil {
		sentry.CaptureException(err)
		return "", 500, err
	}

	defer fileLocal.Close()

	part1, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		sentry.CaptureException(err)
		return "", 500, err
	}

	_, err = io.Copy(part1, fileLocal)
	if err != nil {
		sentry.CaptureException(err)
		return "", 500, err
	}

	writer.WriteField("path", bucketS3)
	writer.Close()

	client := &http.Client{}
	req, err := http.NewRequest("POST", endpointPulzoAwsS3, payload)
	if err != nil {
		sentry.CaptureException(err)
		return "", 500, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := client.Do(req)
	if err != nil {
		sentry.CaptureException(err)
		return "", 500, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		sentry.CaptureException(err)
		return "", 500, err
	}

	deleteFileLocal(filePath)
	return depureBodyResponse(string(body)), 200, nil
}

func depureBodyResponse(body string) string {
	body = strings.ReplaceAll(body, "urlFileUploaded", "")
	body = strings.ReplaceAll(body, "\"", "")
	body = strings.ReplaceAll(body, ":http", "http")
	body = strings.ReplaceAll(body, "{", "")
	body = strings.ReplaceAll(body, "}", "")
	return body
}

func pathFile(filaname string) string {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	return string(rootPath) + filaname
}
