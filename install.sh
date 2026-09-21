#!/bin/bash
set -e

OS=$(uname -s)

cd platform || exit 1

if [ "$OS" = "Darwin" ]; then
    echo "Detected macOS"

    GOOS=darwin GOARCH=$(uname -m | sed 's/arm64/arm64/;s/x86_64/amd64/') go build -o platform

    mv platform /opt/homebrew/bin/platform

elif [ "$OS" = "Linux" ]; then
    echo "Detected Linux"

    GOOS=linux GOARCH=$(uname -m | sed 's/x86_64/amd64/;s/aarch64/arm64/') go build -o platform

    mv platform /usr/local/bin/platform

else
    echo "Unsupported operating system: $OS"
    exit 1
fi

echo "Platform CLI installed successfully."
