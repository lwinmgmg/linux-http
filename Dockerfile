FROM golang:1.21-bookworm AS compiler

ARG OUTPUT_DIR=bin
ARG OUTPUT_BINARY_NAME=linux-http-go
ARG WORKDIR_NAME=/build

WORKDIR ${WORKDIR_NAME}

COPY . ${WORKDIR_NAME}/.

RUN pwd

RUN go get && go mod vendor

RUN go build -o "${OUTPUT_DIR}/${OUTPUT_BINARY_NAME}"

FROM debian:bookworm

ARG OUTPUT_DIR=bin
ARG OUTPUT_BINARY_NAME=linux-http-go
ARG WORKDIR_NAME=/build

ENV USER_GOLANG_VERSION=1.21
ENV LINUX_USER=linux
ENV LINUX_USER_HOME_DIR="/home/${LINUX_USER}"
ENV LINUX_USER_UID=2023
ENV LINUX_WORKDIR="/etc/linux-http"
ENV LINUX_BINARY_DIR="/usr/local/linux-http/bin"
ENV PATH=${PATH}:${LINUX_BINARY_DIR}
ENV GIN_MODE=release

# APP ENV
ENV LH_ENV_PATH="${LINUX_WORKDIR}/.env"
ENV DB_DIR="/var/lib/linux-http"
ENV LH_DB_PATH="${DB_DIR}/app.db"

RUN groupadd --gid ${LINUX_USER_UID} ${LINUX_USER} \
    && useradd --uid ${LINUX_USER_UID} --gid ${LINUX_USER_UID} \
    --home-dir ${LINUX_USER_HOME_DIR} --shell /bin/bash ${LINUX_USER}

RUN apt update
RUN apt install -y ca-certificates

RUN mkdir ${LINUX_USER_HOME_DIR}
RUN mkdir ${LINUX_WORKDIR}
RUN mkdir ${DB_DIR}
RUN chown ${LINUX_USER}:${LINUX_USER} -R ${LINUX_USER_HOME_DIR}
RUN chown ${LINUX_USER}:${LINUX_USER} -R ${LINUX_WORKDIR}
RUN chown ${LINUX_USER}:${LINUX_USER} -R ${DB_DIR}

# Copy compiled binary
COPY --from=compiler --chown=${LINUX_USER}:${LINUX_USER} "${WORKDIR_NAME}/${OUTPUT_DIR}/${OUTPUT_BINARY_NAME}" "${LINUX_BINARY_DIR}/${OUTPUT_BINARY_NAME}"

# To save vendor package if they go unsupported
COPY --from=compiler --chown=${LINUX_USER}:${LINUX_USER} "${WORKDIR_NAME}/vendor" "${LINUX_USER_HOME_DIR}/vendor"

COPY --from=compiler --chown=${LINUX_USER}:${LINUX_USER} "${WORKDIR_NAME}/.env.example" ${LINUX_WORKDIR}/.env.example

VOLUME [ "${DB_DIR}" ]

USER ${LINUX_USER}

WORKDIR ${LINUX_WORKDIR}

RUN mv .env.example .env

EXPOSE 8888

CMD [ "linux-http-go" ]