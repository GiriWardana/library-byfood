FROM node:20

WORKDIR /app

COPY package.json yarn.lock ./
RUN yarn install

# Copy environment file for build stage
# COPY .env.production .env.production

# Copy all other files
COPY . .

# Build Next.js app (will read .env.production automatically)
RUN yarn build

EXPOSE 3000

CMD ["yarn", "start"]
