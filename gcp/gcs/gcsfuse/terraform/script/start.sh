#!/bin/bash

sudo -u demo-user bash -c 'WHO_AM_I=$(whoami);
echo WHO_AM_I; $WHO_AM_I &>> debug.txt;'

sudo su elasticsearch

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

mkdir -p /home/elasticsearch/gcsfuse/test1
gcsfuse stg-gcsfuse-test1 /home/elasticsearch/gcsfuse/test1

# elasticsearch の設定
sudo yum install java-1.8.0-openjdk -y
cd /root || exit
curl -L -O https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-5.3.3.tar.gz
tar -xvf elasticsearch-5.3.3.tar.gz
cd elasticsearch-5.3.3/bin || exit

./elasticsearch
