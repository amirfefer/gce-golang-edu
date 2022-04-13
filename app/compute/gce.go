package compute

import (
	"fmt"
	"context"
	"github.com/create-go-app/fiber-go-template/app/models"
	compute "cloud.google.com/go/compute/apiv1"
	"google.golang.org/api/iterator"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
)
// get all zones for a project
func ZoneList(projectName string) ([]*computepb.Zone, error) {
	ctx := context.Background()
	zones := []*computepb.Zone{}
	c, err := compute.NewZonesRESTClient(ctx)
	if err != nil {
		return zones, err
	}
	defer c.Close()

	req := &computepb.ListZonesRequest{
		Project:    projectName,
	}
	it := c.List(ctx, req)
	for {
		item, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return zones, err
		}
		zones = append(zones, item)
	}
	return zones, nil
}


func ListInstances(client *compute.InstancesClient, compute *models.Compute) ([]*models.GCEInstance, error) {
	GCEinstances := []*models.GCEInstance{}
	ctx := context.Background()

	// Use the `MaxResults` parameter to limit the number of results that the API returns per response page.
	req := &computepb.ListInstancesRequest{
		Project:    compute.ProjectID,
		Zone:       compute.Zone,
	}

	it := client.List(ctx, req)
	// Despite using the `MaxResults` parameter, you don't need to handle the pagination
	// yourself. The returned iterator object handles pagination
	// automatically, returning separated pages as you iterate over the results.
	for {
		instance, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

			GCEinstances = append(GCEinstances, &models.GCEInstance{Name: instance.GetName(), Zone: instance.GetZone(), ID: instance.GetId(), Status: instance.GetStatus()})
		}
	
	return GCEinstances, nil
}

func StartInstance(client *compute.InstancesClient, instance *models.GCEInstance) error {
	req := &computepb.StartInstanceRequest{
		Project:  instance.Compute.ProjectID,
		Zone:     instance.Compute.Zone,
		Instance: instance.Name,
	}
    ctx := context.Background()
	op, err := client.Start(ctx, req)
	if err != nil {
		return fmt.Errorf("unable to start instance: %v", err)
	}

	if err = op.Wait(ctx); err != nil {
		return fmt.Errorf("unable to wait for the operation: %v", err)
	}

	fmt.Printf("Instance started\n")

	return nil
}

// stopInstance stops a started Google Compute Engine instance
func StopInstance(client *compute.InstancesClient, instance *models.GCEInstance) error {
	ctx := context.Background()

	req := &computepb.StopInstanceRequest{
		Project:  instance.Compute.ProjectID,
		Zone:     instance.Compute.Zone,
		Instance: instance.Name,
	}

	op, err := client.Stop(ctx, req)
	if err != nil {
		return fmt.Errorf("unable to stop instance: %v", err)
	}

	if err = op.Wait(ctx); err != nil {
		return fmt.Errorf("unable to wait for the operation: %v", err)
	}

	fmt.Printf("Instance stopped\n")

	return nil
}
