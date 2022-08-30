FROM nginx:1.23.1-alpine

WORKDIR /usr/share/nginx/html
RUN rm -rf * && \
  apk --no-cache add tzdata && \
  cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
  apk del tzdata
