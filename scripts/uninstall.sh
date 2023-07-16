#!/bin/bash

# Removing CLI
sudo rm /usr/bin/gbc
# Removing config directory
rm -rf $HOME/.config/gbc

echo "gbc uninstalled successfully."