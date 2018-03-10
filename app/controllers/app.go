package controllers

import (
	"github.com/revel/revel"

	"time"

        "fmt"
        "log"

        "golang.org/x/net/context"
        "golang.org/x/oauth2/google"
        "google.golang.org/api/compute/v1"

        "strings"
        "os/exec"
)


// project ID
var projectvar string = "aaron-project-197520"
// zone ID
var zonevar string = "europe-west3-a"
// Instance NAME
var instancevar string = "test111"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Healthcheck() revel.Result {
    return c.Render()
}


func (c App) Create(username, password string) revel.Result {
  // Checking the credential.
  if username == "hylee" && password == "aaron11!" {
  	// Creating an instance.
        createinstance(username, password)
	// return the IP 
        var resultmsg string = checkingip()
  }
  
  var resultmsg string = "Please check your credential."
  return c.Render(resultmsg)

}

//creating an instance.
func createinstance(username,password string) {

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
        project := projectvar 

        // The name of the zone for this request.
        zone := zonevar

        rb := &compute.Instance{
                Name:        instancevar,
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

        fmt.Printf("%#v\n", resp)
}

// Checking the IP
func checkingip() string {
	// wait for the time to create an instance.
	time.Sleep(20000 * time.Millisecond)	
	
	// Checking the IP
	cmdip := "gcloud compute instances list |grep 'test111' |awk '{print $5}'"
	outip, _ := exec.Command("bash", "-c", cmdip).Output()

	// Ping test
	out, _ := exec.Command("ping", string(outip), "-c 10", "-i 5", "-w 15").Output()
	if strings.Contains(string(out), "64 bytes from") {
    		 return string(outip)
	} else {
    		 return "error"
	}
}

