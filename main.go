package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/aws/aws-sdk-go/service/cloudfront/sign"
	"io/ioutil"
	"log"
	url2 "net/url"
	"os"
	"time"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <s3ObjectPath>")
	}

	// Set your AWS credentials and CloudFront distribution ID
	accessKeyID := "sampleKeyID"               // CloudFront key-pair public key ID
	privKeyPath := "private_key.pem"           // Path to CloudFront key-pair pirvate key
	expires := time.Now().Add(1 * time.Minute) // expiry time
	s3ObjectPath := os.Args[1]
	cloudFrontDistributionHost := "sample.cloudfront.net"

	// Read the private key file - Instead of using the sign.LoadPEMPrivKeyFile reading into a byte slice due to error : x509: failed to parse private key (use ParsePKCS8PrivateKey instead for this key format)
	keyBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("failed to read private key file: %v", err)
	}

	// Decode the PEM block
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		log.Fatal("failed to parse PEM block containing the key")
	}

	// Parse the PKCS#8 private key
	privKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		log.Fatalf("failed to parse private key: %v", err)
	}

	// Generate the signed URL
	URLSigner := sign.NewURLSigner(accessKeyID, privKey.(*rsa.PrivateKey))

	rawURL := &url2.URL{
		Scheme: "https",
		Host:   cloudFrontDistributionHost, // CloudFront distribution host
		Path:   s3ObjectPath,               // path to s3 object
	}

	signedURL, err := URLSigner.Sign(rawURL.String(), expires)
	if err != nil {
		log.Fatal("generate signed url failed:", err)
	}

	fmt.Printf("access signedURL: %s", signedURL)
}
