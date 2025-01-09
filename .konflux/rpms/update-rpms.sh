#!/bin/bash

BASE_IMAGE=registry.access.redhat.com/ubi9/ubi-minimal:latest
podman run $BASE_IMAGE cat /etc/yum.repos.d/ubi.repo > ./ubi.repo

podman run \
  -e BASE_IMAGE=$BASE_IMAGE \
  -v ./:/rpms -w /rpms -it  fedora sh -c "
    set -x
    dnf install -y python3-dnf skopeo pip
    python -m pip install https://github.com/konflux-ci/rpm-lockfile-prototype/archive/refs/tags/v0.13.1.tar.gz
    rpm-lockfile-prototype --image $BASE_IMAGE rpms.in.yaml
  "

