#!/bin/bash

# Directory where the validator keys are stored
KEYS_DIR="./bls-keys"

# Output file for storing all public keys
PUBKEYS_FILE="./all_validator_pubkeys.txt"

# Number of validator keys to process
NUM_KEYS=50

# Create or clear the pubkeys file
echo "" > $PUBKEYS_FILE

# Loop through each generated key file to extract the public key
for i in $(seq 0 $((NUM_KEYS-1)))
do
    # Path to the validator's private key
    PRIVATE_KEY_PATH="$KEYS_DIR/validator${i}-key.pem"

    # Check if the private key file exists
    if [ -f "$PRIVATE_KEY_PATH" ]; then
        echo "Extracting public key for validator $i..."

        # Extract the public key from the private key
        # Replace 'bls-keygen' with your actual command that extracts the public key from the private key
        PUBLIC_KEY=$(bls-keygen publickey $PRIVATE_KEY_PATH)

        # Append the public key to the file
        echo "$PUBLIC_KEY" >> $PUBKEYS_FILE
    else
        echo "Warning: Private key file $PRIVATE_KEY_PATH not found!"
    fi
done

echo "🎉 All BLS public keys have been extracted and saved to $PUBKEYS_FILE"
