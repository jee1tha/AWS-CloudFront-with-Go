# AWS-CloudFront-with-Go

This Go application generates a signed URL for AWS CloudFront, allowing controlled access to private content stored on S3 through CloudFront distribution.

## Requirements

- Go 1.21.6
- AWS SDK for Go
- RSA private key in PEM format (PKCS#8)

## Configuration

Before running the application, update the following configuration parameters in the `main` function:
- `accessKeyID`: Your CloudFront key-pair ID.
- `privKeyPath`: Path to your RSA private key file.
- `expires`: Set the expiry time for the signed URL.
- `cloudFrontDistributionHost`: CloudFront distribution host

## Usage

1. Clone this repository.
2. Place your RSA private key in the root directory or update the `privKeyPath` to its location.
3. Run the application:

```bash
go run main.go "upload/sample_image.jpeg"
