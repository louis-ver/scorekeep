FROM alpine

ENV ANSIBLE_HOST_KEY_CHECKING=False

RUN apk update \
  && apk add \
  ansible \
  gcc \
  python3-dev \
  libc-dev \
  libffi-dev \
  openssl-dev \
  && pip3 install \
    --upgrade paramiko

COPY docker-entrypoint.sh /docker-entrypoint.sh
COPY ssh_key /ssh_key
COPY ansible-root/ /ansible

RUN mkdir -p /etc/ansible/

COPY hosts /etc/ansible/hosts

ENTRYPOINT ["/docker-entrypoint.sh"]
