FROM golang:1.18 as builder

WORKDIR /src
COPY . ./
# RUN go get -d -v ./...
RUN go build -o /src/app
# RUN CGO_ENABLED=0 go build -o /go/bin/app cmd/github-telegram-notify/main.go

FROM gcr.io/distroless/base

LABEL version="1.0.0"
LABEL maintainer="wrench"
LABEL repository="https://github.com/EverythingSuckz/github-telegram-notify"
LABEL homepage="https://github.com/EverythingSuckz/github-telegram-notify"
LABEL "com.github.actions.name"="Github Telegram Notify"
LABEL "com.github.actions.description"="Notify each git commit to Telegram"
LABEL "com.github.actions.icon"="bell"
LABEL "com.github.actions.color"="blue"


COPY --from=builder /src/app .
CMD ["./app"]