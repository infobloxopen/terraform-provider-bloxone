package cdc_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
	"github.com/infobloxopen/universal-ddi-go-client/cdc"
)

// Self-signed, non-sensitive test certificates used only to satisfy the CDC
// API's x509 PEM validation for ca_certificate/certificate fields.
const testAccSourceApiCertPEM = `-----BEGIN CERTIFICATE-----
MIIDDTCCAfWgAwIBAgIUD2dETKme6j/Jd8qywdTYLCxAKJgwDQYJKoZIhvcNAQEL
BQAwFjEUMBIGA1UEAwwLdGYtYWNjLXRlc3QwHhcNMjYxMDE2MDc1MDUxWhcNMzYx
MDEzMDc1MDUxWjAWMRQwEgYDVQQDDAt0Zi1hY2MtdGVzdDCCASIwDQYJKoZIhvcN
AQEBBQADggEPADCCAQoCggEBAM3Cgl0P9Ey69fqcq4QV97/EjW/X0x/JUX01CjU8
du8MdjOq2lfr5k1LXAn12BDOZjYc2f40MC1Y7Di0ZPWgJvky1UJJbuGC6tUlQ9TT
JVtKw5ukJzCPqojDy0Hz3EhpfgM+hBjhR3aDCAHZSCplv1G1G6JM0eqVeb5hvdpU
GOiPNAclkJenGOfFJzAifz9qZT07UyX037/48CuhvPCImM72hUHzYILOqYFXAxCS
wx4NmzPvGtTu1VkKvbsBSVrxJfu9mircsQD3ge9yUSdqKzp7WEl4d5fCW9cr/k7X
f2rX0MLZCpPPAX8EJVC1nkYdgpi6w4THrkMm3iz32ZOjUaECAwEAAaNTMFEwHQYD
VR0OBBYEFOIoE67RqlWnW6Soyb7LTVNHChRyMB8GA1UdIwQYMBaAFOIoE67RqlWn
W6Soyb7LTVNHChRyMA8GA1UdEwEB/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEB
AAHbsr9N/95pmxyXlxp69YwiD8Vdgw4Hj4nLa2dxKSmHmEAz6WRtHI94nKnd1cHH
d5tsHLKLOe2l6/7aCXYvsGenTh7HFCZlBg/bWYTVf+DZtCP8x59NkzOmdFfmwDo7
Ley97lbNZEssA87H9PoDSSnNC9irxpvq9q2oI1qjNoxMCvL2iBbH5/IOvjZj/Ace
2uohH4khyDR7kuRUCmxgnQW4po2D6BjUWfs0FrrCbKk5oFl33ilepLwOmDWbP5RH
5pPEmH14Ug1X3lwBCmaM88/stzU5r/ulofVnBnQYkbvgBU54h5mZUPKHhvZsOqm/
6qP1HFvy0k/V9Zc0fBIoS+k=
-----END CERTIFICATE-----
`

const testAccSourceApiCertPEMUpdated = `-----BEGIN CERTIFICATE-----
MIIDHTCCAgWgAwIBAgIUDL7FwIy1JzRvRy3FYesBG5SY0ZswDQYJKoZIhvcNAQEL
BQAwHjEcMBoGA1UEAwwTdGYtYWNjLXRlc3QtdXBkYXRlZDAeFw0yNjEwMTYwNzUw
NTZaFw0zNjEwMTMwNzUwNTZaMB4xHDAaBgNVBAMME3RmLWFjYy10ZXN0LXVwZGF0
ZWQwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCTrx1fI1V3Lk5jU7qo
pPNTOUXpunmSElTQdGk6CpkowwT8oGYX9FtoE2+Cgr+IKCdJZqivtknyAracYuIN
gynyT7OtAYCsq4pbZEPSut0TjobrQLC4RpwW4NTQxa8qhuTBNvAvrjkZSBTunmsp
3mghBAOzX/9H+EePa4eECZsFxhb6UqRHkL1Vd3A+yFc/JkrKNUpGXrn9m+JcVwi1
8UreNBwD3hqZwByRjkhOBppLhvmFTE+n9KLabHOTrkUxDt0zRtmvxv/iIknbzbF7
WOWNAHCWKZ4w+m4WKTdTmlJPnlMopSL62A/WDLABQEaTcniTeKP/vMNeArp2Otj/
LTHtAgMBAAGjUzBRMB0GA1UdDgQWBBS4sKpHZgcPAXMgB0kZzdb3WzS4XzAfBgNV
HSMEGDAWgBS4sKpHZgcPAXMgB0kZzdb3WzS4XzAPBgNVHRMBAf8EBTADAQH/MA0G
CSqGSIb3DQEBCwUAA4IBAQBs8J2v72qVT/V37WRvfCP2Fi9z50c+WIDTtIQY4AXw
D6BJzdPQC6wkBQ7l8TduMtsK0m/GCiEK71FkPAdvmewYV1dPNZwvomt7+q0NjwaK
ff8UXTLCvgo3+vSERqs/8IMgPpXkjudtnPvwRKBvBvBVWV1PcvkwhzZ/6uflzTy+
+vNwIY9EueajIYeLsUEVy7etsNuaiB9O1vbPBA2ZZc9Pq5wPJEM9nBTgJi2+orcq
iof6iugolWPASdVe4lI/tvM5HidCIdtf1RvtaJhoHZ8whIMn/z5rq0q3KXs6hDFg
/N8aii5n3bPla9rVBy8kTbgCSyJTCsysEvXaH1gZXnAl
-----END CERTIFICATE-----
`

// testAccSourceApiCertKeyPEM/Updated are PEM bundles (private key + cert) for
// the "certificate" field, which the API validates as a client cert+key pair,
// unlike "ca_certificate" which only requires a certificate.
const testAccSourceApiCertKeyPEM = `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDNwoJdD/RMuvX6
nKuEFfe/xI1v19MfyVF9NQo1PHbvDHYzqtpX6+ZNS1wJ9dgQzmY2HNn+NDAtWOw4
tGT1oCb5MtVCSW7hgurVJUPU0yVbSsObpCcwj6qIw8tB89xIaX4DPoQY4Ud2gwgB
2UgqZb9RtRuiTNHqlXm+Yb3aVBjojzQHJZCXpxjnxScwIn8/amU9O1Ml9N+/+PAr
obzwiJjO9oVB82CCzqmBVwMQksMeDZsz7xrU7tVZCr27AUla8SX7vZoq3LEA94Hv
clEnais6e1hJeHeXwlvXK/5O139q19DC2QqTzwF/BCVQtZ5GHYKYusOEx65DJt4s
99mTo1GhAgMBAAECggEAAKGvYcFp8rmuMJyTgDPLD2PMkAJ9wOlkHnoaXmshmFPq
9vdNdUX27oIj4FnzVas0hXJVjvddCylB4bNYDf2gW6p2x3Mklnz0OJRFNDLeY1HZ
mNw090FcTiV/tbjzElEp6bD4RH/zMV9vQluG4lkZRb3EaHKg6v3rd5Si5GKBXbwK
uDyxWntByYfjo54KSiSWyxDswVw2z8sDsI9neJm7STwAt7rYrEpyeYDlTy0IrxmJ
j369UDl9qznNL9sNRefanUViqaasWQ3YaHbTmKScXTHY5rQs+b7Au8kXR8SpwvD7
RvK1NzpaOym3D8WI2w1gBEKWrxQEj4lUkVA3xZpb4QKBgQDqH+1UsV7L5j4KUSGZ
viMEApa3p9W9HhMGXPMfXXU6yqZPQsPUHKn+NwYn0l+YyYcaVEeiyYFQvuGdau2Z
C7ZIUUb7uzjY4YWemLY2+SQKykmPuBunejWao/mZpY2H0sJuR+lbzJjJD9fhjrAg
KYpq1yTYvtricJavKMTZz7DlGQKBgQDg/B3RoTXuQcfMRChAr3VGGG4dJRSGGK5d
uqAnOApkNgAZffALDtGdJIP/a5NsVYQtsNLEqDQs8V6vfV+q2Vrd7lWdwN6YoQB2
/O6BFgsnr1egaqnJpCb5x+MykuDxGScihEmdd80SqVsQcpM5EeNjugc/ZjmcPnJb
lfhaImAZyQKBgQCXAWyz+Yro/GH7iYSggvBWv8yyPAAx8FnY+s6c5kycVOSbrrZb
8vj5zEeeG6QpCaKoKr/hkRJlIBiozbTdx54UWsfryT/LMvpZ5qlwH5VNmFeF6Pt9
M+erDCQLI6p7dkaNlrg+y2GoHR3ij4Ot2t06N/W+Dqti7bdi1mrRCZJQ0QKBgErz
kteIdlgGcBJbznEmQtJ0RxLMjJbT2Br/lmUaZVPXLZdQfkygxjIsvRC1fOqFXsB9
/OlpA+DJFHswMFPRj1UKFlm0rDN6x5XcU3SLrHGvF2kov6GeXyGtdsIdbnTDpWb5
XTDpfVcvNlg3Y1aK4A8UJV/4pU6o+x4Jyy/z7cIJAoGAYXApLX0f2Nr6hvEkC6Be
5mYL96YvL7RgrdGm3FF/cVs+g1XnvLxt0aIQ09L1aOLPJjAMvLpyLlk/vzVU9qZw
X5wKn+4hS99zfUEgTeoIf0dtVpvAUX2jllIyLNGai8RKzEIueiyMzEvVjqwc0Zmt
Lwy14dFmBCRV2Yk2qD6SyDY=
-----END PRIVATE KEY-----
-----BEGIN CERTIFICATE-----
MIIDDTCCAfWgAwIBAgIUD2dETKme6j/Jd8qywdTYLCxAKJgwDQYJKoZIhvcNAQEL
BQAwFjEUMBIGA1UEAwwLdGYtYWNjLXRlc3QwHhcNMjYxMDE2MDc1MDUxWhcNMzYx
MDEzMDc1MDUxWjAWMRQwEgYDVQQDDAt0Zi1hY2MtdGVzdDCCASIwDQYJKoZIhvcN
AQEBBQADggEPADCCAQoCggEBAM3Cgl0P9Ey69fqcq4QV97/EjW/X0x/JUX01CjU8
du8MdjOq2lfr5k1LXAn12BDOZjYc2f40MC1Y7Di0ZPWgJvky1UJJbuGC6tUlQ9TT
JVtKw5ukJzCPqojDy0Hz3EhpfgM+hBjhR3aDCAHZSCplv1G1G6JM0eqVeb5hvdpU
GOiPNAclkJenGOfFJzAifz9qZT07UyX037/48CuhvPCImM72hUHzYILOqYFXAxCS
wx4NmzPvGtTu1VkKvbsBSVrxJfu9mircsQD3ge9yUSdqKzp7WEl4d5fCW9cr/k7X
f2rX0MLZCpPPAX8EJVC1nkYdgpi6w4THrkMm3iz32ZOjUaECAwEAAaNTMFEwHQYD
VR0OBBYEFOIoE67RqlWnW6Soyb7LTVNHChRyMB8GA1UdIwQYMBaAFOIoE67RqlWn
W6Soyb7LTVNHChRyMA8GA1UdEwEB/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEB
AAHbsr9N/95pmxyXlxp69YwiD8Vdgw4Hj4nLa2dxKSmHmEAz6WRtHI94nKnd1cHH
d5tsHLKLOe2l6/7aCXYvsGenTh7HFCZlBg/bWYTVf+DZtCP8x59NkzOmdFfmwDo7
Ley97lbNZEssA87H9PoDSSnNC9irxpvq9q2oI1qjNoxMCvL2iBbH5/IOvjZj/Ace
2uohH4khyDR7kuRUCmxgnQW4po2D6BjUWfs0FrrCbKk5oFl33ilepLwOmDWbP5RH
5pPEmH14Ug1X3lwBCmaM88/stzU5r/ulofVnBnQYkbvgBU54h5mZUPKHhvZsOqm/
6qP1HFvy0k/V9Zc0fBIoS+k=
-----END CERTIFICATE-----
`

const testAccSourceApiCertKeyPEMUpdated = `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCTrx1fI1V3Lk5j
U7qopPNTOUXpunmSElTQdGk6CpkowwT8oGYX9FtoE2+Cgr+IKCdJZqivtknyArac
YuINgynyT7OtAYCsq4pbZEPSut0TjobrQLC4RpwW4NTQxa8qhuTBNvAvrjkZSBTu
nmsp3mghBAOzX/9H+EePa4eECZsFxhb6UqRHkL1Vd3A+yFc/JkrKNUpGXrn9m+Jc
Vwi18UreNBwD3hqZwByRjkhOBppLhvmFTE+n9KLabHOTrkUxDt0zRtmvxv/iIknb
zbF7WOWNAHCWKZ4w+m4WKTdTmlJPnlMopSL62A/WDLABQEaTcniTeKP/vMNeArp2
Otj/LTHtAgMBAAECggEAEhN1d18WR2OiEZM+fgTOiYmxjnWWCrR2sBsHhCSm8W/J
TPh+9oCgrYDvfZxuQAaWJGxNyk/4Bgih/Es4GupKeI+b0Nmlu7y1TxFbIdBgXdPs
nJ/K4Ys28K/tLG59zak8F5AXXuams3r4oOo/+Jxxx1S79TkTQGZZr+q76APa9x5D
s8DmRxsNrcFsrmRgR0EKBWlLBKGLQXndLWJdRYVsmqsyA+8ppAwIfy+gL7oSjrAn
JMA6Hh2BIu5QDtVtbkECQh6Ts/3Souk+3a34VZDY3svaLn0RQqyzKzayy7FoApaU
R+nsBdbDBXJVd4Kaxzz/c7WmDW3P1LiJ9mkzmrZL4QKBgQDIhUIXafvYzqeZIhe9
JDyDY2YJ6lVVaYlcfxvQDQ0WrrzVAOgDTr2LjL7Of5g2eyhtGgpOv7H+xKB9vtM5
jDoPrS29Np7CrHONHrQOXo4hVHRgbnaBM30JsX+4FHJirHT/DWa6IPhr6PXvQ0SW
RJo9N7/TtoAmKJmhuSupKSnufQKBgQC8i3ssuUjZE1OyKUc3Ow6A32tcs6e9K/LH
gPtfNz0UuyEF91rTa7NanS5YTRcs2W0fa0MPl8TLjJVGjch6mpa2mvWUlI4CbqMk
v+HFVt7DRlvfkhwWT/kxRPWk5JvfvoZpGzB/OwcC0QJMc57nouXef0mBHs+mjvdM
86jKIMp8MQKBgQC8TN9Gq4EUFOQwGmLefDPTUjlVIFgoECYs1TRiCI98+HRdNi+q
Q3wdmvrnYcynhLnu0mBRmRRmYXExBEzIotU3OXJVWAB4YYdVQzrkr4ia5fsvPyxL
ofVOb/+29oSAAuEPTh9ktGd0tlyU5b93uU56JGsgPiZTLxcuVnwYq44z5QKBgQCK
cLwzvaSwx5En+0kwU6u1ktcpTTz4GSH9vUZVxqCR7Uk1MUe+f5FFEq6p5YkDxG5s
qHyoIV89um8NTCQNah8J28RAoceYH5ao8D+df8i3h50EDyuScoxd0e2wcq/FXQci
e/VjJntiZQeUDT17hc/VTI8zQOVlvgn1td+U8vjasQKBgAMg1RM5vjrxVrD9WMB9
R3RgE8pRyPbJ5622zFJza11Sct6BtD+vY3O4MGZhYUzil570HBePWUrn4nrplrM/
6eVUeT/GFgmsl5FPp1cEAHg0Sz68NoAfApu7HkATnisUy0yQ9PQ01aWN6xvO9GSz
7GMvomkZAQ2pxip/xdh/jTx1
-----END PRIVATE KEY-----
-----BEGIN CERTIFICATE-----
MIIDHTCCAgWgAwIBAgIUDL7FwIy1JzRvRy3FYesBG5SY0ZswDQYJKoZIhvcNAQEL
BQAwHjEcMBoGA1UEAwwTdGYtYWNjLXRlc3QtdXBkYXRlZDAeFw0yNjEwMTYwNzUw
NTZaFw0zNjEwMTMwNzUwNTZaMB4xHDAaBgNVBAMME3RmLWFjYy10ZXN0LXVwZGF0
ZWQwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCTrx1fI1V3Lk5jU7qo
pPNTOUXpunmSElTQdGk6CpkowwT8oGYX9FtoE2+Cgr+IKCdJZqivtknyAracYuIN
gynyT7OtAYCsq4pbZEPSut0TjobrQLC4RpwW4NTQxa8qhuTBNvAvrjkZSBTunmsp
3mghBAOzX/9H+EePa4eECZsFxhb6UqRHkL1Vd3A+yFc/JkrKNUpGXrn9m+JcVwi1
8UreNBwD3hqZwByRjkhOBppLhvmFTE+n9KLabHOTrkUxDt0zRtmvxv/iIknbzbF7
WOWNAHCWKZ4w+m4WKTdTmlJPnlMopSL62A/WDLABQEaTcniTeKP/vMNeArp2Otj/
LTHtAgMBAAGjUzBRMB0GA1UdDgQWBBS4sKpHZgcPAXMgB0kZzdb3WzS4XzAfBgNV
HSMEGDAWgBS4sKpHZgcPAXMgB0kZzdb3WzS4XzAPBgNVHRMBAf8EBTADAQH/MA0G
CSqGSIb3DQEBCwUAA4IBAQBs8J2v72qVT/V37WRvfCP2Fi9z50c+WIDTtIQY4AXw
D6BJzdPQC6wkBQ7l8TduMtsK0m/GCiEK71FkPAdvmewYV1dPNZwvomt7+q0NjwaK
ff8UXTLCvgo3+vSERqs/8IMgPpXkjudtnPvwRKBvBvBVWV1PcvkwhzZ/6uflzTy+
+vNwIY9EueajIYeLsUEVy7etsNuaiB9O1vbPBA2ZZc9Pq5wPJEM9nBTgJi2+orcq
iof6iugolWPASdVe4lI/tvM5HidCIdtf1RvtaJhoHZ8whIMn/z5rq0q3KXs6hDFg
/N8aii5n3bPla9rVBy8kTbgCSyJTCsysEvXaH1gZXnAl
-----END CERTIFICATE-----
`

func TestAccSourceApiResource_basic(t *testing.T) {
	var resourceName = "bloxone_cdc_source_nios.test"
	var v cdc.SourceNios
	name := acctest.RandomNameWithPrefix("source")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSourceApiBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSourceApiResource_disappears(t *testing.T) {
	resourceName := "bloxone_cdc_source_nios.test"
	var v cdc.SourceNios
	name := acctest.RandomNameWithPrefix("source")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckSourceApiDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccSourceApiBasicConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					testAccCheckSourceApiDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccSourceApiResource_Address(t *testing.T) {
	var resourceName = "bloxone_cdc_source_nios.test_address"
	var v cdc.SourceNios
	name := acctest.RandomNameWithPrefix("source")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSourceApiAddress(name, name+"-1.example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", name+"-1.example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccSourceApiAddress(name, name+"-2.example.com"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "address", name+"-2.example.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSourceApiResource_CaCertificate(t *testing.T) {
	var resourceName = "bloxone_cdc_source_nios.test_ca_certificate"
	var v cdc.SourceNios
	name := acctest.RandomNameWithPrefix("source")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSourceApiCaCertificate(name, testAccSourceApiCertPEM),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ca_certificate", testAccSourceApiCertPEM),
				),
			},
			// Update and Read
			{
				Config: testAccSourceApiCaCertificate(name, testAccSourceApiCertPEMUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ca_certificate", testAccSourceApiCertPEMUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSourceApiResource_Certificate(t *testing.T) {
	var resourceName = "bloxone_cdc_source_nios.test_certificate"
	var v cdc.SourceNios
	name := acctest.RandomNameWithPrefix("source")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSourceApiCertificate(name, testAccSourceApiCertKeyPEM),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "certificate", testAccSourceApiCertKeyPEM),
				),
			},
			// Update and Read
			{
				Config: testAccSourceApiCertificate(name, testAccSourceApiCertKeyPEMUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "certificate", testAccSourceApiCertKeyPEMUpdated),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSourceApiResource_Description(t *testing.T) {
	var resourceName = "bloxone_cdc_source_nios.test_description"
	var v cdc.SourceNios
	name := acctest.RandomNameWithPrefix("source")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSourceApiDescription(name, "DESCRIPTION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "DESCRIPTION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccSourceApiDescription(name, "DESCRIPTION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "description", "DESCRIPTION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSourceApiResource_Enabled(t *testing.T) {
	var resourceName = "bloxone_cdc_source_nios.test_enabled"
	var v cdc.SourceNios
	name := acctest.RandomNameWithPrefix("source")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSourceApiEnabled(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccSourceApiEnabled(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSourceApiResource_InsecureMode(t *testing.T) {
	var resourceName = "bloxone_cdc_source_nios.test_insecure_mode"
	var v cdc.SourceNios
	name := acctest.RandomNameWithPrefix("source")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSourceApiInsecureMode(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "insecure_mode", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccSourceApiInsecureMode(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "insecure_mode", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSourceApiResource_Name(t *testing.T) {
	var resourceName = "bloxone_cdc_source_nios.test_name"
	var v cdc.SourceNios
	name := acctest.RandomNameWithPrefix("source")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSourceApiName(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
				),
			},
			// Update and Read
			{
				Config: testAccSourceApiName(name + "-updated"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name+"-updated"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSourceApiResource_RpzLogInsecureMode(t *testing.T) {
	var resourceName = "bloxone_cdc_source_nios.test_rpz_log_insecure_mode"
	var v cdc.SourceNios
	name := acctest.RandomNameWithPrefix("source")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSourceApiRpzLogInsecureMode(name, "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_log_insecure_mode", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccSourceApiRpzLogInsecureMode(name, "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "rpz_log_insecure_mode", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccSourceApiResource_Tags(t *testing.T) {
	var resourceName = "bloxone_cdc_source_nios.test_tags"
	var v cdc.SourceNios
	name := acctest.RandomNameWithPrefix("source")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccSourceApiTags(name, map[string]string{"tag1": "value1"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
				),
			},
			// Update and Read
			{
				Config: testAccSourceApiTags(name, map[string]string{"tag1": "value1changed", "tag2": "value2"}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSourceApiExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckSourceApiExists(ctx context.Context, resourceName string, v *cdc.SourceNios) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		id, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}
		apiRes, _, err := acctest.BloxOneClient.CDCAPI.
			SourceApiAPI.
			ReadSourceNios(ctx, id).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.HasResults() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetResults()
		return nil
	}
}

func testAccCheckSourceApiDestroy(ctx context.Context, v *cdc.SourceNios) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		if v.Id == nil {
			return nil
		}
		_, httpRes, err := acctest.BloxOneClient.CDCAPI.
			SourceApiAPI.
			ReadSourceNios(ctx, *v.Id).
			Execute()
		if err != nil {
			if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
				// resource was deleted
				return nil
			}
			return err
		}
		return errors.New("expected to be deleted")
	}
}

func testAccCheckSourceApiDisappears(ctx context.Context, v *cdc.SourceNios) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.CDCAPI.
			SourceApiAPI.
			DeleteSourceNios(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccSourceApiBasicConfig(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_source_nios" "test" {
    name          = %q
    address       = "%s.example.com"
    nios_username = "tf_auto"
    nios_password = "tf_auto_password1!"
}
`, name, name)
}

func testAccSourceApiAddress(name string, address string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_source_nios" "test_address" {
    name          = %q
    address       = %q
    nios_username = "tf_auto"
    nios_password = "tf_auto_password1!"
}
`, name, address)
}

func testAccSourceApiCaCertificate(name string, caCertificate string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_source_nios" "test_ca_certificate" {
    name           = %q
    address        = "%s.example.com"
    nios_username  = "tf_auto"
    nios_password  = "tf_auto_password1!"
    ca_certificate = %q
}
`, name, name, caCertificate)
}

func testAccSourceApiCertificate(name string, certificate string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_source_nios" "test_certificate" {
    name          = %q
    address       = "%s.example.com"
    nios_username = "tf_auto"
    nios_password = "tf_auto_password1!"
    certificate   = %q
}
`, name, name, certificate)
}

func testAccSourceApiDescription(name string, description string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_source_nios" "test_description" {
    name          = %q
    address       = "%s.example.com"
    nios_username = "tf_auto"
    nios_password = "tf_auto_password1!"
    description   = %q
}
`, name, name, description)
}

func testAccSourceApiEnabled(name string, enabled string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_source_nios" "test_enabled" {
    name          = %q
    address       = "%s.example.com"
    nios_username = "tf_auto"
    nios_password = "tf_auto_password1!"
    enabled       = %s
}
`, name, name, enabled)
}

func testAccSourceApiInsecureMode(name string, insecureMode string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_source_nios" "test_insecure_mode" {
    name           = %q
    address        = "%s.example.com"
    nios_username  = "tf_auto"
    nios_password  = "tf_auto_password1!"
    ca_certificate = %q
    insecure_mode  = %s
}
`, name, name, testAccSourceApiCertPEM, insecureMode)
}

func testAccSourceApiName(name string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_source_nios" "test_name" {
    name          = %q
    address       = "%s.example.com"
    nios_username = "tf_auto"
    nios_password = "tf_auto_password1!"
}
`, name, name)
}

func testAccSourceApiRpzLogInsecureMode(name string, rpzLogInsecureMode string) string {
	return fmt.Sprintf(`
resource "bloxone_cdc_source_nios" "test_rpz_log_insecure_mode" {
    name                   = %q
    address                = "%s.example.com"
    nios_username          = "tf_auto"
    nios_password          = "tf_auto_password1!"
    rpz_log_insecure_mode  = %s
}
`, name, name, rpzLogInsecureMode)
}

func testAccSourceApiTags(name string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf("\t\t%s = %q\n", k, v)
	}
	tagsStr += "\t}"

	return fmt.Sprintf(`
resource "bloxone_cdc_source_nios" "test_tags" {
    name          = %q
    address       = "%s.example.com"
    nios_username = "tf_auto"
    nios_password = "tf_auto_password1!"
    tags          = %s
}
`, name, name, tagsStr)
}
