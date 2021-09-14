#!/usr/bin/env bash

set -e

if [[ -n "${BRIDGE_CONFIG}" ]]; then
    echo "${BRIDGE_CONFIG}" > /etc/config.json
fi

if [[ -n "${KEYSTORE_KEYS}" ]]; then
    chains_number="$(echo "${BRIDGE_CONFIG}" | jq -r '.chains | length')"
    for chain_index in $(seq 1 "$chains_number")
    do
        address=$(echo "${BRIDGE_CONFIG}" | jq -r ".chains[${chain_index}].from")
        keyfile="/etc/keys/${address}.key"
        echo "${KEYSTORE_KEYS}" | jq ".[${chain_index}]" > "${keyfile}"
    done
else
    if [[ -n "${KEY1}" ]]; then
        KEYFILE1="/etc/keys/$(echo ${KEY1} | jq -r .address).key"
        echo ${KEY1} > ${KEYFILE1}
    fi

    if [[ -n "${KEY2}" ]]; then
        KEYFILE2="/etc/keys/$(echo ${KEY2} | jq -r .address).key"
        echo ${KEY2} > ${KEYFILE2}
    fi

    if [[ -n "${KEY3}" ]]; then
        KEYFILE3="/etc/keys/$(echo ${KEY3} | jq -r .address).key"
        echo ${KEY3} > ${KEYFILE3}
    fi
fi

script -q -c 'usr/bin/bridge --config /etc/config.json --keystore /etc/keys --blockstore ${BLOCKSTORE} --graylog ${GRAYLOG} ${PARAMS} --metrics --metricsPort 9880'

