
# SPDX-License-Identifier: Apache-2.0

#########################################################
##    docker build --no-cache -t vela-ansible:local .    ##
#########################################################

FROM alpine:3.24.3@sha256:be2656061b536c1883c4034790e21d0a09c5aaaf51968b89368a3c0371ea2c9a

ENV ANSIBLE_VERSION=11.2.0

RUN apk --update --no-cache add \
  git \
  openssl \
  openssh-client \
  python3\
  sshpass \
  py3-pip

RUN apk --update add --virtual \
  .build-deps \
  python3-dev \
  libffi-dev \
  openssl-dev \
  build-base

RUN pip3 install --upgrade \
  pip \
  cffi --break-system-packages

RUN pip3 install \
  wheel \
  ansible==${ANSIBLE_VERSION} \
  ansible-lint --break-system-packages

RUN apk del \
  .build-deps

RUN rm -rf /var/cache/apk/*

COPY release/vela-ansible /bin/vela-ansible

ENV ANSIBLE_HOST_KEY_CHECKING=false
ENV StrictHostKeyChecking=no

ENTRYPOINT ["/bin/vela-ansible"]
