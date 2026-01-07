import { errorResponse } from "@helpers/generateErrorResponse";
import { Injectable } from "@nestjs/common";
import { BaseResponse } from "@utils/baseResponse";
import { Res } from "@utils/response";
import { IParamListTemplate } from "@template/interfaces/iparam-list-template";
import { IParamTemplate } from "@template/interfaces/iparam-template";
import { TemplateRepository } from "@template/repository/template.repository";

@Injectable()
export class TemplateService {
    constructor(private readonly variableRepository: TemplateRepository) { }

    async create(params: IParamTemplate): Promise<BaseResponse<IParamTemplate>> {
        let Response = new BaseResponse<IParamTemplate>();
        try {
            const result = await this.variableRepository.save(params);
            if (result.success) {
                Response = Res({ result: result.data, statusCode: 200, message: result.message });
            } else {
                Response = Res({ statusCode: 400, message: result.message });
            }

        } catch (error) {
            const { statusCode, message } = errorResponse(error);
            Response = Res({ statusCode, message });
        }
        return Response;
    }

    async update(params: IParamTemplate): Promise<BaseResponse<IParamTemplate>> {
        let Response = new BaseResponse<IParamTemplate>();
        try {
            const exist = await this.variableRepository.findOne(params.id);
            if (exist.success) {
                const param = { ...params, id: exist.data.id }
                const result = await this.variableRepository.save(param);
                if (result.success) {
                    Response = Res({ result: result.data, statusCode: 200, message: result.message });
                } else {
                    Response = Res({ statusCode: 400, message: result.message });
                }
            } else {
                Response = Res({ statusCode: 404, message: exist.message });
            }


        } catch (error) {
            const { statusCode, message } = errorResponse(error);
            Response = Res({ statusCode, message });
        }
        return Response;
    }

    async detail(id: string): Promise<BaseResponse<IParamTemplate>> {
        let Response = new BaseResponse<IParamTemplate>();
        try {
            const result = await this.variableRepository.findOne(id);
            if (result.success) {
                Response = Res({ result: result.data, statusCode: 200, message: result.message });
            } else {
                Response = Res({ statusCode: 404, message: result.message });
            }

        } catch (error) {
            const { statusCode, message } = errorResponse(error);
            Response = Res({ statusCode, message });
        }
        return Response;
    }

    async fetchAll(params: IParamListTemplate): Promise<BaseResponse<IParamTemplate[]>> {
        let Response = new BaseResponse<IParamTemplate[]>();
        try {

            const result = await this.variableRepository.findAll(params);
            if (result.success) {
                Response = Res({ result: result.data, statusCode: 200, message: result.message, page: params.page, limit: params.limit, total: result.total });
            } else {
                Response = Res({ statusCode: 400, message: result.message, page: params.page, limit: params.limit });
            }

        } catch (error) {
            const { statusCode, message } = errorResponse(error);
            Response = Res({ statusCode, message });
        }
        return Response;
    }

    async delete(id: string): Promise<BaseResponse<IParamTemplate>> {
        let Response = new BaseResponse<IParamTemplate>();
        try {

            const result = await this.variableRepository.remove(id);
            if (result.success) {
                Response = Res({ result: result.data, statusCode: 200, message: result.message });
            } else {
                Response = Res({ statusCode: 404, message: result.message });
            }

        } catch (error) {
            const { statusCode, message } = errorResponse(error);
            Response = Res({ statusCode, message });
        }
        return Response;
    }
}

