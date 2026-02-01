#!/bin/bash
set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

REPO="mofax/pkbin"
BINARY_NAME="pk"
INSTALL_DIR="${INSTALL_DIR:-$HOME/.local/bin}"

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$OS" in
  linux)
    OS_NAME="linux"
    ;;
  darwin)
    OS_NAME="darwin"
    ;;
  mingw*|msys*)
    OS_NAME="windows"
    BINARY_NAME="pk.exe"
    ;;
  *)
    echo -e "${RED}Unsupported OS: $OS${NC}"
    exit 1
    ;;
esac

case "$ARCH" in
  x86_64)
    ARCH_NAME="amd64"
    ;;
  aarch64|arm64)
    ARCH_NAME="arm64"
    ;;
  *)
    echo -e "${RED}Unsupported architecture: $ARCH${NC}"
    exit 1
    ;;
esac

# Fetch latest release
echo "Fetching latest release for $OS_NAME/$ARCH_NAME..."
DOWNLOAD_URL=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | \
  grep "browser_download_url" | \
  grep "$OS_NAME" | \
  grep "$ARCH_NAME" | \
  head -1 | \
  cut -d'"' -f4)

if [ -z "$DOWNLOAD_URL" ]; then
  echo -e "${RED}Failed to find release for $OS_NAME/$ARCH_NAME${NC}"
  exit 1
fi

echo "Downloading from: $DOWNLOAD_URL"

# Create install directory if it doesn't exist
mkdir -p "$INSTALL_DIR"

# Download and install
TEMP_FILE=$(mktemp)
trap "rm -f $TEMP_FILE" EXIT

curl -sL "$DOWNLOAD_URL" -o "$TEMP_FILE"
chmod +x "$TEMP_FILE"
mv "$TEMP_FILE" "$INSTALL_DIR/$BINARY_NAME"

echo -e "${GREEN}Successfully installed $BINARY_NAME to $INSTALL_DIR/$BINARY_NAME${NC}"
echo ""

# Check if install directory is in PATH
if [[ ":$PATH:" == *":$INSTALL_DIR:"* ]]; then
  echo -e "${GREEN}finished"
else
  echo -e "${RED}Warning: $INSTALL_DIR is not in your PATH${NC}"
  echo "Add the following line to your shell profile (~/.bashrc, ~/.zshrc, etc.):"
  echo "export PATH=\"\$HOME/.local/bin:\$PATH\""
fi

echo ""
echo "To get started, run: pk --help"
