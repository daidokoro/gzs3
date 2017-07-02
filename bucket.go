package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// -- Contains all things S3

// S3Read - Reads the content of a given s3 url endpoint and returns the content string.
func S3Read(url string, sess *session.Session) (string, error) {
	svc := s3.New(sess)

	// Parse s3 url
	bucket := strings.Split(strings.Replace(strings.ToLower(url), `s3://`, "", -1), `/`)[0]
	key := strings.Replace(strings.ToLower(url), fmt.Sprintf("s3://%s/", bucket), "", -1)

	params := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	log.Debug(fmt.Sprintln("Calling S3 [GetObject] with parameters:", params))
	resp, err := svc.GetObject(params)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	log.Debug("Reading from S3 Response Body")
	buf.ReadFrom(resp.Body)
	return buf.String(), nil

}

// S3write - Writes a file to s3 and returns the presigned url
func S3write(bucket string, key string, body *bytes.Buffer, sess *session.Session) (string, error) {
	svc := s3.New(sess)
	params := &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   bytes.NewReader(body.Bytes()),
		Metadata: map[string]*string{
			"created_by": aws.String("gzs3"),
		},
	}

	log.Debug(fmt.Sprintln("Calling S3 [PutObject] with parameters:", params))
	_, err := svc.PutObject(params)
	if err != nil {
		return "", err
	}

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})

	url, err := req.Presign(10 * time.Minute)
	if err != nil {
		return "", err
	}

	return url, nil

}
