import { UseGuards, UseFilters } from "@nestjs/common";
import {
  Args,
  Mutation,
  Query,
  Resolver,
  Subscription
} from "@nestjs/graphql";
import { PubSub } from "graphql-subscriptions";
import { BaseResponse } from "@utils/baseResponse";
import { AuthGuard } from "@auth/auth.guard";
import { HttpExceptionFilter } from "@utils/httpException.filter";
import { TemplateService } from "@template/providers/template.service";
import { IParamTemplate } from "@template/interfaces/iparam-template";
import { validate } from "@utils/validate";
import schemaSave from "@template/api/jsonschema/template.schema-request";
import schemaParamList from "@template/api/jsonschema/template.schema-list";
import { IParamList } from "@common/interfaces/iparam-list";

const pubSub = new PubSub();

@Resolver("Template")
@UseFilters(HttpExceptionFilter)
@UseGuards(AuthGuard)
export class TemplateResolvers {
  constructor(private readonly variableService: TemplateService) { }

  @Query("variableDetail")
  async variableDetail(
    @Args("id")
    id: string
  ): Promise<BaseResponse<IParamTemplate>> {
    const Response = await this.variableService.detail(id);
    if (Response.statusCode != 200) {
      throw new Error(Response.message);
    }

    return Response;
  }

  @Query("variableList")
  async variableList(
    @Args("filter")
    filter: IParamList
  ): Promise<BaseResponse<IParamTemplate[]>> {
    let Response = new BaseResponse<IParamTemplate[]>();
    const validateRequest = validate(schemaParamList.graphql, filter);
    if (validateRequest.valid) {
      Response = await this.variableService.fetchAll(filter);
      if (Response.statusCode != 200) {
        throw new Error(Response.message);
      }
    } else {
      throw new Error(validateRequest.errorText);
    }
    return Response;
  }

  @Mutation("createTemplate")
  async createTemplate(
    @Args("createTemplateInput") args: IParamTemplate
  ): Promise<BaseResponse<IParamTemplate>> {
    let Response = new BaseResponse<IParamTemplate>();
    const validateRequest = validate(schemaSave.post, args);
    if (validateRequest.valid) {
      Response = await this.variableService.create(args);

      if (Response.statusCode == 200) {
        // publish subscription data
        this.publish(Response.result.id);
      } else {
        throw new Error(Response.message);
      }

    } else {
      throw new Error(validateRequest.errorText);
    }

    return Response;
  }


  @Mutation("updateTemplate")
  async updateTemplate(
    @Args("id") id: string,
    @Args("updateTemplateInput") args: IParamTemplate
  ): Promise<BaseResponse<IParamTemplate>> {
    let Response = new BaseResponse<IParamTemplate>();
    const input = { ...args, id }
    const validateRequest = validate(schemaSave.post, input);
    if (validateRequest.valid) {
      Response = await this.variableService.update(input);

      if (Response.statusCode == 200) {
        // publish subscription data
        this.publish(Response.result.id);
      } else {
        throw new Error(Response.message);
      }

    } else {
      throw new Error(validateRequest.errorText);
    }

    return Response;
  }

  @Mutation("deleteTemplate")
  async deleteTemplate(
    @Args("id") id: string
  ): Promise<BaseResponse<IParamTemplate>> {
    const Response = await this.variableService.delete(id);

    if (Response.statusCode == 200) {
      // publish subscription data
      this.publish(id);
    } else {
      throw new Error(Response.message);
    }


    return Response;
  }

  /* istanbul ignore next */
  @Subscription("variableData", {
    filter(payload: any, variables: any) {
      return payload.variableData.result.id === variables.id
    }
  })
  variableData(@Args('id') id: string) {
    pubSub.subscribe("postTemplateData", (message: any) => this.newPublish(message, { id }))
    this.publish(id);
    return pubSub.asyncIterator("variableData");
  }

  /* istanbul ignore next */
  async publish(data: any) {
    pubSub.publish("postTemplateData", data);
  }

  /* istanbul ignore next */
  async newPublish(message: any, filter: any) {
    if (message == filter.id) {
      const getData = await this.variableService.detail(filter.id)
      if (getData.success) {
        pubSub.publish(`variableData`, { variableData: getData })
      }
    }
  }
}
