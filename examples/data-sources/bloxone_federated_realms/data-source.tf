# Get Federated Realm filtered by an attribute
data "bloxone_federated_realms" "example_by_attribute" {
  filters = {
    name = "example_federated_realm"
  }
}

#Get Federated Realm filtered by tag
data "bloxone_federated_realms" "example_by_tag" {
  tag_filters = {
    key1 = "value1"
  }
}

#Get all Federated Realm
data "bloxone_federated_realms" "example_all" {}
