#!/bin/bash

# Get the code from the user
read -p "Enter the CODE: " code

echo "Setting up the monitor..."

# Split the code by '+'
IFS='+' read -ra code_parts <<< "$code"

# Check if there are at least two parts
if [ ${#code_parts[@]} -ne 2 ]; then
  echo "Invalid Code."
  exit 1
fi

api_url="${code_parts[0]}"
api_key="${code_parts[1]}"

# Split the api_key by '&'
IFS='&' read -ra api_key_parts <<< "$api_key"

if [ ${#api_key_parts[@]} -ne 2 ]; then
  echo "Invalid Code."
  exit 1
fi

id="${api_key_parts[0]}"
token="${api_key_parts[1]}"

req_url="${api_url}/monitors/register"
req_data='{"id": "'$id'", "token": "'$token'"}'

# Make the POST request using curl
response=$(curl -s -X POST -H "Content-Type: application/json" -d "$req_data" "$req_url")

# Check for errors
if [ $? -ne 0 ]; then
  echo "Error making API request"
  exit 1
fi

api_key=$(echo "$response" | jq -r '.data.api_key')
uid=$(echo "$response" | jq -r '.data.id')

conf_data='{"api_url": "'$api_url'", "uid": "'$uid'", "secret": "'$api_key'"}'

# Write the configuration to the file
echo "$conf_data" > .conf

# Ask the user if they want to run the script at startup
read -p "Setup the Montinger Sentinel as s service? (y/n): " choice

if [[ "$choice" == "y" || "$choice" == "Y" ]]; then
  echo "Setting up the Montinger Sentinel as a service..."

  user=$(whoami)
  work_dir=$(pwd)
  file_name=$(find . -maxdepth 1 -name "montinger-sentinels-*linux" -print -quit)

  sed "s|<USER>|$user|g" "s|<WORK_DIR>|$work_dir|g" "s|<EXEC_START>|$file_name|g" montinger-sentinel.service.backup > montinger-sentinel.service

  sudo cp montinger-sentinel.service /etc/systemd/system/
  sudo systemctl enable montinger-sentinel
  sudo systemctl start montinger-sentinel
  
  if [ $? -ne 0 ]; then
    echo "Error setting up the Montinger Sentinel as a service."
    exit 1
  fi
  echo "Montinger Sentinel setup as a service successfully."
else
  echo "Montinger Sentinel setup as a service skipped."
fi

echo "Monitor setup successfully."