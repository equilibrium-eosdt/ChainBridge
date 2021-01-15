#!/usr/bin/env bash

set -e 

if [[ -n "${BRIDGE_CONFIG}" ]]; then
    echo ${BRIDGE_CONFIG} > /etc/config.json
fi

KEY1N=$(cat /etc/config.json | jq '.chains[] | select(.id == "0") | .from')
KEY1N="${KEY1N%\"}"
KEY1N="${KEY1N#\"}"
KEY1N="/etc/keys/${KEY1N}.key"

if [[ -n "${KEY1}" ]]; then
    echo ${KEY1} > ${KEY1N}
fi

KEY2N=$(cat /etc/config.json | jq '.chains[] | select(.id == "1") | .from')
KEY2N="${KEY2N%\"}"
KEY2N="${KEY2N#\"}"
KEY2N="/etc/keys/${KEY2N}.key"

if [[ -n "${KEY2}" ]]; then
    echo ${KEY2} > ${KEY2N}
fi

#exec "$@"
script -q -c 'usr/bin/bridge --config /etc/config.json --keystore /etc/keys --latest --fresh --graylog gl.oxygen-dev.net:12201 --metrics --metricsPort 9880' << ENDDOC
1
1
ENDDOC

