#!/bin/bash

# gcs fuse の設定
sudo tee /etc/yum.repos.d/gcsfuse.repo > /dev/null <<EOF
[gcsfuse]
name=gcsfuse (packages.cloud.google.com)
baseurl=https://packages.cloud.google.com/yum/repos/gcsfuse-el7-x86_64
enabled=1
gpgcheck=1
repo_gpgcheck=0
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg
       https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
EOF

sudo yum install gcsfuse -y

mkdir -p /root/gcsfuse/1
gcsfuse stg-gcsfuse-test1 /root/gcsfuse/1

mkdir -p /root/gcsfuse/2
gcsfuse stg-gcsfuse-test2 /root/gcsfuse/2
