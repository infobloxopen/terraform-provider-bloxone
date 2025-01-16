resource "bloxone_ipam_ip_space" "example" {
  name = "example_ip_space"
}

resource "bloxone_ipam_address_block" "example" {
  address = "10.0.0.0"
  cidr    = 8
  space   = bloxone_ipam_ip_space.example.id
}

resource "bloxone_ipam_subnet" "example" {
  address = "10.1.0.0"
  cidr    = 24
  space   = bloxone_ipam_ip_space.example.id
}

resource "bloxone_ipam_range" "example" {
  start      = "10.1.0.10"
  end        = "10.1.0.20"
  space      = bloxone_ipam_ip_space.example.id
  depends_on = [bloxone_ipam_subnet.example]
}


resource "bloxone_ipam_address" "example" {
  address = "10.1.0.5"
  space   = bloxone_ipam_ip_space.example.id

  # Other optional fields
  comment = "reservation for Site A"
  names = [
    {
      name = "bby-1"
      type = "user"
    }
  ]
  tags = {
    site = "Site A"
  }

  depends_on = [bloxone_ipam_subnet.example]
}

# Next available address in subnet
resource "bloxone_ipam_address" "example_na_s" {
  next_available_id = bloxone_ipam_subnet.example.id
  space             = bloxone_ipam_ip_space.example.id

  # Other optional fields
  comment = "reservation for Site A"
  names = [
    {
      name = "bby-1"
      type = "user"
    }
  ]
  tags = {
    site = "Site A"
  }
}

# Next available address in address block
resource "bloxone_ipam_address" "example_na_ab" {
  next_available_id = bloxone_ipam_address_block.example.id
  space             = bloxone_ipam_ip_space.example.id

  # Other optional fields
  comment = "reservation for Site A"
  names = [
    {
      name = "bby-1"
      type = "user"
    }
  ]
  tags = {
    site = "Site A"
  }
}

# Next available address in range
resource "bloxone_ipam_address" "example_na_rng" {
  next_available_id = bloxone_ipam_range.example.id
  space             = bloxone_ipam_ip_space.example.id

  # Other optional fields
  comment = "reservation for Site A"
  names = [
    {
      name = "bby-1"
      type = "user"
    }
  ]
  tags = {
    site = "Site A"
  }
}
