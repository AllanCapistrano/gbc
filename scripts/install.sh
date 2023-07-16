#!/bin/bash

# Downloading the necessary files.
curl -O https://github.com/AllanCapistrano/gbc/releases/download/v1.0.0/gbc
curl -O https://raw.githubusercontent.com/AllanCapistrano/gbc/main/config/gbc.conf

# Creating the config directory for gbc.
mkdir -p $HOME/.config/gbc

# Moving the CLI and settings file.
sudo mv gbc /usr/bin
mv gbc.conf $HOME/.config/gbc

# Giving permission to run the CLI.
sudo chmod +x /usr/bin/gbc

gbc --help