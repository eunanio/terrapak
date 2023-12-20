package s3

import (
	"bytes"
	"context"
	"fmt"
	"terrapak/internal/config"
	"terrapak/internal/config/mid"
	"terrapak/internal/storage/providers/shared"
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

func (s *S3Config) NewProvider() *S3Config {
	return &S3Config{}
}

func (s *S3Config) Type() string {
	return "s3"
}

func loadClient() *S3Config {
	config := S3Config{}
	cfg, err := aws_config.LoadDefaultConfig(context.TODO()); if err != nil {
		panic("failed to load config")
	}
	config.Client = s3.NewFromConfig(cfg)
	return &config
}

func Upload(mid mid.MID, data []byte) {
	s3c := loadClient()
	gc := config.GetDefault()
	S3Client := s3c.Client
	prefix, filename := shared.BuldPathValues(mid)
	ctx := context.TODO()
	
	key := fmt.Sprintf("%s/%s",prefix,filename)
	if S3Client != nil  {
		_, err := S3Client.PutObject(ctx,&s3.PutObjectInput{
			Bucket: &gc.BucketName,
			Key:    &key,
			Body:   bytes.NewReader(data),
		})
		if err != nil {
			panic(fmt.Sprintf("failed to upload file to s3, %v", err))
		}
	}

	fmt.Println("Successfully uploaded file to s3?")
}

func (c *S3Config) Download(mid mid.MID) (url string, err error) {
	prefix, filename := shared.BuldPathValues(mid)
	gc := config.GetDefault()
	ctx := context.TODO()
	key := fmt.Sprintf("%s/%s",prefix,filename)

	return getPresignUrl(gc.BucketName,key,ctx)
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
	fmt.Println(request)

	if err != nil {
		return "", fmt.Errorf("failed to presign request, %v", err)
	}
	fmt.Println(request.URL)
	return request.URL, nil
}

func getPresignUrl(bucket,key string, ctx context.Context) (url string, err error) {
	client := loadClient()
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

func (s *S3Config)DeleteObject(mid mid.MID){
	client := loadClient()
	gc := config.GetDefault()
	S3Client := client.Client
	prefix, filename := shared.BuldPathValues(mid)
	ctx := context.TODO()
	key := fmt.Sprintf("%s/%s",prefix,filename)

	_, err := S3Client.DeleteObject(ctx,&s3.DeleteObjectInput{
		Bucket: &gc.BucketName,
		Key:    &key,
	})

	if err != nil {
		panic(fmt.Sprintf("failed to delete file from s3, %v", err))
	}
}

