package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// validateCredentials verifica las credenciales de AWS usando STS GetCallerIdentity
func validateCredentials(ctx context.Context, cfg aws.Config) error {
	client := sts.NewFromConfig(cfg)
	_, err := client.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return fmt.Errorf("failed to validate AWS credentials: %v", err)
	}
	fmt.Println("AWS credentials validated successfully")
	return nil
}

func listS3Buckets(ctx context.Context, cfg aws.Config) error {
	client := s3.NewFromConfig(cfg)
	result, err := client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return fmt.Errorf("failed to list S3 buckets: %v", err)
	}
	fmt.Println("S3 Buckets:")
	for _, bucket := range result.Buckets {
		fmt.Printf("- %s (Created: %s)\n", *bucket.Name, bucket.CreationDate)
	}
	return nil
}

func listEC2Instances(ctx context.Context, cfg aws.Config) error {
	client := ec2.NewFromConfig(cfg)
	result, err := client.DescribeInstances(ctx, &ec2.DescribeInstancesInput{})
	if err != nil {
		return fmt.Errorf("failed to list EC2 instances: %v", err)
	}
	fmt.Println("EC2 Instances:")
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			instanceID := "<no-id>"
			if instance.InstanceId != nil {
				instanceID = *instance.InstanceId
			}
			state := "<no-state>"
			if instance.State != nil && instance.State.Name != "" {
				state = string(instance.State.Name)
			}
			fmt.Printf("- %s (State: %s)\n", instanceID, state)
		}
	}
	return nil
}

func listLambdaFunctions(ctx context.Context, cfg aws.Config) error {
	client := lambda.NewFromConfig(cfg)
	result, err := client.ListFunctions(ctx, &lambda.ListFunctionsInput{})
	if err != nil {
		return fmt.Errorf("failed to list Lambda functions: %v", err)
	}
	fmt.Println("Lambda Functions:")
	for _, function := range result.Functions {
		fmt.Printf("- %s (Runtime: %s)\n", *function.FunctionName, function.Runtime)
	}
	return nil
}

func main() {
	region := flag.String("region", "us-east-1", "AWS region")
	operation := flag.String("operation", "", "Cloud operation (list-s3, list-ec2, list-lambda)")
	flag.Parse()

	if *operation == "" {
		fmt.Fprintln(os.Stderr, "Error: -operation flag is required (list-s3, list-ec2, list-lambda)")
		os.Exit(1)
	}

	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(*region))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading AWS config: %v\n", err)
		os.Exit(1)
	}

	// Validar credenciales antes de ejecutar cualquier operación
	if err := validateCredentials(ctx, cfg); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	switch *operation {
	case "list-s3":
		if err := listS3Buckets(ctx, cfg); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	case "list-ec2":
		if err := listEC2Instances(ctx, cfg); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	case "list-lambda":
		if err := listLambdaFunctions(ctx, cfg); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "Error: Invalid operation. Use list-s3, list-ec2, or list-lambda\n")
		os.Exit(1)
	}
}