FROM node:10
WORKDIR /usr/src/app
COPY package*.json ./
RUN npm i -g @nestjs/cli
RUN npm install
COPY . .
EXPOSE 5000
CMD [ "npm", "run", "start:prod" ]
