#!/bin/bash

# このスクリプトは、 root ユーザとして SSH 接続して実行する。
set -eu

useradd -m elasticsearch
passwd -u -f elasticsearch

# gcs fuse の設定（https://github.com/GoogleCloudPlatform/gcsfuse/blob/master/docs/installing.md）
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

# elasticsearch の設定
sudo yum install java-1.8.0-openjdk -y
curl -L -o /root/elasticsearch-5.3.2.tar.gz https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-5.3.2.tar.gz
tar -xvf /root/elasticsearch-5.3.2.tar.gz
mv /root/elasticsearch-5.3.2/ /root/elasticsearch/
mv /root/elasticsearch/ /usr/share/
/usr/share/elasticsearch/bin/elasticsearch-plugin install analysis-kuromoji
/usr/share/elasticsearch/bin/elasticsearch-plugin install analysis-icu
yes | /usr/share/elasticsearch/bin/elasticsearch-plugin install repository-gcs
mkdir -p /usr/share/elasticsearch/config/user_dictionary

chmod 777 -R /usr/share/elasticsearch

# 外部永続ディスクのマウント設定
sudo lsblk # デバイス名を表示するが、今回なら「sdb」。
yes | sudo mkfs.ext4 -m 0 -E lazy_itable_init=0,lazy_journal_init=0,discard /dev/sdb
sudo mkdir -p /usr/share/elasticsearch/data
sudo mount -o discard,defaults /dev/sdb /usr/share/elasticsearch/data
sudo chmod a+w /usr/share/elasticsearch/data
