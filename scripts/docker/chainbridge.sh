#!/bin/sh
# This script is intended to run in a container automatically.
# It's useless to try launch it out of that context.

export PATH=$PATH:/root/go/bin

chainbridge \
    --config /root/config.json \
    --verbosity info \
    --testkey alice \
    --latest
