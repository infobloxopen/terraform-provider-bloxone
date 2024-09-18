# Get Federated Realm filtered by an attribute
data "bloxone_federation_federated_realms" "example_by_attribute" {
  filters = {
    name = "example_federation_federated_realm"
  }
}

#Get Federated Realm filtered by tag
data "bloxone_federation_federated_realms" "example_by_tag" {
  tag_filters = {
    site = "Site A"
  }
}

#Get all Federated Realm
data "bloxone_federation_federated_realms" "example_all" {}
