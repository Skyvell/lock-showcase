# AWS S3 Locking Example

This Go application demonstrates the use of AWS S3 locks with the `s3lock` package to create 10 competing lock instances. Each lock attempts to acquire a shared lock on an S3 object, and only one lock can acquire it at a time. The locks are managed using the AWS SDK for Go (v2).

## Prerequisites

Before you begin, ensure you have the following set up:

- [Go](https://golang.org/dl/) installed on your machine.
- An AWS account with necessary IAM permissions to interact with S3.
- An S3 bucket (`versioningbucketcrossbreed` in this example) with versioning enabled.
- AWS credentials configured on your machine or in the environment.

## Setup

1. Clone the repository (or save the file locally).

2. Install dependencies:
   - You will need the AWS SDK for Go v2 and the `s3lock` package.
   Run the following command to initialize the Go module and install dependencies:

   ```bash
   go mod init your-module-name
   go get github.com/aws/aws-sdk-go-v2/config
   go get github.com/skyvell/locksv2/s3lock
   ```

3. Ensure that your AWS credentials are correctly configured. You can do this using the AWS CLI or by setting environment variables:
   - `AWS_ACCESS_KEY_ID`
   - `AWS_SECRET_ACCESS_KEY`

   You can configure AWS CLI by running:

   ```bash
   aws configure
   ```

4. Make sure that the S3 bucket (`versioningbucketcrossbreed`) exists and that S3 object versioning is enabled on the bucket.

## How to Run

1. In the `main.go` file, make sure to set the correct bucket name and key for your S3 object:

   ```go
   const bucketName = "versioningbucketcrossbreed"
   const key = "key-201"
   ```

2. Build and run the Go program:

   ```bash
   go run main.go
   ```

   The program will attempt to create 10 competing lock instances. Each lock will try to acquire the lock on the S3 object and execute code if successful.

3. If successful, you'll see output similar to the following:

   ```bash
   Lock: Lock0 acquired lock. Executing:
   Sleeping for 1 second.

   Lock: Lock1 acquired lock. Executing:
   Sleeping for 1 second.
   ```

## Additional Information

- The lock retry duration is set to 60 seconds. If a lock cannot be acquired within this time, the program will skip it.
- Each successful lock holds the lock for 1 second before releasing it.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```

Feel free to modify the bucket name, key, and other configuration values as per your needs.
