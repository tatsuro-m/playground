terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }

  backend "s3" {
    bucket = "tf-state-81931"
    key    = "tf-tutorial/terraform.tfstate"
    region = "ap-northeast-1"
  }
}

# Configure the AWS Provider
provider "aws" {
  region = "ap-northeast-1"
}

resource "aws_s3_bucket" "example" {
  bucket = "my-tf-test-bucket-901q7r0278802"

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}
