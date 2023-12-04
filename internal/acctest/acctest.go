package acctest

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
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

	Host01 = "TF_TEST_HOST_01"
	Host02 = "TF_TEST_HOST_02"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// RandomNameWithPrefix generates a random name with the given prefix.
// This is used in the acceptance tests where a unique name is required for the resource.
func RandomNameWithPrefix(prefix string) string {
	b := make([]byte, 6)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return fmt.Sprintf("%s-%s", prefix, b)
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

// TestAccBaseConfig_DhcpHosts creates a Terraform datasource config that allows you to filter by Hostname for 2 hosts
func TestAccBaseConfig_DhcpHosts() string {
	config := fmt.Sprintf(`
data "bloxone_dhcp_hosts" "test_02" {
	filters = {
		name = %q
	}
}`, Host02)
	return strings.Join([]string{TestAccBaseConfig_DhcpHost(), config}, "")
}

// TestAccBaseConfig_DhcpHost creates a Terraform datasource config that allows you to filter by Hostname for a single host
func TestAccBaseConfig_DhcpHost() string {
	return fmt.Sprintf(`
data "bloxone_dhcp_hosts" "test_01" {
	filters = {
		name = %q
	}
}`, Host01)
}
