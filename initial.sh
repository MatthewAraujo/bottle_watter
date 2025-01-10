#!/bin/bash

os=$(uname -s)
arch=$(uname -m)

if [ "$os" == "Linux" ]; then
  BINAR_NAME="bottle_watter-linux-amd64"
elif [ "$os" == "Darwin" ]; then
  BINAR_NAME="bottle_watter-darwin-amd64"
elif [ "$os" == "Cygwin" ] || [ "$os" == "mingw" ]; then
  BINAR_NAME="bottle_watter.exe"
else
  echo "OS not supported"
  exit 1
fi

echo "Downloading the binary..."
curl -L -o "/usr/local/bin/bottle_watter" "https://github.com/matthewaraujo/bottle_watter/releases/download/Binary/$BINAR_NAME"

chmod +x /usr/local/bin/bottle_watter

if command -v bottle_watter &> /dev/null
then
  echo "Installation successful! You can now use the 'bottle_watter' command."
else
  echo "Error installing the binary."
fi

bottle_watter --help