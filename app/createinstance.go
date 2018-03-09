package main

// BEFORE RUNNING:
// ---------------
// 1. If not already done, enable the Compute Engine API
//    and check the quota for your project at
//    https://console.developers.google.com/apis/api/compute
// 2. This sample uses Application Default Credentials for authentication.
//    If not already done, install the gcloud CLI from
//    https://cloud.google.com/sdk/ and run
//    `gcloud beta auth application-default login`.
//    For more information, see
//    https://developers.google.com/identity/protocols/application-default-credentials
// 3. Install and update the Go dependencies by running `go get -u` in the
//    project directory.

import (
        "fmt"
        "log"

        "golang.org/x/net/context"
        "golang.org/x/oauth2/google"
        "google.golang.org/api/compute/v1"
)

func main() {
        ctx := context.Background()

        c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
        if err != nil {
                log.Fatal(err)
        }

        computeService, err := compute.New(c)
        if err != nil {
                log.Fatal(err)
        }

        // Project ID for this request.
        project := "aaron-project-197520" // TODO: Update placeholder value.

        // The name of the zone for this request.
        zone := "europe-west3-a" // TODO: Update placeholder value.

	rb := &compute.Instance{
		Name:        "test111",
		Description: "compute sample instance",
		MachineType: "https://www.googleapis.com/compute/v1/projects/aaron-project-197520/zones/europe-west3-a/machineTypes/n1-standard-1",
		Disks: []*compute.AttachedDisk{
			{
				AutoDelete: true,
				Boot:       true,
				Type:       "PERSISTENT",
				InitializeParams: &compute.AttachedDiskInitializeParams{
					DiskName:    "testdisk",
					SourceImage: "projects/ubuntu-os-cloud/global/images/ubuntu-1604-xenial-v20180306",
				},
			},
		},
		NetworkInterfaces: []*compute.NetworkInterface{
			&compute.NetworkInterface{
				AccessConfigs: []*compute.AccessConfig{
					&compute.AccessConfig{
						Type: "ONE_TO_ONE_NAT",
						Name: "External NAT",
					},
				},
				Network: "https://www.googleapis.com/compute/v1/projects/aaron-project-197520/global/networks/default",
			},
		},
		ServiceAccounts: []*compute.ServiceAccount{
			{
				Email: "default",
				Scopes: []string{
					compute.DevstorageFullControlScope,
					compute.ComputeScope,
				},
			},
		},
	}

        resp, err := computeService.Instances.Insert(project, zone, rb).Context(ctx).Do()
        if err != nil {
                log.Fatal(err)
        }

        // TODO: Change code below to process the `resp` object:
        fmt.Printf("%#v\n", resp)
}
