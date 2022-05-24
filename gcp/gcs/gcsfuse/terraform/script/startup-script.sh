#!/bin/bash

set -eu

# メタデータサーバーから、ファイルを取得する。
curl -X GET --location "http://metadata.google.internal/computeMetadata/v1/instance/attributes/init-script" -H "Metadata-Flavor: Google" > /root/init.sh
curl -X GET --location "http://metadata.google.internal/computeMetadata/v1/instance/attributes/es-config" -H "Metadata-Flavor: Google" > /usr/share/elasticsearch/config/elasticsearch.yml

sudo mount -o discard,defaults /dev/sdb /usr/share/elasticsearch/data
sudo -u elasticsearch bash -c '
gcsfuse stg-gcsfuse-test1 /usr/share/elasticsearch/config/user_dictionary
/usr/share/elasticsearch/bin/elasticsearch
'
