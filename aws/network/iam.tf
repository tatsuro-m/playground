resource "aws_iam_role" "handson_ssm" {
  name = "handson-ssm"

  assume_role_policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "sts:AssumeRole"
            ],
            "Principal": {
                "Service": [
                    "ec2.amazonaws.com"
                ]
            }
        }
    ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "handson_ssm" {
  role       = aws_iam_role.handson_ssm.id
  # AWS 管理のポリシーを直接アタッチする
  policy_arn = "arn:aws:iam::aws:policy/AmazonSSMFullAccess"
}
