#!/bin/bash

# Check if the correct number of arguments is provided
if [ "$#" -ne 2 ]; then
  echo "Usage: $0 <json_file>" "<output_folder>"
  exit 1
fi

# Set variables
json_file=$1
output_dir=$2

# Create the output directory
mkdir -p $output_dir

# Extract "alloc" entries with "code" from the JSON data
alloc_entries=$(jq -c '.alloc | to_entries | map(select(.value.code != null)) | .[]' $json_file)

# Iterate over each entry and save the code to a file
index=1
while read -r entry; do
  code=$(echo $entry | jq -r '.value.code')
  address=$(echo $entry | jq -r '.key')

  # Check if the code is not null
  if [ "$code" != "null" ]; then
    # Remove the 0x prefix from the code
    code=$(echo $code | sed 's/^0x//')

    # Save the code to a file
    output_file="$output_dir/$address.txt"
    echo $code > $output_file

    echo "Code $index for address $address has been saved to $output_file"
    ((index++))
  fi
done <<< "$alloc_entries"

echo "Extraction completed successfully."