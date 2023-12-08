# Use the official AWS CLI Docker image as the base image
FROM amazon/aws-cli:amd64

# Set the working directory
WORKDIR /workspace

# Copy the script into the container
COPY download_artifacts.sh /workspace/download_artifacts.sh

# Make the script executable
RUN chmod +x /workspace/download_artifacts.sh

# Entry point for the container
ENTRYPOINT ["/workspace/download_artifacts.sh"]

