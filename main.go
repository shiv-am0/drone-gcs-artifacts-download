package main

import (
	"fmt"
	"os"
	"os/exec"
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

	// Set AWS credentials and region
	os.Setenv("AWS_ACCESS_KEY_ID", awsAccessKeyID)
	os.Setenv("AWS_SECRET_ACCESS_KEY", awsSecretAccessKey)
	os.Setenv("AWS_DEFAULT_REGION", awsDefaultRegion)

	// Download artifacts from S3 bucket to the specified target
	cmd := exec.Command("aws", "s3", "cp", fmt.Sprintf("s3://%s/%s", awsBucketName, fetchDir), downloadTarget, "--recursive")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		os.Exit(1)
	}
}
