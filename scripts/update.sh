#!/bin/bash

# Removing older version.
sudo rm /usr/bin/gbc

# Downloading the latest version.
curl -LO https://github.com/AllanCapistrano/gbc/releases/download/v1.2.0/gbc

# Moving the CLI and settings file.
sudo mv gbc /usr/bin

# Giving permission to run the CLI.
sudo chmod +x /usr/bin/gbc

gbc --help

echo "gbc updated to version 1.2.0"