#!/bin/bash

# Directory where keys will be stored
KEYS_DIR="./bls-keys"

# Number of validator keys to generate
NUM_KEYS=50

# Create folder if it doesn't exist
mkdir -p $KEYS_DIR

# Loop to generate 50 validator keys
for i in $(seq 0 $((NUM_KEYS-1)))
do
    echo "Generating key pair for validator $i..."
    
    # Run the Go command to generate keys for each validator
    go run main.go
    
    # Move the generated keys into the KEYS_DIR folder
    mv validatorKey.pem $KEYS_DIR/validator${i}-key.pem
done

echo "🎉 All BLS keys generated and saved in $KEYS_DIR"
