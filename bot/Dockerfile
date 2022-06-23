FROM node:16.15.1

WORKDIR /bot

COPY ./package*.json ./

RUN npm install

# copy source files
COPY ./src ./src

COPY ./tsconfig.json ./

CMD npm start
