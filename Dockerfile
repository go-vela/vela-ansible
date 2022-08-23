# Copyright (c) 2022 Target Brands, Inc. All rights reserved.
#
# Use of this source code is governed by the LICENSE file in this repository.

#########################################################
##    docker build --no-cache -t vela-ansible:local .    ##
#########################################################

FROM alpine:latest

ENV ANSIBLE_VERSION 6.2.0

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
        cffi

RUN pip3 install \
        wheel \
        ansible==${ANSIBLE_VERSION} \
        ansible-lint

RUN apk del \
        .build-deps

RUN rm -rf /var/cache/apk/*

COPY release/vela-ansible /bin/vela-ansible

ENV ANSIBLE_HOST_KEY_CHECKING false
ENV StrictHostKeyChecking no

ENTRYPOINT ["/bin/vela-ansible"]
