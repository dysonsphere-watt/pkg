package pkg

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/goravel/framework/facades"
)

var azBlobClient *azblob.Client

// Sets up the Azure Blob Storage client for file read/writes
func SetupAzureBlob() error {
	// For default credentials, ensure AZURE_CLIENT_ID and AZURE_TENANT_ID are set as environment variables
	// Then depending on the type of authentication,
	// `AZ_CLIENT_SECRET``
	// or `AZURE_CLIENT_CERTIFICATE_PATH` and `AZURE_CLIENT_CERTIFICATE_PASSWORD`
	// must be set.
	azCred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return err
	}

	serviceURL := facades.Config().GetString("AZURE_STORAGE_SERVICE_URL", "https://fail.blob.core.windows.net/")
	azBlobClient, err = azblob.NewClient(serviceURL, azCred, nil)
	if err != nil {
		return err
	}

	return nil
}

func GetAzBlobClient() *azblob.Client {
	return azBlobClient
}

// Uploads a file to the container in Azure Blob Storage
func UploadMultipartFile(container string, fileHeader *multipart.FileHeader) (string, error) {
	blobName := fmt.Sprintf("%s-%d", GenerateRandomAlphaNum(16), time.Now().Unix())

	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileBuffer, err := ReadFileHeader(fileHeader)
	if err != nil {
		return "", err
	}

	_, err = azBlobClient.UploadBuffer(context.TODO(), container, blobName, fileBuffer, nil)
	if err != nil {
		return "", err
	}

	blobURL := fmt.Sprintf("%s%s/%s", azBlobClient.URL(), container, blobName)
	return blobURL, nil
}
