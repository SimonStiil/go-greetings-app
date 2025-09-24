FROM scratch

ARG TARGETARCH

WORKDIR /app
COPY go-greetings-app-${TARGETARCH} /usr/bin/go-greetings-app
ENTRYPOINT ["go-greetings-app"]