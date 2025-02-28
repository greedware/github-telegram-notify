FROM golang:1.23 as builder
WORKDIR /src
COPY . .
RUN go build -o /src/app -ldflags="-w -s" .
FROM gcr.io/distroless/base

LABEL version="1.1.3"
LABEL maintainer="corel-frim"
LABEL repository="https://github.com/greedware/github-telegram-notify"
LABEL homepage="https://github.com/greedware/github-telegram-notify"
LABEL "com.github.actions.name"="Github Telegram Notify"
LABEL "com.github.actions.description"="Notify each git action to Telegram"
LABEL "com.github.actions.icon"="bell"
LABEL "com.github.actions.color"="blue"

COPY --from=builder /src/app /
ENTRYPOINT ["/app"]