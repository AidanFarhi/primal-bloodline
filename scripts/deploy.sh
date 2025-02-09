#!/bin/bash

ssh_key=./kashtrack-key-pair.pem

echo "starting deployment"

echo "building binary"
GOOS=linux GOARCH=amd64 go build -o kashtrack app.go

echo "stopping app"
ssh -o StrictHostKeyChecking=no -i $ssh_key ec2-user@kash-track.com sudo -n pkill -f kashtrack || true

echo "creating log file"
ssh -i $ssh_key ec2-user@kash-track.com touch /home/ec2-user/app/app.log

echo "copying binary and static files"
scp -i $ssh_key .env_prod ec2-user@kash-track.com:/home/ec2-user/app/.env
scp -i $ssh_key kashtrack ec2-user@kash-track.com:/home/ec2-user/app/
scp -r -i $ssh_key web ec2-user@kash-track.com:/home/ec2-user/app/

echo "starting app"
ssh -i $ssh_key ec2-user@kash-track.com "cd /home/ec2-user/app && echo 'sudo -n ./kashtrack > /dev/null 2>&1 &' | at now 2>/dev/null"

rm kashtrack
echo "deployment complete"
