# AWS S3 Locking Demonstration

This Go application demonstrates the use of AWS S3 locks with the `s3lock` package to create 10 competing lock instances. Each lock attempts to acquire a shared lock on an S3 object, and only one lock can acquire it at a time. The locks are managed using the AWS SDK for Go (v2).

## Prerequisites
- Access to a AWS account with a bucket and versioning enabled. Configure the constants in the script according to your aws settings.
- Aws command line interface installed and configured with credentials to your AWS account.

## Installation
To retreive the library:
```
go get github.com/skyvell/locksv2/s3lock
```

## Showcase
Run the script with:
```go
go run main.go
```
