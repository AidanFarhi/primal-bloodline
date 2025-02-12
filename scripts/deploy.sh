#!/bin/bash

ssh_key=./FIXME.pem

echo "starting deployment"

echo "building binary"
GOOS=linux GOARCH=amd64 go build -o FIXME app.go

echo "stopping app"
ssh -o StrictHostKeyChecking=no -i $ssh_key ec2-user@FIXME.com sudo -n pkill -f FIXME || true

echo "creating log file"
ssh -i $ssh_key ec2-user@FIXME.com touch /home/ec2-user/app/app.log

echo "copying binary and static files"
scp -i $ssh_key .env_prod ec2-user@FIXME.com:/home/ec2-user/app/.env
scp -i $ssh_key FIXME ec2-user@FIXME.com:/home/ec2-user/app/
scp -r -i $ssh_key web ec2-user@FIXME.com:/home/ec2-user/app/

echo "starting app"
ssh -i $ssh_key ec2-user@FIXME.com "cd /home/ec2-user/app && echo 'sudo -n ./FIXME > /dev/null 2>&1 &' | at now 2>/dev/null"

rm FIXME
echo "deployment complete"
