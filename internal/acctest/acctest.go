package acctest

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"

	bloxoneclient "github.com/infobloxopen/bloxone-go-client/client"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/provider"
)

var (
	// BloxOneClient will be used to do verification tests
	BloxOneClient *bloxoneclient.APIClient

	// ProtoV6ProviderFactories are used to instantiate a provider during
	// acceptance testing. The factory function will be invoked for every Terraform
	// CLI command executed to create a provider server to which the CLI can
	// reattach.
	ProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"bloxone": providerserver.NewProtocol6WithError(provider.New("test", "none")()),
	}
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// RandomNameWithPrefix generates a random name with the given prefix.
// This is used in the acceptance tests where a unique name is required for the resource.
func RandomNameWithPrefix(prefix string) string {
	return fmt.Sprintf("%s-%s", prefix, RandomName())
}

func RandomName() string {
	b := make([]byte, 6)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func PreCheck(t *testing.T) {
	cspURL := os.Getenv("BLOXONE_CSP_URL")
	if cspURL == "" {
		t.Fatal("BLOXONE_CSP_URL must be set for acceptance tests")
	}

	apiKey := os.Getenv("BLOXONE_API_KEY")
	if apiKey == "" {
		t.Fatal("BLOXONE_API_KEY must be set for acceptance tests")
	}

	var err error
	BloxOneClient, err = bloxoneclient.NewAPIClient(bloxoneclient.Configuration{
		ClientName: "acceptance-test",
		CSPURL:     cspURL,
		APIKey:     apiKey,
	})
	if err != nil {
		t.Fatal("Cannot create bloxone client")
	}
}

// TestAccBaseConfig_DhcpHosts creates a Terraform datasource config that allows you to filter by tags
func TestAccBaseConfig_DhcpHosts() string {
	return `
data "bloxone_dhcp_hosts" "test" {
	tag_filters = {
		used_for = "Terraform Provider Acceptance Tests"
	}
}`
}
