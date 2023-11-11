#!/bin/bash
_create_voulumes() {
	# ENV LH_ENV_PATH="${LINUX_WORKDIR}/.env"
    # ENV DB_DIR="/var/lib/linux-http"
    # ENV LH_DB_PATH="${DB_DIR}/app.db"
    mkdir -p $DB_DIR | :
    chmod 00700 $DB_DIR | :
}

_main() {
    echo "$@"
    _create_voulumes
}

echo "$@"
_main "hello"

