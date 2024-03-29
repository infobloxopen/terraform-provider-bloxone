---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "bloxone_keys_tsig Resource - terraform-provider-bloxone"
subcategory: "Keys"
description: |-
  Manages a TSIG Key.
---

# bloxone_keys_tsig (Resource)

Manages a TSIG Key.

## Example Usage

```terraform
resource "bloxone_keys_tsig" "example" {
  name = "example_tsig_key"

  # Other optional fields
  comment   = "key created through terraform"
  algorithm = "hmac_sha256"
  secret    = "your-base64-encoded-secret-here"
  tags = {
    location = "site1"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The TSIG key name in the absolute domain name format.

### Optional

- `algorithm` (String) TSIG key algorithm.

  Valid values are:
  * _hmac_sha1_
  * _hmac_sha224_
  * _hmac_sha256_
  * _hmac_sha384_
  * _hmac_sha512_

  Defaults to _hmac_sha256_.
- `comment` (String) The description for the TSIG key. May contain 0 to 1024 characters. Can include UTF-8.
- `secret` (String, Sensitive) The TSIG key secret as a Base64 encoded string.
- `tags` (Map of String) The tags for the TSIG key in JSON format.

### Read-Only

- `created_at` (String) Time when the object has been created.
- `id` (String) The resource identifier.
- `protocol_name` (String) The TSIG key name supplied during a create/update operation that is converted to canonical form in punycode.
- `updated_at` (String) Time when the object has been updated. Equals to _created_at_ if not updated after creation.
