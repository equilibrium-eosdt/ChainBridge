#!/usr/bin/env bash

set -e 

if [[ -n "${BRIDGE_CONFIG}" ]]; then
    echo ${BRIDGE_CONFIG} > /etc/config.json
fi

exec "$@"

