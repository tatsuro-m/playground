#!/bin/bash

# メタデータサーバーから、初回のみ必要なスクリプトを書き込む。
curl -X GET --location "http://metadata.google.internal/computeMetadata/v1/instance/attributes/init-script" -H "Metadata-Flavor: Google" > /root/init.sh

sudo mount -o discard,defaults /dev/sdb /usr/share/elasticsearch/data
sudo -u elasticsearch bash -c '
gcsfuse stg-gcsfuse-test1 /usr/share/elasticsearch/config/user_dictionary
/usr/share/elasticsearch/bin/elasticsearch
'
