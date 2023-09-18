resource "aws_instance" "web" {
  ami           = "ami-0f89bdd365c3d966d"
  instance_type = "t2.micro"
  subnet_id     = aws_subnet.public_a.id

  tags = {
    Name = "Web"
  }

  associate_public_ip_address = true
  vpc_security_group_ids      = [
    aws_security_group.main.id
  ]

  user_data = <<EOF
#!/bin/bash

yum -y update
yum -y install httpd
systemctl enable httpd.service
systemctl start httpd.service
EOF

  iam_instance_profile = aws_iam_instance_profile.main.id
}

resource "aws_iam_instance_profile" "main" {
  name = "main"
  role = aws_iam_role.handson_ssm.id
}
