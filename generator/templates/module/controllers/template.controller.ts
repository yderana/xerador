import {
  Controller,
  Delete,
  Get,
  Param,
  Post,
  Body,
  Query,
  UseGuards,
  UseFilters,
  Req,
  UseInterceptors,
  Res as Resp,
  Put
} from "@nestjs/common";
import { AuthGuard } from "@auth/auth.guard";
import { RestHttpExceptionFilter } from "@utils/httpException.filter";
import { Response } from "express";
import schemaParamList from "@template/api/jsonschema/template.schema-list";
import schemaSave from "@template/api/jsonschema/template.schema-request";
import { validate } from "@utils/validate";
import { Res } from "@utils/response";
import { TemplateService } from "@template/providers/template.service";
import { IParamListTemplate } from "@template/interfaces/iparam-list-template";
import { IParamTemplate } from "@template/interfaces/iparam-template";
import { TemplateResolvers } from "@template/controllers/template.resolvers";

@Controller("template")
@UseFilters(RestHttpExceptionFilter)
@UseGuards(AuthGuard)
export class TemplateController {
  constructor(private readonly variableService: TemplateService, private readonly variableResolvers: TemplateResolvers) { }

  @Get()
  async findAll(
    @Query("page") page: number,
    @Query("limit") limit: number,
    @Query("sort") sort: string,
    @Query("orderBy") orderBy: string,
    @Query("noCache") noCache: boolean,
    @Query("q") query: string,
    @Resp() response: Response
  ) {
    let result: any = null;
    const params: IParamListTemplate = {
      page,
      limit,
      sort,
      orderBy,
      query,
      noCache,
    };

    const validateParam = validate(schemaParamList.rest, params);
    if (validateParam.valid) {
      result = await this.variableService.fetchAll(params);
    } else {
      result = Res({ statusCode: 400, message: validateParam.errorText });
    }
    return response.status(result.statusCode).send(result);
  }

  @Get("/:id")
  async findOneById(
    @Param("id") id: string, @Resp() response: Response
  ) {
    const result = await this.variableService.detail(id);
    return response.status(result.statusCode).send(result);
  }

  @Post()
  @UseInterceptors()
  async post(@Body() input: IParamTemplate, @Resp() response: Response) {
    let result: any = null;
    const validateRequest = validate(schemaSave.post, input);
    if (validateRequest.valid) {
      result = await this.variableService.create(input);

      if (result.statusCode == 200) {
        // publish to subscription
        this.variableResolvers.publish(result.result.id);
      }

    } else {
      result = Res({ statusCode: 400, message: validateRequest.errorText });
    }

    return response.status(result.statusCode).send(result);
  }

  @Put("/:id")
  @UseInterceptors()
  async put(@Param("id") id: string, @Body() input: IParamTemplate, @Resp() response: Response) {
    let result: any = null;
    input = { ...input, id }
    const validateRequest = validate(schemaSave.put, input);
    if (validateRequest.valid) {
      result = await this.variableService.update(input);

      if (result.statusCode == 200) {
        // publish to subscription
        this.variableResolvers.publish(result.result.id);
      }

    } else {
      result = Res({ statusCode: 400, message: validateRequest.errorText });
    }

    return response.status(result.statusCode).send(result);
  }

  @Delete("/:id")
  async remove(@Param("id") id: string, @Resp() response: Response) {
    const result = await this.variableService.delete(id);
    return response.status(result.statusCode).send(result);
  }
}
