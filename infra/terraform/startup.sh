#!/bin/bash
apt-get update
apt-get install -y curl git docker.io
usermod -aG docker $USER
