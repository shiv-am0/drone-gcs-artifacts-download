# drone-s3-artifacts-download
This plugin is designed to download objects/artifacts from an AWS S3 bucket. The primary goal of this plugin is to use AWS CLI to authenticate into an AWS account and download objects from an S3 bucket.

## Build and Run using terminal

Clone the repo and use the following commands to run the script directly from the terminal. Make sure to export the respective environment variables and have AWS CLI installed on your local machine.

```bash
./path/to/repo/download_artifacts_s3.sh
```

## Docker

Build the Docker image with the following commands. Using the following command, the image can be built for different OS and architecture. 

```
docker buildx create --use
docker buildx build --platform linux/amd64,linux/arm64 -t DOCKER_ORG/drone-s3-artifacts-download .
```

## Usage

Use the following command to run the container using docker
```bash
docker run --rm \
-e PLUGIN_AWS_ACCESS_KEY_ID=<"YourAccessKeyId"> \
-e PLUGIN_AWS_SECRET_ACCESS_KEY=<"YourSecretAccessKey"> \
-e PLUGIN_AWS_DEFAULT_REGION=<"YourAWSRegion"> \
-e PLUGIN_AWS_BUCKET_NAME=<"YourS3BucketName"> \
-e PLUGIN_FETCH_DIR=<"YourFetchDirectory"> \
-e PLUGIN_DOWNLOAD_TARGET=<"YourDownloadTarget"> \
DOCKER_ORG/drone-s3-artifacts-download
```

In Harness CI, the following YAML can be used to implement the plugin as a step
```yaml
              - step:
                  type: Plugin
                  name: drone-s3-artifacts-download
                  identifier: drones3artifactsdownload
                  spec:
                    connectorRef: account.harnessImage
                    image: harnesscommunity/drone-s3-artifacts-download
                    settings:
                      aws_access_key_id: <+secrets.getValue("awsaccesskeyid")>
                      aws_secret_access_key: <+secrets.getValue("awssecretaccesskey")>
                      aws_default_region: ap-south-1
                      aws_bucket_name: mybucket
                      download_target: /harness
                      fetch_dir: mydir
```
