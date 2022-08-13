// https://portal.azure.com
// https://github.com/Azure-Samples/azure-sdk-for-go-samples/blob/main/sdk/resourcemanager/storage/blob/main.go
package azure

import (
	"context"
	"log"
	"scanner-go/config"
	_storage "scanner-go/storage" // TODO fix name

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

type ScanAzure struct {
	Config  *config.AppConfig
	Store   *_storage.Storage
	Sub     string
	Context context.Context
	_cred   *azidentity.DefaultAzureCredential
}

//func (scan *ScanAzure) NewScanAzure(app_config config.AppConfig, store *_storage.Storage) *ScanAzure {
func NewScanAzure(app_config config.AppConfig, store *_storage.Storage) *ScanAzure {
	//sub := app_config.AzureSubscriptionId
	// cred, err := azidentity.NewDefaultAzureCredential(nil)
	// if err != nil {
	// 	return err
	// }
	//ctx := context.Background()

	return &ScanAzure{
		Config:  &app_config,
		Store:   store,
		Sub:     app_config.AzureSubscriptionId,
		Context: context.Background(),
	}
}

func Scan(app_config config.AppConfig, store *_storage.Storage) error {
	log.Print("Scan Azure")

	scan := NewScanAzure(app_config, store)

	cred, err := scan.Cred()
	if err != nil {
		log.Printf("credential error %+v", err)
		return err
	}

	//if err := scan.ListBlobContainer(sub, ctx, cred, store); err != nil {
	if err := scan.ListBlobContainer(cred, store); err != nil {
		log.Printf("Azure Blob error %v", err)
	}

	return nil
}

func (scan *ScanAzure) Cred() (*azidentity.DefaultAzureCredential, error) {
	if scan._cred != nil {
		return scan._cred, nil
	}

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Printf("credential error %+v", err)
		return nil, err
	}
	return cred, nil
}

func (scan *ScanAzure) ListBlobContainer(cred azcore.TokenCredential, store *_storage.Storage) error {
	// func listBlobContainer(sub string, ctx context.Context, cred azcore.TokenCredential, store *_storage.Storage) error {
	subscriptionID := scan.Sub //sub // os.Getenv("AZURE_SUBSCRIPTION_ID")
	resourceGroupName := "storage-resource-group"
	storageAccountName := "ihasabucket"
	blobContainerClient, err := armstorage.NewBlobContainersClient(subscriptionID, cred, nil)
	if err != nil {
		return err
	}

	containerItemsPager := blobContainerClient.NewListPager(resourceGroupName, storageAccountName, nil)

	for containerItemsPager.More() {
		pageResp, err := containerItemsPager.NextPage(scan.Context)
		if err != nil {
			return err
		}
		store.Save(store.NewStorageData(pageResp))
	}

	return nil
}
