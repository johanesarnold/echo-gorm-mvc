package api

import (
	"bytes"
	"encoding/base64"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func UploadFileS3(session *session.Session, bucket string, filename string, base64Req string) error {
	path, err := Base64ToFile(base64Req)
	if err != nil {
		return err
	}

	upFile, err := os.Open(path)
	if err != nil {
		return err
	}

	upFileInfo, _ := upFile.Stat()
	var fileSize int64 = upFileInfo.Size()
	fileBuffer := make([]byte, fileSize)
	upFile.Read(fileBuffer)

	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(filename),
		ACL:         aws.String("public-read"),
		Body:        bytes.NewReader(fileBuffer),
		ContentType: aws.String(http.DetectContentType(fileBuffer)),
	})
	upFile.Close()

	err = os.Remove(path)
	if err != nil {
		return err
	}

	return err
}

func WriteLog(err error) {
	f, _ := os.OpenFile("log/echo.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	defer f.Close()
	log.SetOutput(f)
	log.Println(err)
}

func Base64ToFile(base64Req string) (string, error) {
	decodedBase64, err := base64.StdEncoding.DecodeString(base64Req)
	if err != nil {
		return "", err
	}

	now := time.Now()
	sec := now.Unix()
	filename := strconv.FormatInt(sec, 10) + ".jpg"
	path := "public/" + filename
	file, _ := os.Create(path)
	file.Write(decodedBase64)
	file.Close()

	return path, err
}
