#!/bin/bash

# Installation script for GitHub-Sync on Linux

# Colors for better output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${GREEN}Starting GitHub-Sync installation...${NC}"

# Get Git commit hash
GIT_COMMIT=$(git rev-parse --short HEAD 2>/dev/null)
if [ -z "$GIT_COMMIT" ]; then
    GIT_COMMIT="unknown"
fi

# Get current date
BUILD_DATE=$(date +%Y-%m-%d)

# Build the binary with version information
echo -e "Building GitHub-Sync..."
if ! go build -o ghs -ldflags "-X 'github.com/mostafa-mahmood/GitHub-Sync/cmd.Version=1.0.0' -X 'github.com/mostafa-mahmood/GitHub-Sync/cmd.BuildDate=$BUILD_DATE' -X 'github.com/mostafa-mahmood/GitHub-Sync/cmd.GitCommit=$GIT_COMMIT'"; then
    echo -e "${RED}❌ Failed to build GitHub-Sync. Please ensure Go is installed and your project is set up correctly.${NC}"
    exit 1
fi

# Determine installation directory
if [ -w "/usr/local/bin" ]; then
    # User has write permissions to /usr/local/bin
    INSTALL_DIR="/usr/local/bin"
    SUDO=""
elif [ -d "$HOME/.local/bin" ] || mkdir -p "$HOME/.local/bin"; then
    # User has a local bin directory or we can create one
    INSTALL_DIR="$HOME/.local/bin"
    SUDO=""
    
    # Check if ~/.local/bin is in PATH
    if ! echo "$PATH" | grep -q "$HOME/.local/bin"; then
        echo -e "${YELLOW}Adding $HOME/.local/bin to your PATH...${NC}"
        echo 'export PATH="$HOME/.local/bin:$PATH"' >> "$HOME/.bashrc"
        
        # If user is using zsh, update .zshrc as well
        if [ -f "$HOME/.zshrc" ]; then
            echo 'export PATH="$HOME/.local/bin:$PATH"' >> "$HOME/.zshrc"
        fi
    fi
else
    # Fall back to /usr/local/bin with sudo
    INSTALL_DIR="/usr/local/bin"
    SUDO="sudo"
fi

# Check for sudo permissions if needed
if [ -n "$SUDO" ]; then
    if ! $SUDO -v; then
        echo -e "${RED}❌ Sudo access is required but not granted.${NC}"
        exit 1
    fi
fi

# Install the binary
echo -e "Installing GitHub-Sync to ${INSTALL_DIR}..."
$SUDO mv ghs "$INSTALL_DIR/"
if [ $? -ne 0 ]; then
    echo -e "${RED}❌ Failed to move ghs to $INSTALL_DIR.${NC}"
    exit 1
fi

$SUDO chmod +x "$INSTALL_DIR/ghs"
if [ $? -ne 0 ]; then
    echo -e "${RED}❌ Failed to set executable permissions on ghs.${NC}"
    exit 1
fi

# Update PATH for the current session
export PATH="$INSTALL_DIR:$PATH"

echo -e "${GREEN}Installation complete!${NC}"
echo -e "You can now use GitHub-Sync by running: ${YELLOW}ghs${NC}"

if [[ "$INSTALL_DIR" == "$HOME/.local/bin" && ! "$PATH" == *"$HOME/.local/bin"* ]]; then
    echo -e "${YELLOW}Please restart your terminal or run:${NC}"
    echo -e "  source ~/.bashrc"
    echo -e "to update your PATH."
fi