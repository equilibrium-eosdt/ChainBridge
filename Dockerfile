# Copyright 2020 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only

FROM  golang:1.13-stretch AS builder
ADD . /src
WORKDIR /src
RUN go mod download \
    && cd cmd/chainbridge && go build -o /bridge . \
    && chmod +x /bridge

# # final stage
FROM debian:stretch-slim
COPY --from=builder /bridge /usr/bin/bridge
COPY entrypoint.sh /
RUN apt-get -y update && apt-get -y upgrade && apt-get install ca-certificates wget -y
RUN wget -P /usr/local/bin/ https://chainbridge.ams3.digitaloceanspaces.com/subkey-rc6 \
  && mv /usr/local/bin/subkey-rc6 /usr/local/bin/subkey \
  && chmod +x /usr/local/bin/subkey \
  && mkdir -p /etc/keys

ENTRYPOINT ["/entrypoint.sh"]
EXPOSE 9880
#CMD ["/usr/bin/bridge","--config","/etc/config.json","--verbosity","info","--latest"]
