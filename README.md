# http2smtp
take http request and send to smtp server

# Usage
## Step 1: get smtp from yahoo
* Register a yahoo email if you don't have.
* Go to https://login.yahoo.com/myaccount/security/
* Click Generate and manage app passwords
* Enter the app name, it is only a description.
* Click Generate password
* Write down the password generated, we need it.

## Step 2: create .env file
```
PORT=8888
SMTP_HOST=smtp.mail.yahoo.com
SMTP_PORT=465
SMTP_USER=<username>@yahoo.com
SMTP_PASSWORD=<password-from-step-1>
```

## Step 3: create docker-compose file
```
services:
  http2smtp:
    image: jindongh/http2smtp:latest
    restart: unless-stopped
    env_file:
      .env
    ports:
      - "8888:8888"
```

## Step 4: Start the docker
```
docker compose up -d
```

## Step 5: Verify
```
YAHOO_EMAIL=<user>@yahoo.com
curl "localhost:8888?to=${YAHOO_EMAIL}&subject=test%20http2smtp&content=hello%20from%20http2smtp"
```
