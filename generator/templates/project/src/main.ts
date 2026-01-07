import { ValidationPipe } from "@nestjs/common";
import { NestFactory } from "@nestjs/core";
import { ApplicationModule } from "./app.module";
import { HttpExceptionFilter } from "@utils/httpException.filter";
import * as bodyParser from "body-parser";
import * as cookieParser from 'cookie-parser';
// import { emitter } from "nock";
require('events').EventEmitter.defaultMaxListeners = 0;

async function bootstrap() {
  const app = await NestFactory.create(ApplicationModule);
  app.useGlobalPipes(new ValidationPipe());
  // app.useGlobalFilters(new HttpExceptionFilter());
  app.use(bodyParser.json({ limit: "50mb" }));
  app.use(bodyParser.urlencoded({ limit: "50mb", extended: true }));
  app.use(cookieParser());
  process.setMaxListeners(0);
  // emitter.setMaxListeners(0);
  app.enableCors();
  await app.listen(5002);
}
bootstrap();
