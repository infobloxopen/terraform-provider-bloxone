resource "bloxone_cdc_application" "example" {
  name = "example_cdc_application"

  # data has no fixed schema, so it is passed as JSON; the value returned by
  # the backend is available in file_data
  data = jsonencode({
    scripts = [{
      filename = "script_1.py"
      content  = "print(\"automation\")"
    }]
  })
}

resource "bloxone_cdc_application" "example_with_optional_fields" {
  name        = "example_cdc_application_optional_fields"
  description = "Example CDC application"
  enabled     = false

  data = jsonencode({
    scripts = [{
      filename = "script_1.py"
      content  = "print(\"automation\")"
    }]
  })
  variables = {
    var1 = "value1"
  }

  tags = {
    site = "Site A"
  }
}

resource "bloxone_cdc_application" "example_with_script_file" {
  name = "example_cdc_application_script_file"

  data = jsonencode({
    scripts = [{
      # Replace with the path to your own script
      filename = "handler.py"
      content  = file("${path.module}/scripts/handler.py")
    }]
  })
}
