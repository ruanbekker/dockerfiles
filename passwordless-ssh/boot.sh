#!/usr/bin/env bash

create_user() {
  SSH_USER="user-$(openssl rand -hex 12)"
  export SSH_USER=${SSH_USER}
  adduser -D -s /bin/bash -G ssh-users ${SSH_USER}
  echo -e "\n" | passwd ${SSH_USER}
}

add_user(){
  adduser -D -s /bin/bash -G ssh-users ${SSH_USER}
  echo -e "\n" | passwd ${SSH_USER}
}

if [ -z ${SSH_USER} ]
then
  create_user
else
  add_user
fi

export SSH_USER=${SSH_USER}
echo "SSH User is: ${SSH_USER}"

bash /motd-generator.sh
/usr/sbin/sshd -f /etc/ssh/sshd_config -D
