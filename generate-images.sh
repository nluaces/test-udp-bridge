#! /bin/sh

podman build -t quay.io/nluaces/udp-server:latest -f Dockerfile.server
podman build -t quay.io/nluaces/udp-client:latest -f Dockerfile.client

podman push quay.io/nluaces/udp-server
podman push quay.io/nluaces/udp-client