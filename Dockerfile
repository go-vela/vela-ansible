# SPDX-License-Identifier: Apache-2.0

#########################################################
##    docker build --no-cache -t vela-ansible:local .    ##
#########################################################

FROM alpine:latest@sha256:c5b1261d6d3e43071626931fc004f70149baeba2c8ec672bd4f27761f8e1ad6b

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
        cffi --break-system-packages

RUN pip3 install \
        wheel \
        ansible==${ANSIBLE_VERSION} \
        ansible-lint --break-system-packages

RUN apk del \
        .build-deps

RUN rm -rf /var/cache/apk/*

COPY release/vela-ansible /bin/vela-ansible

ENV ANSIBLE_HOST_KEY_CHECKING false
ENV StrictHostKeyChecking no

ENTRYPOINT ["/bin/vela-ansible"]
