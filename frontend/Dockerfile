FROM node:20-alpine

WORKDIR /app
COPY . .
RUN npm install && npm run build

# Serve with a basic web server
RUN npm install -g serve
CMD ["serve", "-s", "dist", "-l", "3000"]
