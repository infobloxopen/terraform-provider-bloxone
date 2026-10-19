package acctest

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/provider"
	universalddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
	"github.com/infobloxopen/universal-ddi-go-client/option"
)

const (
	letterBytes  = "abcdefghijklmnopqrstuvwxyz"
	defaultKey   = "managed_by"
	defaultValue = "terraform"
)

var (
	// BloxOneClient will be used to do verification tests
	BloxOneClient *universalddiclient.APIClient

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

// lookupEnvAny returns the value of the first set environment variable among names.
func lookupEnvAny(names ...string) string {
	for _, name := range names {
		if v, ok := os.LookupEnv(name); ok {
			return v
		}
	}
	return ""
}

func PreCheck(t *testing.T) {
	t.Helper()

	cspURL := lookupEnvAny("INFOBLOX_PORTAL_URL", "BLOXONE_CSP_URL")
	if cspURL == "" {
		t.Fatal("INFOBLOX_PORTAL_URL (or the deprecated BLOXONE_CSP_URL) must be set for acceptance tests")
	}

	apiKey := lookupEnvAny("INFOBLOX_PORTAL_KEY", "BLOXONE_API_KEY")
	if apiKey == "" {
		t.Fatal("INFOBLOX_PORTAL_KEY (or the deprecated BLOXONE_API_KEY) must be set for acceptance tests")
	}

	BloxOneClient = universalddiclient.NewAPIClient(
		option.WithClientName("terraform-acceptance-tests"),
		option.WithCSPUrl(cspURL),
		option.WithAPIKey(apiKey),
		option.WithDebug(true),
	)
}

func VerifyDefaultTag(resourceName string) resource.TestCheckFunc {
	return resource.TestCheckResourceAttr(resourceName, fmt.Sprintf("tags_all.%s", defaultKey), defaultValue)
}

// ReadTestdata reads a fixture file from internal/acctest/testdata/<service>/,
// shared across every service's tests. Used to keep large or sensitive-looking
// content (e.g. test certificates) out of test source files.
//
// The path is resolved relative to this source file (via runtime.Caller)
// rather than the working directory, since go test sets the working
// directory to the *calling* package's directory, not acctest's own -
// a plain "testdata/..." relative path would only work for tests that
// happen to live in this same package.
func ReadTestdata(t *testing.T, service, name string) string {
	t.Helper()
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatalf("unable to determine acctest package location")
	}
	path := filepath.Join(filepath.Dir(thisFile), "testdata", service, name)
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("unable to read %s: %v", path, err)
	}
	return string(content)
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
