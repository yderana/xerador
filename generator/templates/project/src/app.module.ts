import { Module } from "@nestjs/common";
import { GraphQLModule } from "@nestjs/graphql";
import { TypeOrmModule } from "@nestjs/typeorm";
import { TemplateModule } from "@template/template.module";
import { ScheduleModule } from '@nestjs/schedule';
// tslint:disable-next-line:no-var-requires
require("dotenv").config();

@Module({
  imports: [
    TemplateModule,
    ScheduleModule.forRoot(),
    GraphQLModule.forRoot({
      formatError: (error: any) => {
        const custom: any = {
          message:
            process.env.ENVIRONTMENT.toLowerCase() === "production"
              ? "Bad Request"
              : error.message
        };
        return custom;
      },
      playground: process.env.PLAYGROUND === "true" ? {
        settings: {
          "request.credentials": 'include'
        }
      } : false,
      typePaths: ["./**/*.graphql"],
      installSubscriptionHandlers: true,
      context: ({ req, connection }) => {
        return {
          auth: req ? req.headers.authorization : connection.context.auth,
          userToken: req ? req.headers["x-cookie"] : connection.context.userToken
        }
      },
      subscriptions: {
        onConnect: (connectionParams: any) => {
          return {
            headers: connectionParams,
            auth: connectionParams.Authorization,
            userToken: connectionParams["x-cookie"]
          };
        },
      },
    }),
    TypeOrmModule.forRoot({
      type: "mysql",
      host: process.env.DATABASE_HOST,
      port: parseInt(process.env.DATABASE_PORT, 10),
      username: process.env.DATABASE_USERNAME,
      password: process.env.DATABASE_PASSWORD,
      database: process.env.DATABASE_NAME,
      entities: ["src/**/*.entity{.ts,.js}"],
      synchronize: process.env.DATABASE_SYNC === "true"
    })
  ]
})
export class ApplicationModule { }
