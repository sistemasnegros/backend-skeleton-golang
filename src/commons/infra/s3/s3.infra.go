package s3Infra

import (
	"context"

	configService "backend-skeleton-golang/commons/app/services/config-service"
	logService "backend-skeleton-golang/commons/app/services/log-service"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func New() *s3.Client {

	s3Config := configService.GetS3()

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			PartitionID:   "aws",
			URL:           s3Config.AWS_ENDPOINT,
			SigningRegion: s3Config.AWS_REGION,
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(s3Config.AWS_ACCESS_KEY_ID, s3Config.AWS_SECRET_ACCESS_KEY, "")),
	)

	cfg.EndpointResolverWithOptions = customResolver

	if err != nil {
		logService.Error(err.Error())
		panic(err)
	}

	logService.Info("successfully connection S3!")
	client := s3.NewFromConfig(cfg)


	return client
}
