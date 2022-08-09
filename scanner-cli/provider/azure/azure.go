// https://portal.azure.com
// https://github.com/Azure-Samples/azure-sdk-for-go-samples/blob/main/sdk/resourcemanager/storage/blob/main.go
package azure

import (
	"context"
	"log"
	"scanner-go/config"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

func Scan(app_config config.AppConfig) {
	sub := app_config.AzureSubscriptionId
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	containerItems, err := listBlobContainer(sub, ctx, cred)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("blob container list:")
	for _, item := range containerItems {
		log.Println("\t", *item.ID)
	}
}

func listBlobContainer(sub string, ctx context.Context, cred azcore.TokenCredential) (listItems []*armstorage.ListContainerItem, err error) {
	subscriptionID := sub // os.Getenv("AZURE_SUBSCRIPTION_ID")
	resourceGroupName := "storage-resource-group"
	storageAccountName := "ihasabucket"
	blobContainerClient, err := armstorage.NewBlobContainersClient(subscriptionID, cred, nil)
	if err != nil {
		return nil, err
	}

	containerItemsPager := blobContainerClient.NewListPager(resourceGroupName, storageAccountName, nil)

	listItems = make([]*armstorage.ListContainerItem, 0)
	for containerItemsPager.More() {
		pageResp, err := containerItemsPager.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		listItems = append(listItems, pageResp.ListContainerItems.Value...)
	}
	return
}
