FROM golang:1.16.3-alpine AS builder

RUN apk add --update --no-cache make
ENV PROJECT_NAME=tfdocs-format-template
ENV SOURCE_DIR=/go/src/${PROJECT_NAME}
# replace tdocs-format-template with actual formatter plugin name.
WORKDIR ${SOURCE_DIR}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN make build

################

FROM quay.io/terraform-docs/terraform-docs:latest
# added because custom plugin should be installed into the $HOME
ENV PROJECT_NAME=tfdocs-format-template
ENV BUILD_CONTAINER_WD=/go/src/${PROJECT_NAME}
ENV USER=docker
ENV UID=12345
ENV GID=23456
# creating user in docker container
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "$(pwd)" \
    --ingroup "root" \
    --uid "$UID" \
    "$USER"
# after this, all commands will be executed from this user
USER docker

COPY --from=builder ${BUILD_CONTAINER_WD}/bin/linux-amd64/${PROJECT_NAME} $HOME/.tfdocs.d/plugins/${PROJECT_NAME}

CMD ["/bin/sh","-c","terraform-docs","$@"]