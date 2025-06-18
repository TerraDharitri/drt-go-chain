#!/bin/bash

# Directory where keys will be stored
KEYS_DIR="./wallet-keys"

# Number of wallet keys to generate
NUM_KEYS=60

# Create folder if it doesn't exist
mkdir -p $KEYS_DIR

# Loop to generate 50 wallet keys
for i in $(seq 0 $((NUM_KEYS-1)))
do
    echo "Generating key pair for wallet $i..."
    
    # Run the Go command to generate keys for each wallet
    go run main.go
    
    # Move the generated keys into the KEYS_DIR folder
    mv walletKey.pem $KEYS_DIR/wallet${i}-key.pem
done

echo "🎉 All BLS keys generated and saved in $KEYS_DIR"
