FROM node:22-alpine AS builder

WORKDIR /app

COPY package*.json ./
COPY .env .env
RUN npm install

COPY . .

RUN npm run build

FROM node:22-alpine

WORKDIR /app

COPY --from=builder /app ./

ENV NODE_ENV=production
EXPOSE 3000

CMD ["npm", "start"]
