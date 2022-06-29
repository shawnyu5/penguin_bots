FROM node:16.15.1 AS build

WORKDIR /app

COPY ./package*.json ./

RUN npm install --verbose

# copy source files
COPY ./src ./src

COPY ./tsconfig.json ./

RUN npm run build

FROM node:16.15.1-alpine AS production

WORKDIR /app

COPY ./package*.json ./
RUN npm install --omit=dev --verbose

COPY --from=build /app/build ./build

CMD npm run run
