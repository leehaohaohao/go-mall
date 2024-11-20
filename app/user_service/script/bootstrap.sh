#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/user_service"
exec "$CURDIR/bin/user_service"
