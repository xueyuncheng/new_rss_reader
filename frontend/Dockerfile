FROM node:latest AS builder

COPY . /app

RUN npm install
RUN npm run build


FROM nginx:latest

COPY --from=builder /app/dist /usr/share/nginx/html

CMD ["nginx", "-g", "daemon off;"]