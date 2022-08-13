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

func Scan(app_config config.AppConfig, store *_storage.Storage) error {
	sub := app_config.AzureSubscriptionId
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return err
	}
	ctx := context.Background()

	if err := listBlobContainer(sub, ctx, cred, store); err != nil {
		log.Printf("Azure Blob error %v", err)
	}

	return nil
}

func listBlobContainer(sub string, ctx context.Context, cred azcore.TokenCredential, store *_storage.Storage) error {
	subscriptionID := sub // os.Getenv("AZURE_SUBSCRIPTION_ID")
	resourceGroupName := "storage-resource-group"
	storageAccountName := "ihasabucket"
	blobContainerClient, err := armstorage.NewBlobContainersClient(subscriptionID, cred, nil)
	if err != nil {
		return err
	}

	containerItemsPager := blobContainerClient.NewListPager(resourceGroupName, storageAccountName, nil)

	for containerItemsPager.More() {
		pageResp, err := containerItemsPager.NextPage(ctx)
		if err != nil {
			return err
		}
		store.Save(store.NewStorageData(pageResp))
	}

	return nil
}
