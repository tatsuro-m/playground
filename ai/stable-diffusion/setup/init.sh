#!/bin/bash

set -e

# このコマンドを GCE インスタンスに ssh で接続してから実行する

sudo apt-get update
sudo apt-get install -y git curl gcc make zlib1g-dev libbz2-dev libreadline-dev libssl-dev libsqlite3-dev liblzma-dev libffi-dev

curl https://pyenv.run | bash

echo 'export PATH="$HOME/.pyenv/bin:$PATH"' >> ~/.bashrc
echo 'eval "$(pyenv init --path)"' >> ~/.bashrc
echo 'eval "$(pyenv init -)"' >> ~/.bashrc
source ~/.bashrc

pyenv install 3.10.6
pyenv global 3.10.6

git clone https://github.com/AUTOMATIC1111/stable-diffusion-webui

# 適当なモデルをダウンロード
wget -P ./stable-diffusion-webui/models/Stable-diffusion  https://huggingface.co/sazyou-roukaku/BracingEvoMix/resolve/main/BracingEvoMix_v1.safetensors
wget -O ./stable-diffusion-webui/models/Stable-diffusion/majicmixRealistic_v6.safetensors https://civitai.com/api/download/models/94640


./stable-diffusion-webui/webui.sh
