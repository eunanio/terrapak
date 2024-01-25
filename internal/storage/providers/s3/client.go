package s3

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"terrapak/internal/config"
	"terrapak/internal/config/mid"
	"time"

	aws_config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Config struct {
	BucketName         string
	BucketRegion       string
	BucketPrefix       string
	AccessKeyID        string
	SecretAccessKey    string
	Client             *s3.Client
	DefaultCredentials bool
}

type S3Presign struct {
	PresignClient *s3.PresignClient
	LinkExpire         int
}

type S3Provider struct {
	Client *S3Config
}

func NewProvider() *S3Provider {
	return &S3Provider{
		Client: loadClient(),
	}
}

func (p *S3Provider) Type() string {
	return "s3"
}

func loadClient() *S3Config {
	config := S3Config{}
	cfg, err := aws_config.LoadDefaultConfig(context.TODO()); if err != nil {
		log.Fatal("failed to load config")
	}
	config.Client = s3.NewFromConfig(cfg)
	return &config
}

func (p *S3Provider) Upload(mid mid.MID, data []byte) error {
	s3c := p.Client
	gc := config.GetDefault()
	S3Client := s3c.Client
	ctx := context.TODO()
	
	key := mid.Filepath()
	if S3Client != nil  {
		_, err := S3Client.PutObject(ctx,&s3.PutObjectInput{
			Bucket: &gc.StorageSource.Path,
			Key:    &key,
			Body:   bytes.NewReader(data),
		})
		if err != nil {
			return fmt.Errorf("failed to upload file to s3, %v", err)
		}
	}
	return nil
}

func (p *S3Provider) Download(mid mid.MID) (url string, err error) {
	gc := config.GetDefault()
	ctx := context.TODO()
	key := mid.Filepath()

	return getPresignUrl(gc.StorageSource.Path,key,ctx,p.Client)
}

func (p *S3Presign) setPresignOptions(options *s3.PresignOptions) {
	p.LinkExpire = 30
	options.Expires = time.Duration(p.LinkExpire) * time.Second
}

func (p *S3Presign) Presign(bucket,key string, ctx context.Context) (url string, err error) {
	request, err := p.PresignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})

	if err != nil {
		return "", fmt.Errorf("failed to presign request, %v", err)
	}
	fmt.Println(request.URL)
	return request.URL, nil
}

func getPresignUrl(bucket,key string, ctx context.Context, client *S3Config) (url string, err error) {

	S3Client := client.Client

	p := S3Presign{
		PresignClient: s3.NewPresignClient(S3Client),
	}

	request, err := p.PresignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	},p.setPresignOptions)

	fmt.Println(request)
	if err != nil {
		return "", fmt.Errorf("failed to presign request, %v", err)
	}

	url = request.URL
	
	return url, nil
}

func (p *S3Provider) Delete(mid mid.MID) error {
	client := p.Client
	gc := config.GetDefault()
	S3Client := client.Client
	ctx := context.TODO()
	key := mid.Filepath()


	_, err := S3Client.DeleteObject(ctx,&s3.DeleteObjectInput{
		Bucket: &gc.StorageSource.Path,
		Key:    &key,
	})

	if err != nil {
		return fmt.Errorf("failed to delete file from s3, %v", err)
	}

	return nil
}

