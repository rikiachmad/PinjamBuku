FROM node:18 as builder
WORKDIR '/app'

COPY package.json .
COPY package-lock.json .
RUN npm install -g npm@6.14.8
RUN npm install

COPY . .

RUN npm run build

FROM nginx:alpine
WORKDIR /usr/share/nginx/html

COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /app/build .

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]