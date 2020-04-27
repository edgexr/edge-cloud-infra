#!/bin/sh

# exit immediately on failure
set -e

REGION=$1

if [ -z "$REGION" ]; then
    echo "Usage: setup-region.sh <region>"
    exit 1
fi
echo "Setting up infra Vault region $REGION"

# autoprov approle
# Just need access to influx db credentials
cat > /tmp/autoprov-pol.hcl <<EOF
path "auth/approle/login" {
  capabilities = [ "create", "read" ]
}

path "secret/data/+/accounts/influxdb" {
  capabilities = [ "read" ]
}

path "pki-regional/issue/*" {
  capabilities = [ "read", "update" ]
}
EOF
vault policy write autoprov /tmp/autoprov-pol.hcl
rm /tmp/autoprov-pol.hcl
vault write auth/approle/role/autoprov period="720h" policies="autoprov"
# get autoprov app roleID and generate secretID
vault read auth/approle/role/autoprov/role-id
vault write -f auth/approle/role/autoprov/secret-id

# Note: Shepherd uses CRM's Vault access creds.