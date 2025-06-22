#FROM golang:1.17-alpine as builder
#
#FROM sajari/docd
#RUN apt-get update && apt-get install -y \
#    poppler-utils \
#    wv \
#    unrtf \
#    tidy
#
#RUN apt update && apt install golang -y
#
#WORKDIR /app
#
#COPY go.mod go.sum ./
#
#COPY . .
#
#COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
#
#RUN go get code.sajari.com/docconv/...
#
#RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
#
#EXPOSE 8081
#
## Adding Env variables.
#ARG DEPLOYMENT_ID
#ENV APY_BUILD=${DEPLOYMENT_ID}
#
#ENTRYPOINT ["./main"]

# --------------------------------------------------

# Use the specified base image
FROM golang:1.21-alpine

# Update and upgrade system packages
RUN apk update && apk upgrade

# Add edge and community repositories
RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/main" >> /etc/apk/repositories && \
    echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories && \
    echo "http://dl-cdn.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories

# Install general dependencies
RUN apk add --no-cache \
    tiff \
    icu \
    icu-libs \
    icu-dev \
    leptonica-dev \
    tesseract-ocr \
    unrtf

# Set up dependencies for docconv
RUN apk add --no-cache \
    poppler-utils \
    wv \
    tidyhtml \
    git \
    ca-certificates

# Install dependencies for Chrome
RUN apk add --no-cache \
    nss \
    freetype \
    freetype-dev \
    harfbuzz \
    ca-certificates \
    ttf-freefont \
    libc6-compat

# Install Chrome
RUN apk add --no-cache chromium

# Set working directory
WORKDIR /app

# Copy required files to the container
COPY . .

# Install go dependencies and build the application
RUN go get github.com/JalfResi/justext && \
    go get -tags ocr code.sajari.com/docconv/... && \
    go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Expose the required port
EXPOSE 8081

# Add environment variables
ARG DEPLOYMENT_ID
ENV APY_BUILD=${DEPLOYMENT_ID}

# Set the entry point for the container
ENTRYPOINT ["./main"]