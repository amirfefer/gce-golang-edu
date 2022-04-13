package compute

import (
	"fmt"
	"context"
	"github.com/create-go-app/fiber-go-template/app/models"
	compute "cloud.google.com/go/compute/apiv1"
	"google.golang.org/api/option"
)

func CreateClient(c *models.Compute) (*compute.InstancesClient, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile(c.JsonPath)
	instancesClient, err := compute.NewInstancesRESTClient(ctx, opt)
	if err != nil {
		return instancesClient, fmt.Errorf("NewInstancesRESTClient: %v", err)
	}
	return instancesClient, nil
}

func CloseClient(c *compute.InstancesClient) {
	c.Close()
}