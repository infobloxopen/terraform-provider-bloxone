# The automation script (application) that the destination sends data to (Required as parent)
resource "bloxone_cdc_application" "example" {
  name = "example_cdc_application_for_destination"

  data = jsonencode({
    scripts = [{
      filename = "script_1.py"
      content  = "print(\"example automation script\")"
    }]
  })
}

resource "bloxone_cdc_destination_application" "example" {
  name        = "example_cdc_destination_application"
  application = bloxone_cdc_application.example.id
}

resource "bloxone_cdc_destination_application" "example_with_optional_fields" {
  name               = "example_cdc_destination_application_optional_fields"
  application        = bloxone_cdc_application.example.id
  description        = "Example CDC application destination with optional fields"
  enabled            = false
  output_data_format = "STREAM"

  # Variables passed to the application
  context = {
    var1 = "value1"
  }

  tags = {
    site = "Site A"
  }
}
