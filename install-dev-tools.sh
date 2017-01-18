#!/bin/bash

set -e

if [ -z $(which glide) ]; then
  sudo add-apt-repository ppa:masterminds/glide && sudo apt-get update
  sudo apt-get install glide
else
  sudo apt-get update
fi

# Install Node.js
if [ -z $(which node) ]; then
  curl -sL https://deb.nodesource.com/setup | sudo bash -
  sudo apt-get install -y nodejs
fi

# Install NPM
if [ -z $(which npm) ]; then
  sudo npm install npm -g
fi

# Install Webpack
if [ -z $(which webpack) ]; then
  sudo npm install webpack -g
fi
