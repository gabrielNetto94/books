package s3

import (
	"books/pkg/storage"
	"context"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3Impl struct {
	bucket        string
	client        *s3.Client
	presignClient *s3.PresignClient
}

type S3Config struct {
	Bucket            string
	Region            string
	DisableAutoConfig bool
	AWSConfig         aws.Config
	Endpoint          string
	UsePathStyle      bool
}

func New(ctx context.Context, cfg S3Config) (storage.Storage, error) {
	var awsCfg aws.Config
	var err error
	if cfg.DisableAutoConfig {
		awsCfg = cfg.AWSConfig
	} else {
		awsCfg, err = config.LoadDefaultConfig(ctx, config.WithRegion(cfg.Region))
		if err != nil {
			return nil, fmt.Errorf("storage: loading aws config: %w", err)
		}
	}
	client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = cfg.UsePathStyle
		if cfg.Endpoint != "" {
			o.BaseEndpoint = aws.String(cfg.Endpoint)
		}
	})
	presigner := s3.NewPresignClient(client)

	return &s3Impl{bucket: cfg.Bucket, client: client, presignClient: presigner}, nil
}

func (s s3Impl) Put(ctx context.Context, key string, tags map[string]string, r io.Reader, size int64, contentType string) error {
	input := &s3.PutObjectInput{
		Bucket:        aws.String(s.bucket),
		Key:           aws.String(key),
		Body:          r,
		ContentType:   aws.String(contentType),
		ContentLength: &size, // The interface provides size, so we use it.
	}

	// Convert tags map[string]string to S3's URL-encoded string format
	if len(tags) > 0 {
		var tagParts []string
		for k, v := range tags {
			tagParts = append(tagParts, url.QueryEscape(k)+"="+url.QueryEscape(v))
		}
		input.Tagging = aws.String(strings.Join(tagParts, "&"))
	}

	_, err := s.client.PutObject(ctx, input)
	return err
}

func (s s3Impl) Get(ctx context.Context, key string) (io.ReadCloser, error) {

	input := &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	}

	output, err := s.client.GetObject(ctx, input)
	if err != nil {
		return nil, err
	}

	// The output body is already an io.ReadCloser
	return output.Body, nil
}

func (s s3Impl) Delete(ctx context.Context, key string) error {
	return nil
}

func (s s3Impl) Exists(ctx context.Context, key string) (bool, error) {
	return false, nil
}

func (s s3Impl) URL(ctx context.Context, key string, expiry time.Duration) (string, error) {
	return "", nil
}
