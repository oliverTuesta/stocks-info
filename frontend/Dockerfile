# Minimal multi-stage Dockerfile for the Vite + Vue frontend

FROM node:20.19-alpine AS builder
WORKDIR /app

# Install build deps (git may be needed for some packages)
RUN apk add --no-cache git build-base python3

# Install dependencies
COPY package.json package-lock.json* ./
RUN npm install --silent

# Copy source and build
COPY . .
RUN npm run build-only

FROM nginx:stable-alpine
# Remove default nginx index page
RUN rm -rf /usr/share/nginx/html/*

# Copy built files from builder
COPY --from=builder /app/dist /usr/share/nginx/html

# Add SPA-friendly nginx config
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
