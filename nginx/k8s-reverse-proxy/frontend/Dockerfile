# pull official base image
FROM node:15.10.0-alpine3.12

# set working directory
WORKDIR /app

# add `/app/node_modules/.bin` to $PATH
ENV PATH /app/node_modules/.bin:$PATH

# install app dependencies
COPY package.json ./
COPY yarn.lock ./
RUN yarn install --silent

# add app
COPY . ./

# 本当だった build して静的コンテンツを nginx で配信するだろけど今回は割愛
# start app
CMD ["yarn", "start"]
