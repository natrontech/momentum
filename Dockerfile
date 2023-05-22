FROM golang:1.20-alpine AS backend-builder
WORKDIR /build
COPY momentum-backend/go.mod momentum-backend/go.sum momentum-backend/main.go ./
COPY momentum-backend/hooks ./hooks
RUN apk --no-cache add upx make git gcc libtool musl-dev ca-certificates dumb-init \
  && go mod tidy \
  && CGO_ENABLED=0 go build \
  && upx pocketbase

FROM node:lts-slim as ui-builder
WORKDIR /build
COPY ./momentum-ui/package*.json ./
RUN rm -rf ./node_modules
RUN rm -rf ./build
COPY ./momentum-ui .
RUN npm install
RUN npm run build

FROM alpine as runtime
WORKDIR /app/momentum
COPY --from=backend-builder /build/pocketbase /app/momentum/pocketbase
COPY ./momentum-backend/pb_migrations ./pb_migrations
COPY --from=ui-builder /build/build /app/momentum/pb_public
EXPOSE 8090
CMD ["/app/momentum/pocketbase","serve", "--http", "0.0.0.0:8090"]
