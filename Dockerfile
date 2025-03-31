FROM golang:1.24.1 AS builder

WORKDIR /app
COPY main.go .
RUN GOOS=js GOARCH=wasm go build -o main.wasm main.go
RUN cp $(go env GOROOT)/lib/wasm/wasm_exec.js .

FROM nginx:alpine

COPY --from=builder /app/main.wasm /usr/share/nginx/html/
COPY --from=builder /app/wasm_exec.js /usr/share/nginx/html/
COPY index.html /usr/share/nginx/html/

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]