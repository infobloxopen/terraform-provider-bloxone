package acctest

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	bloxoneclient "github.com/infobloxopen/bloxone-go-client/client"
	"github.com/infobloxopen/bloxone-go-client/option"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/provider"
)

const (
	letterBytes  = "abcdefghijklmnopqrstuvwxyz"
	defaultKey   = "managed_by"
	defaultValue = "terraform"
)

var (
	// BloxOneClient will be used to do verification tests
	BloxOneClient *bloxoneclient.APIClient

	// ProtoV6ProviderFactories are used to instantiate a provider during
	// acceptance testing. The factory function will be invoked for every Terraform
	// CLI command executed to create a provider server to which the CLI can
	// reattach.
	ProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"bloxone": providerserver.NewProtocol6WithError(provider.New("test", "test")()),
	}
	ProtoV6ProviderFactoriesWithTags = map[string]func() (tfprotov6.ProviderServer, error){
		"bloxone": providerserver.NewProtocol6WithError(provider.NewWithTags(map[string]string{defaultKey: defaultValue})()),
	}
)

// RandomNameWithPrefix generates a random name with the given prefix.
// This is used in the acceptance tests where a unique name is required for the resource.
func RandomNameWithPrefix(prefix string) string {
	return fmt.Sprintf("%s-%s", prefix, RandomName())
}

func RandomIP() string {
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
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

	BloxOneClient = bloxoneclient.NewAPIClient(
		option.WithClientName("terraform-acceptance-tests"),
		option.WithCSPUrl(cspURL),
		option.WithAPIKey(apiKey),
		option.WithDebug(true),
	)
}

func VerifyDefaultTag(resourceName string) resource.TestCheckFunc {
	return resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("tags_all.%s", defaultKey), defaultValue)
}

// TestAccBase_DhcpHosts creates a Terraform datasource config that allows you to filter by tags
func TestAccBase_DhcpHosts() string {
	return `
data "bloxone_dhcp_hosts" "test" {
	tag_filters = {
		used_for = "Terraform Provider Acceptance Tests"
	}
}`
}
