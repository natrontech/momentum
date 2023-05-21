FROM golang:1.20-alpine AS builder
WORKDIR /build
COPY momentum-backend/go.mod momentum-backend/go.sum momentum-backend/main.go ./
COPY momentum-backend/hooks ./hooks
RUN apk --no-cache add upx make git gcc libtool musl-dev ca-certificates dumb-init \
  && go mod tidy \
  && CGO_ENABLED=0 go build \
  && upx pocketbase

FROM alpine
WORKDIR /app/pb
COPY --from=builder /build/pocketbase /app/momentum-backend/pocketbase
COPY momentum-backend/pb_migrations ./pb_migrations
COPY ./momentum-ui/build /app/momentum-backend/pb_public
EXPOSE 8090
CMD ["/app/momentum-backend/pocketbase","serve", "--http", "0.0.0.0:8090"]
