FROM node:8-alpine

# Drake support
ENV BRIGDRAKE_VERSION=v0.21.0
RUN apk update \
  && apk add ca-certificates \
  && rm -rf /var/cache/apk/* \
  && mkdir -p /brigade/bin \
  && wget -q -O /brigade/bin/brigdrake-worker \
    https://github.com/lovethedrake/brigdrake/releases/download/${BRIGDRAKE_VERSION}/brigdrake-worker-linux-amd64 \
  && chmod 755 /brigade/bin/brigdrake-worker

# brigade.js support 
WORKDIR /home/src
COPY brigade-worker/ /home/src/
RUN yarn build

CMD yarn run test
