FROM ubuntu:16.04

# apt-get -y install openssh-client

LABEL version="0.0.1"
LABEL name="Streisand Builder"
LABEL author="Ray Hu"
LABEL modified="2019-09-13"

RUN apt-get update && apt-get install -y \
    git \
    python-pip \
    && rm -rf /var/lib/apt/lists/*

RUN ssh-keygen -q -t rsa -N '' -f /keys/id_rsa
