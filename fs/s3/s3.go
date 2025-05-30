// Package s3 provides an AWS S3 access layer
package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	s3 "github.com/clicknclear/afero-s3"
	"github.com/spf13/afero"

	"github.com/clicknclear/ftpserver/config/confpar"
)

// LoadFs loads a file system from an access description
func LoadFs(access *confpar.Access) (afero.Fs, error) {
	endpoint := access.Params["endpoint"]
	region := access.Params["region"]
	bucket := access.Params["bucket"]
	keyID := access.Params["access_key_id"]
	secretAccessKey := access.Params["secret_access_key"]
	basePath := access.Params["base_path"]

	conf := aws.Config{
		Region:           aws.String(region),
		DisableSSL:       aws.Bool(access.Params["disable_ssl"] == "true"),
		S3ForcePathStyle: aws.Bool(access.Params["path_style"] == "true"),
	}

	if keyID != "" && secretAccessKey != "" {
		conf.Credentials = credentials.NewStaticCredentials(keyID, secretAccessKey, "")
	}

	if endpoint != "" {
		conf.Endpoint = aws.String(endpoint)
	}

	sess, errSession := session.NewSession(&conf)

	if errSession != nil {
		return nil, errSession
	}

	s3Fs := s3.NewFs(bucket, sess, basePath)

	return s3Fs, nil
}
