#!/bin/bash

# Config
# Config
REMOTE_USER="root"
REMOTE_HOST="vps-spada-2"
REMOTE_PORT="22"
REMOTE_PATH="~/akademik-deployment/"
ZIP_FILE="deploy-package.zip"
INCLUDE_FILES=("deploy.sh" "api-spada-bridging" "api-spada-migrate")

# Clean previous zip
rm -f "$ZIP_FILE"

# Create zip package
echo "Creating zip package: $ZIP_FILE"
zip -r "$ZIP_FILE" "${INCLUDE_FILES[@]}"

if [ $? -ne 0 ]; then
    echo "Failed to create zip package."
    exit 1
fi

# Upload zip to remote server
echo "Uploading $ZIP_FILE to $REMOTE_USER@$REMOTE_HOST:$REMOTE_PATH"
scp -P "$REMOTE_PORT" "$ZIP_FILE" "$REMOTE_USER@$REMOTE_HOST:$REMOTE_PATH"

if [ $? -ne 0 ]; then
    echo "SCP upload failed."
    exit 1
fi

echo "Upload completed successfully."

# Run deploy.sh on remote server
echo "Running deploy.sh on remote server..."
ssh -p "$REMOTE_PORT" "$REMOTE_USER@$REMOTE_HOST" "cd $REMOTE_PATH && bash deploy.sh"

if [ $? -ne 0 ]; then
    echo "Remote deploy.sh execution failed."
    exit 1
fi

echo "Deployment script executed successfully on remote server."