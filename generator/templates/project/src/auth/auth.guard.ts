import { Injectable, CanActivate, ExecutionContext } from "@nestjs/common";
import { GqlExecutionContext } from "@nestjs/graphql";
import { Observable } from "rxjs";
// tslint:disable-next-line:no-var-requires
const env = require("dotenv").config();

@Injectable()
export class AuthGuard implements CanActivate {
  canActivate(
    context: ExecutionContext
  ): boolean | Promise<boolean> | Observable<boolean> {
    const ctx = GqlExecutionContext.create(context);
    const req = ctx.getArgs().req;
    let token = [];
    let valid = false;
    /* istanbul ignore next */
    if (req && req.headers.authorization) {
      token = req.headers.authorization.split(" ");
    } else if (context.getArgByIndex(2) && context.getArgByIndex(2).auth) {
      token = context.getArgByIndex(2).auth.split(" ");
    }
    if (token.length > 0) {
      const [login, password] = Buffer.from(token[1], "base64")
        .toString()
        .split(":");

      valid =
        token[0].toLowerCase() === "basic" &&
        login === process.env.AUTH_USER &&
        password === process.env.AUTH_PASSWORD;
    }

    return valid;
  }
}
