#!/bin/bash

# Check if required environment variables are set
if [ -z "$PLUGIN_AWS_ACCESS_KEY_ID" ] || [ -z "$PLUGIN_AWS_SECRET_ACCESS_KEY" ] || [ -z "$PLUGIN_AWS_DEFAULT_REGION" ] || [ -z "$PLUGIN_AWS_BUCKET_NAME" ] || [ -z "$PLUGIN_FETCH_DIR" ] || [ -z "$PLUGIN_DOWNLOAD_TARGET" ]; then
  echo "Error: Please provide all required environment variables."
  exit 1
fi

# Set AWS credentials and region
export AWS_ACCESS_KEY_ID="$PLUGIN_AWS_ACCESS_KEY_ID"
export AWS_SECRET_ACCESS_KEY="$PLUGIN_AWS_SECRET_ACCESS_KEY"
export AWS_DEFAULT_REGION="$PLUGIN_AWS_DEFAULT_REGION"

# Download artifacts from S3 bucket to the specified target
aws s3 cp "s3://$PLUGIN_AWS_BUCKET_NAME/$PLUGIN_FETCH_DIR" "$PLUGIN_DOWNLOAD_TARGET" --recursive

