resource "bloxone_federated_realm" "test_name" {
  name = "test_name_federated_realm"
  tags = {
    key1 = "value1"
    key2 = "value2"
  }
}