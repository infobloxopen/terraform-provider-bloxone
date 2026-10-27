resource "bloxone_cdc_application" "example" {
  name = "example_cdc_application"

  data = {
    scripts = [{
      filename = "script_1.py"
      content  = "print(\"hello\")"
    }]
  }
}

resource "bloxone_cdc_application" "example_with_optional_fields" {
  name        = "example_cdc_application_optional_fields"
  description = "Example CDC application"
  enabled     = false

  data = {
    scripts = [{
      filename = "script_1.py"
      content  = "print(\"hello\")"
    }]
  }
  variables = {
    var1 = "value1"
  }

  tags = {
    site = "Site A"
  }
}

resource "bloxone_cdc_application" "example_with_script_file" {
  name = "example_cdc_application_script_file"

  data = {
    scripts = [{
      filename = "cdc_application_handler_sample.py"
      # Load the script content from a file next to this configuration
      content = file("${path.module}/cdc_application_handler_sample.py")
    }]
  }
}
