#!/bin/sh

echo "$VAULT_PASS" > ~/.vault_pass.txt

mkdir -p /etc/ansible
echo "$SERVER_IP" > /etc/ansible/hosts

mkdir ~/.ssh

ansible-vault --vault-password-file="~/.vault_pass.txt" view ~/ssh_key > ~/.ssh/id_rsa

chmod 0600 ~/.ssh/id_rsa

ansible all \
  -m ping \
  -u githubactions
