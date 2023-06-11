#!/bin/bash

set -e

# このコマンドを GCE インスタンスに ssh で接続してから実行する

sudo apt-get update
sudo apt-get install -y git curl gcc make zlib1g-dev libbz2-dev libreadline-dev libssl-dev libsqlite3-dev

curl https://pyenv.run | bash

echo 'export PATH="$HOME/.pyenv/bin:$PATH"' >> ~/.bashrc
echo 'eval "$(pyenv init --path)"' >> ~/.bashrc
echo 'eval "$(pyenv init -)"' >> ~/.bashrc
source ~/.bashrc

pyenv install 3.10.6
pyenv global 3.10.6

git clone https://github.com/AUTOMATIC1111/stable-diffusion-webui
./stable-diffusion-webui/webui.sh
