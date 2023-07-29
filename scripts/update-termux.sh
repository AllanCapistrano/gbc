#!/bin/bash

# Removing older version.
sudo rm /usr/bin/gbc

# Downloading the latest version.
curl -LO https://github.com/AllanCapistrano/gbc/releases/download/v1.2.0/gbc-arm

# Renaming file
mv gbc-arm gbc

# Moving the CLI and settings file.
mv gbc $PATH

# Giving permission to run the CLI.
chmod +x $PATH/gbc

gbc --help

echo "gbc updated to version 1.2.0"