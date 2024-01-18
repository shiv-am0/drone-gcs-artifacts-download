// package main

// import (
// 	"fmt"
// 	"os"
// 	"os/exec"
// )

// func main() {
// 	// Check if required environment variables are set
// 	awsAccessKeyID := os.Getenv("PLUGIN_AWS_ACCESS_KEY_ID")
// 	awsSecretAccessKey := os.Getenv("PLUGIN_AWS_SECRET_ACCESS_KEY")
// 	awsDefaultRegion := os.Getenv("PLUGIN_AWS_DEFAULT_REGION")
// 	awsBucketName := os.Getenv("PLUGIN_AWS_BUCKET_NAME")
// 	fetchDir := os.Getenv("PLUGIN_FETCH_DIR")
// 	downloadTarget := os.Getenv("PLUGIN_DOWNLOAD_TARGET")

// 	if awsAccessKeyID == "" || awsSecretAccessKey == "" || awsDefaultRegion == "" || awsBucketName == "" || fetchDir == "" || downloadTarget == "" {
// 		fmt.Println("Error: Please provide all required environment variables.")
// 		os.Exit(1)
// 	}

// 	// Set AWS credentials and region
// 	os.Setenv("AWS_ACCESS_KEY_ID", awsAccessKeyID)
// 	os.Setenv("AWS_SECRET_ACCESS_KEY", awsSecretAccessKey)
// 	os.Setenv("AWS_DEFAULT_REGION", awsDefaultRegion)

// 	// Download artifacts from S3 bucket to the specified target
// 	cmd := exec.Command("aws", "s3", "cp", fmt.Sprintf("s3://%s/%s", awsBucketName, fetchDir), downloadTarget, "--recursive")
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr

// 	err := cmd.Run()
// 	if err != nil {
// 		fmt.Printf("Error executing command: %v\n", err)
// 		os.Exit(1)
// 	}
// }

package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	// Check if required environment variables are set
	awsAccessKeyID := os.Getenv("PLUGIN_AWS_ACCESS_KEY_ID")
	awsSecretAccessKey := os.Getenv("PLUGIN_AWS_SECRET_ACCESS_KEY")
	awsDefaultRegion := os.Getenv("PLUGIN_AWS_DEFAULT_REGION")
	awsBucketName := os.Getenv("PLUGIN_AWS_BUCKET_NAME")
	fetchDir := os.Getenv("PLUGIN_FETCH_DIR")
	downloadTarget := os.Getenv("PLUGIN_DOWNLOAD_TARGET")

	if awsAccessKeyID == "" || awsSecretAccessKey == "" || awsDefaultRegion == "" || awsBucketName == "" || fetchDir == "" || downloadTarget == "" {
		fmt.Println("Error: Please provide all required environment variables.")
		os.Exit(1)
	}

	// Create a new AWS session
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsDefaultRegion),
		Credentials: credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, ""),
	})
	if err != nil {
		fmt.Printf("Error creating AWS session: %v\n", err)
		os.Exit(1)
	}

	// Create an S3 client
	s3Client := s3.New(sess)

	// List objects in the S3 bucket
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(awsBucketName),
		Prefix: aws.String(fetchDir),
	}

	result, err := s3Client.ListObjectsV2(input)
	if err != nil {
		fmt.Printf("Error listing objects: %v\n", err)
		os.Exit(1)
	}

	// Download each object
	for _, obj := range result.Contents {
		key := obj.Key
		output, err := s3Client.GetObject(&s3.GetObjectInput{
			Bucket: aws.String(awsBucketName),
			Key:    key,
		})
		if err != nil {
			fmt.Printf("Error downloading object %s: %v\n", *key, err)
			os.Exit(1)
		}

		// Save the object to the local file
		filePath := downloadTarget + "/" + *key
		err = saveToFile(output, filePath)
		if err != nil {
			fmt.Printf("Error saving object to file: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Downloaded object %s to %s\n", *key, filePath)
	}

	fmt.Println("Artifacts downloaded successfully.")
}

func saveToFile(output *s3.GetObjectOutput, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.ReadFrom(output.Body)
	if err != nil {
		return err
	}

	return nil
}
