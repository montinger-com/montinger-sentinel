#!/bin/bash

api_url="https://api.github.com/repos/montinger-com/montinger-sentinel/releases/latest"
response=$(curl -s "$api_url")

if [ $? -ne 0 ]; then
  echo "Error making API request"
  exit 1
fi

tag_name=$(echo "$response" | jq -r '.tag_name')
version=${tag_name#v}

echo "Downloading the latest release..."
curl -s -L "https://github.com/montinger-com/montinger-sentinel/releases/download/$tag_name/montinger-sentinel-$version-linux.tar.gz" -o "montinger-sentinel.tar.gz"

if [ $? -ne 0 ]; then
  echo "Error downloading the release"
  exit 1
fi

echo "Extracting the release..."
tar -xzf montinger-sentinel.tar.gz --overwrite

if [ $? -ne 0 ]; then
  echo "Error extracting the release"
  exit 1
fi

echo "Stopping service..."
sudo systemctl stop montinger-sentinel

if [ $? -ne 0 ]; then
  echo "Error stopping service"
  exit 1
fi

echo "Updating files..."
mv -f montinger-sentinel-$version-linux montinger-sentinel

if [ $? -ne 0 ]; then
  echo "Error updating files"
  exit 1
fi

echo "Starting service..."
sudo systemctl start montinger-sentinel

if [ $? -ne 0 ]; then
  echo "Error starting service"
  exit 1
fi

echo "Cleaning up..."
rm montinger-sentinel.tar.gz

if [ $? -ne 0 ]; then
  echo "Error cleaning up"
  exit 1
fi

echo "Update complete"