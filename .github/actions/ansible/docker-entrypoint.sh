#!/bin/sh

echo "$VAULT_PASS" > /ansible/.vault_pass

mkdir ~/.ssh

ansible-vault view /ssh_key --vault-password-file="/ansible/.vault_pass" > ~/.ssh/id_rsa

chmod 0600 ~/.ssh/id_rsa

ansible-playbook \
  -e "build_sha=$GITHUB_SHA" \
  --vault-password-file="/ansible/.vault_pass" \
  /ansible/deploy.yml \
  "$@"