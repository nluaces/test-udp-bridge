#! /bin/sh

podman build -t quay.io/skupper/udp-server:latest -f Dockerfile.server
podman build -t quay.io/skupper/udp-client:latest -f Dockerfile.client

podman push quay.io/skupper/udp-server
podman push quay.io/skupper/udp-client