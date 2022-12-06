#!/bin/zsh

docker buildx build  --platform linux/arm64,linux/amd64/v1,linux/arm/v7 -t xiaoyi510/rustdesk-api-server -f ./DockerfileBuildx . --push