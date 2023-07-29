#!/bin/bash

# Downloading the necessary files.
curl -LO https://github.com/AllanCapistrano/gbc/releases/download/v1.1.2/gbc
curl -O https://raw.githubusercontent.com/AllanCapistrano/gbc/main/config/gbc.conf

# Creating the config directory for gbc.
mkdir -p $HOME/.config/gbc

# Moving the CLI and settings file.
mv gbc $PATH
mv gbc.conf $HOME/.config/gbc

# Giving permission to run the CLI.
chmod +x $PATH/gbc

gbc --help