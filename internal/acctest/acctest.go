package acctest

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/provider"
)

var (
	// ProtoV6ProviderFactories are used to instantiate a provider during
	// acceptance testing. The factory function will be invoked for every Terraform
	// CLI command executed to create a provider server to which the CLI can
	// reattach.
	ProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"bloxone": providerserver.NewProtocol6WithError(provider.New("test", "none")()),
	}
)

func PreCheck(t *testing.T) {
	if cspURL := os.Getenv("BLOXONE_CSP_URL"); cspURL == "" {
		t.Fatal("BLOXONE_CSP_URL must be set for acceptance tests")
	}

	if apiKey := os.Getenv("BLOXONE_API_KEY"); apiKey == "" {
		t.Fatal("BLOXONE_API_KEY must be set for acceptance tests")
	}
}
