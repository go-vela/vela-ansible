# SPDX-License-Identifier: Apache-2.0

#########################################################
##    docker build --no-cache -t vela-ansible:local .    ##
#########################################################

FROM alpine:latest@sha256:51b67269f354137895d43f3b3d810bfacd3945438e94dc5ac55fdac340352f48

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
