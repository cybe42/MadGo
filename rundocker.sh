#!/bin/bash

ROOT_UID=0
E_NOTROOT=87


if [ "${UID:-$(id -u)}" -ne "$ROOT_UID" ]; then
    echo 'Error: Please run this script with root priviliges.'
    exit $E_NOTROOT
fi
docker run -it cybe42/madgo