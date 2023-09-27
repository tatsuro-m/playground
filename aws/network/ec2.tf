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
chmod 777 /var/www/html -R
echo '<!DOCTYPE html>
            <html>
            <head>
                <meta charset="UTF-8">
                <title>EC2(public subnet) Web httpd instance(No.1)</title>
            </head>
            <body>
                <h1>EC2(public subnet) Web httpd instance(No.1)</h1>
            </body>
            </html>' > /var/www/html/index.html
systemctl restart httpd.service
EOF

  iam_instance_profile = aws_iam_instance_profile.main.id
}

resource "aws_iam_instance_profile" "main" {
  name = "main"
  role = aws_iam_role.handson_ssm.id
}

resource "aws_instance" "internal" {
  ami           = "ami-0f89bdd365c3d966d"
  instance_type = "t2.micro"
  subnet_id     = aws_subnet.private_c.id

  tags = {
    Name = "internal"
  }

  associate_public_ip_address = false
  vpc_security_group_ids      = [
    aws_security_group.main.id
  ]

  user_data = <<EOF
#!/bin/bash

amazon-linux-extras install nginx1
systemctl enable nginx
systemctl start nginx

chmod 777 /usr/share/nginx/html -R
echo '<!DOCTYPE html>
<html>
<head>
　  <meta charset="UTF-8">
    <title>EC2(private subnet) Web nginx instance(No.1)</title>
</head>
<body>
    <h1>EC2(private subnet) Web nginx instance(No.1)</h1>
</body>
</html>' > /usr/share/nginx/html/index.html

systemctl restart nginx
EOF

  iam_instance_profile = aws_iam_instance_profile.main.id
}
