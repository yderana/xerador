import {
  ExceptionFilter,
  Catch,
  ArgumentsHost,
  HttpException,
  ValidationError,
  HttpStatus
} from "@nestjs/common";
import { Request, Response } from "express";
// import { config, captureException } from "raven";

// config(process.env.SENTRY_URL).install();
@Catch(HttpException)
export class RestHttpExceptionFilter implements ExceptionFilter {
  catch(exception: ValidationError, host: ArgumentsHost) {
    const ctx = host.switchToHttp();
    const response = ctx.getResponse();
    // const request = ctx.getRequest();

    // captureException(response);
    /* istanbul ignore next */
    const status =
      exception instanceof HttpException
        ? exception.getStatus()
        : HttpStatus.INTERNAL_SERVER_ERROR;

    switch (status) {
      /* istanbul ignore next */
      case 401:
        response.status(401).json({
          statusCode: 401,
          error: "Not Authorized",
          message: "You are not authorized"
          // timestamp: new Date().toISOString(),
          // path: request.url,
        });
        break;
      /* istanbul ignore next */
      case 403:
        response.status(401).json({
          statusCode: 401,
          error: "Not Authorized",
          message: "You are not authorized"
          // timestamp: new Date().toISOString(),
          // path: request.url,
        });
        break;
      /* istanbul ignore next */
      case 500:
        response.status(500).json({
          statusCode: 500,
          error: "Server Error",
          message: "Internal Server Error"
        });
      /* istanbul ignore next */
      default:
        response.status(400).json({
          statusCode: 400,
          error: "Bad Request",
          message: "Invalid request format"
        });
        break;
    }
  }
}

// tslint:disable-next-line:max-classes-per-file
@Catch(HttpException)
export class HttpExceptionFilter implements ExceptionFilter {
  /* istanbul ignore next */
  catch(exception: HttpException, host: ArgumentsHost) {
    const status = exception.getStatus();

    // captureException(exception);
    switch (status) {
      case 401:
        throw new Error("You are not authorized");
      case 403:
        throw new Error("You are not authorized");
      case 500:
        throw new Error("Internal Server Error");
      default:
        throw new Error("Bad Request");
    }
  }
}
