#!/usr/bin/env bash

LOCAL_TROJAN_DIR=$HOME/.cfssl_workdir/trojan-go
SERVER_TROJAN_DIR=$HOME/trojan-go-linux-amd64

mkdir -p config
# rsync -azP $LOCAL_TROJAN_DIR/server.json config
rsync -azP certs/{server,server-key,ca}.pem config/server.json \
    virmach:$SERVER_TROJAN_DIR
rsync -azP certs/*.pem config/server.json \
    $LOCAL_TROJAN_DIR