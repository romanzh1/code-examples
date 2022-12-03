FROM golang:latest as build
RUN mkdir -p ${GO}/temp
COPY . ${GO}/temp
WORKDIR ${GO}/temp
RUN go build -o /exe main.go
RUN go build -o /deps deps.go

FROM ubuntu:bionic AS base

ARG UID=1000
ARG GID=1000

# Install Go
RUN apt-get update && apt-get install -y ca-certificates

# Install (Chromium) Playwright dependencies
RUN apt-get install -y libwoff1 \
                       libopus0 \
                       libwebp6 \
                       libwebpdemux2 \
                       libenchant1c2a \
                       libgudev-1.0-0 \
                       libsecret-1-0 \
                       libhyphen0 \
                       libgdk-pixbuf2.0-0 \
                       libegl1 \
                       libnotify4 \
                       libxslt1.1 \
                       libevent-2.1-6 \
                       libgles2 \
                       libvpx5 \
                       libnss3 \
                       libxss1 \
                       libasound2 \
                       fonts-noto-color-emoji \
                       libdbus-glib-1-2 \
                       libxt6 \
                       ffmpeg

# If group exists, append user to group +audio,video, otherwise add group
RUN ( grep -q -E "^.*:.*:${GID}:.*$" /etc/group && \
      groupadd -r appuser --gid ${GID} && \
      useradd --uid ${UID} -r -g appuser -G audio,video appuser ) || \
    groupadd -r appuser && \
    useradd --uid ${UID} -r -g appuser -G audio,video appuser

RUN mkdir -p /app && chown appuser:appuser /app

# Playwright works from the user's folder ~/.cache/. But there are not enough rights inside the docker to run playwright. 
# Rights in the docker container cannot be configured.
# Therefore, you need to create a new user in the docker container and run playwright from the new user's folder.
RUN mkdir -p "/home/appuser" && chown -R appuser:appuser "/home/appuser"
USER appuser
COPY --from=build /exe /app/exe
COPY --from=build /deps /app/deps
COPY . /app/

WORKDIR /app
RUN /app/deps
ENTRYPOINT ["/app/exe"]
