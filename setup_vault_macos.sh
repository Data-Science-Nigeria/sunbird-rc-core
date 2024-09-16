#!/bin/bash

# This script does the following things:
# * Start a Vault instance with vault/vault.json configuration
# * Unseal the Vault
# * Create a v2 kv engine
# * Prints the unseal keys and root token
# * This script does not automatically unseal Vault on restarts, it only works with fresh installations

COMPOSE_FILE="${1:-docker-compose.yml}"
SERVICE_NAME="${2:-vault}"

#echo "Removing the db-data, es-data and vault-data directories"
#rm -rf db-data es-data vault-data

echo "Setting up $SERVICE_NAME in $COMPOSE_FILE"

docker compose -f "$COMPOSE_FILE" up -d "$SERVICE_NAME"

# Function to check if Vault is ready
check_vault_status() {
  vault_status=$(docker compose -f "$COMPOSE_FILE" exec -T "$SERVICE_NAME" vault status 2>&1)
  if [[ $vault_status == *"connection refused"* ]]; then
    echo "Unable to connect to Vault. Waiting for Vault to start..."
    return 1
  elif [[ $vault_status == *"Sealed             true"* ]]; then
    echo "Vault is sealed. Waiting for unsealing..."
    return 0
  else
    echo "Unsealed and up. Moving to next steps."
    return 0
  fi
}

# Wait for Vault service to become available
until check_vault_status; do
    echo "Waiting for Vault service to start..."
    sleep 1
done

# Initialize Vault if not already initialized
if [[ $vault_status == *"Initialized             true"* ]]; then
    echo "Vault is already initialized. Unsealing if necessary..."
else
  # Initialize Vault and save keys to files
  docker compose -f "$COMPOSE_FILE" exec -T "$SERVICE_NAME" vault operator init > ansi-keys.txt
  sed $'s/\x1B\[[0-9;]*[JKmsu]//g' ansi-keys.txt > keys.txt
fi

# Unseal Vault
for i in {1..3}; do
  key=$(sed -n "s/Unseal Key $i: \(.*\)/\1/p" keys.txt)
  docker compose -f "$COMPOSE_FILE" exec -T "$SERVICE_NAME" vault operator unseal "$key" < /dev/null
done

# Get the root token
root_token=$(sed -n 's/Initial Root Token: \(.*\)/\1/p' keys.txt | tr -dc '[:print:]')

# Update the .env file with the root token
if [[ -f ".env" ]]; then
  sed -i.bak "s/VAULT_TOKEN=.*/VAULT_TOKEN=$root_token/" ".env"
else
  echo "VAULT_TOKEN=$root_token" > .env
fi

# Enable KV v2 secrets engine if Vault was just initialized
if [[ $vault_status != *"Initialized             true"* ]]; then
  docker compose -f "$COMPOSE_FILE" exec -T -e VAULT_TOKEN="$root_token" "$SERVICE_NAME" vault secrets enable -path=kv kv-v2
fi

echo -e "\nNOTE: KEYS ARE STORED IN keys.txt"

# Clean up temporary files
rm -f ansi-keys.txt
