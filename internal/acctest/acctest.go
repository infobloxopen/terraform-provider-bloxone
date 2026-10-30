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

// RandomPublicIP returns a random public (non-RFC-1918) IPv4 address.
// Use this when the API rejects private IP ranges (e.g. FW network lists).
func RandomPublicIP() string {
	for {
		a := rand.Intn(256)
		b := rand.Intn(256)
		c := rand.Intn(256)
		d := rand.Intn(256)
		if a == 0 || a == 10 || a == 127 || a >= 224 {
			continue
		}
		if a == 100 && b >= 64 && b <= 127 { // RFC 6598 shared address
			continue
		}
		if a == 169 && b == 254 { // link-local
			continue
		}
		if a == 172 && b >= 16 && b <= 31 { // RFC 1918
			continue
		}
		if a == 192 && b == 168 { // RFC 1918
			continue
		}
		if a == 192 && b == 0 && c == 2 { // RFC 5737 TEST-NET-1
			continue
		}
		if a == 198 && b == 51 && c == 100 { // RFC 5737 TEST-NET-2
			continue
		}
		if a == 203 && b == 0 && c == 113 { // RFC 5737 TEST-NET-3
			continue
		}
		return fmt.Sprintf("%d.%d.%d.%d", a, b, c, d)
	}
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

func RandomAlphaNumeric(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
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
	path := TestdataPath(t, service, name)
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("unable to read %s: %v", path, err)
	}
	return string(content)
}

// TestdataPath returns the absolute path of a fixture file in
// internal/acctest/testdata/<service>/, for configs that load it themselves
// (e.g. via Terraform's file() function), since Terraform runs acceptance
// test configs from a temporary working directory.
func TestdataPath(t *testing.T, service, name string) string {
	t.Helper()
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatalf("unable to determine acctest package location")
	}
	return filepath.Join(filepath.Dir(thisFile), "testdata", service, name)
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
