#!/bin/bash

os=$(uname -s)
arch=$(uname -m)

if [ "$os" == "Linux" ]; then
  BINAR_NAME="watter_bottle-linux-amd64"
elif [ "$os" == "Darwin" ]; then
  BINAR_NAME="watter_bottle-darwin-amd64"
elif [ "$os" == "Cygwin" ] || [ "$os" == "mingw" ]; then
  BINAR_NAME="watter_bottle.exe"
else
  echo "OS not supported"
  exit 1
fi

echo "Downloading the binary..."
curl -L -o "/usr/local/bin/watter_bottle" "https://github.com/matthewaraujo/watter_bottle/releases/download/Binary/$BINAR_NAME"

chmod +x /usr/local/bin/watter_bottle

if command -v watter_bottle &> /dev/null
then
  echo "Installation successful! You can now use the 'watter_bottle' command."
else
  echo "Error installing the binary."
fi

watter_bottle --help