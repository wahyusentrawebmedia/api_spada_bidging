#!/bin/bash
set -e

# Configurable variables (edit as needed)
APP_DIR="/opt/akademik-spada"
DEPLOY_DIR="$HOME/akademik-deployment"
SERVICE_NAME="api-spada-bridging"
MIGRATE_BIN="api-spada-migrate"
BRIDGING_BIN="api-spada-bridging"
ENV_FILE=".env"
SERVICE_FILE="${SERVICE_NAME}.service"

# # Environment variables (replace with actual secrets or use export before running)
# DB_HOST="${DB_HOST:-localhost}"
# DB_USER="${DB_USER:-admin}"
# DB_PASS="${DB_PASS:-YtvAQGQx0yDX78li}"
# DB_PORT="${DB_PORT:-5433}"
# DB_NAME="${DB_NAME:-database_akademik_user}"
# URL_AKADEMIK_AUTH="${URL_AKADEMIK_AUTH:-https://devakademik-pengaturan.sentrawebmedia.com}"

echo "== Building application binaries =="
GOOS=linux GOARCH=amd64 go build -o $BRIDGING_BIN ./cmd
GOOS=linux GOARCH=amd64 go build -o $MIGRATE_BIN ./cmd/migrate

# echo "== Generating environment file =="
# cat > $ENV_FILE <<EOF
# DB_HOST=$DB_HOST
# DB_USER=$DB_USER
# DB_PASS=$DB_PASS
# DB_PORT=$DB_PORT
# DB_NAME=$DB_NAME
# URL_AKADEMIK_AUTH=$URL_AKADEMIK_AUTH
# EOF

echo "== Preparing deployment directory =="
mkdir -p $DEPLOY_DIR
cp $BRIDGING_BIN $DEPLOY_DIR/
cp $MIGRATE_BIN $DEPLOY_DIR/
# cp $ENV_FILE $DEPLOY_DIR/

echo "== Creating systemd service file =="
cat > $DEPLOY_DIR/$SERVICE_FILE <<EOF
[Unit]
Description=Akademik User Service
After=network.target

[Service]
Type=simple
User=www-data
WorkingDirectory=$APP_DIR
ExecStart=$APP_DIR/$BRIDGING_BIN
Restart=on-failure
RestartSec=5
StandardOutput=journal
StandardError=journal
Environment=GIN_MODE=release

[Install]
WantedBy=multi-user.target
EOF

echo "== Deploying to $APP_DIR =="
sudo mkdir -p $APP_DIR
sudo cp $DEPLOY_DIR/$BRIDGING_BIN $APP_DIR/
sudo cp $DEPLOY_DIR/$MIGRATE_BIN $APP_DIR/
# sudo cp $DEPLOY_DIR/$ENV_FILE $APP_DIR/
sudo chmod +x $APP_DIR/$BRIDGING_BIN
sudo chown -R www-data:www-data $APP_DIR

echo "== Running migration =="
cd $DEPLOY_DIR
./$MIGRATE_BIN

echo "== Installing systemd service =="
sudo cp $DEPLOY_DIR/$SERVICE_FILE /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable $SERVICE_NAME

echo "== Restarting service =="
sudo systemctl stop $SERVICE_NAME || true
sudo systemctl start $SERVICE_NAME

echo "== Checking service status =="
sudo systemctl status $SERVICE_NAME --no-pager

echo "== Deployment completed successfully! =="
