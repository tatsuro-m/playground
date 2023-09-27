resource "aws_lb" "main" {
  name               = "main"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.main.id]
  subnets            = [
    aws_subnet.public_a.id,
    aws_subnet.public_c.id
  ]

  enable_deletion_protection = false

  tags = {
    Main = "true"
  }
}

# public サブネットに存在する、httpd 起動中のインスタンスに対するグループ
resource "aws_lb_target_group" "web" {
  name     = "web"
  port     = 80
  protocol = "HTTP"

  health_check {
    enabled  = true
    timeout  = 3
    interval = 5
  }

  vpc_id = aws_vpc.main.id
}

resource "aws_lb_listener" "web" {
  load_balancer_arn = aws_lb.main.id
  port              = "80"
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.web.arn
  }
}

resource "aws_lb_target_group_attachment" "web" {
  target_group_arn = aws_lb_target_group.web.arn
  target_id        = aws_instance.web.id
  port             = 80
}

# private サブネットに存在する、nginx 起動中のインスタンスに対するグループ
resource "aws_lb_target_group" "internal" {
  name     = "internal"
  port     = 80
  protocol = "HTTP"

  health_check {
    enabled  = true
    timeout  = 3
    interval = 5
  }

  vpc_id = aws_vpc.main.id
}

resource "aws_lb_listener" "internal" {
  load_balancer_arn = aws_lb.main.id
  # リスナーは１つの ALB に関連付けている以上、「プロトコル x ポート」で unique にならないといけないので、ポートを意図的に変える
  port              = 8080
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.internal.arn
  }
}

resource "aws_lb_target_group_attachment" "internal" {
  target_group_arn = aws_lb_target_group.internal.arn
  target_id        = aws_instance.internal.id
  port             = 80
}
