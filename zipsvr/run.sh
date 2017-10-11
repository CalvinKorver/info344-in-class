#!/usr/bin/env bash
docker rm -f zipsvr

docker run -d \
-p 443:443 \
--name zipsvr \
-v  /Users/calvinkorver/Documents/code/go/src/github.com/calvinkorver/info344-in-class/zipsvr:/tls:ro \
-e TLSCERT=/tls/fullchain.pem \
-e TLSKEY=/tls/privkey.pem \
cjkorver/zipsvr