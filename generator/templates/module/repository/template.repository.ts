import { Injectable } from "@nestjs/common";
import { getConnection } from "typeorm";
import { GlobalMessage } from "@utils/globalMessage";
import { decodeHash, encodeHash } from "@helpers/hashids";
import { DateTimeNow } from "@helpers/formatDate";
import { errorResponse } from "@helpers/generateErrorResponse";
import { generateParamQuery } from "@helpers/generateParamQuery";
import { Template } from "@template/repository/template.entity";
import { IParamTemplate } from "@template/interfaces/iparam-template";
import { IParamListTemplate } from "@template/interfaces/iparam-list-template";
import { flushRedis, getRedis, setRedis } from "@helpers/redisCache";

@Injectable()
export class TemplateRepository {

    async getOne(option: any) {
        const link = await getConnection().getRepository(Template).findOne({ ...option });
        return link;
    }

    async getListAndCount(option: any) {
        const [result, total] = await getConnection().getRepository(Template).findAndCount({
            ...option
        });

        return { result, total };
    }

    async saveToDB(field: any) {
        const save = await getConnection().getRepository(Template).save({
            ...field
        });

        return save;
    }

    async updateToDB(id: any, data: any) {
        await getConnection().getRepository(Template).update(
            { id },
            {
                ...data
            }
        );
    }

    async removeFromDB(data: any) {
        await getConnection().getRepository(Template).remove(data);
        return data;
    }

    async save(
        input: IParamTemplate,
    ): Promise<{ data?: IParamTemplate, success: boolean, message: string }> {
        let Response: { data?: IParamTemplate, success: boolean, message: string };
        try {
            const id = decodeHash(input.id) || null;
            const exist = await this.getOne({ id });
            const dateNow = DateTimeNow();
            const user = "system";
            if (exist) {
                await this.updateToDB(exist.id, {
                    ...input,
                    id: exist.id,
                    createdAt: exist.createdAt,
                    createdBy: exist.createdBy,
                    modifiedAt: dateNow,
                    modifiedBy: user
                });

                Response = {
                    data: { ...exist, ...input, id: encodeHash(exist.id) },
                    success: true,
                    message: GlobalMessage.SAVE_SUCCESS
                }

            } else {
                const save: any = await this.saveToDB({
                    ...input,
                    createdBy: user,
                    createdAt: dateNow,
                    modifiedBy: user,
                    modifiedAt: dateNow
                });

                Response = {
                    data: { ...save, id: encodeHash(save.id) },
                    success: true,
                    message: GlobalMessage.SAVE_SUCCESS
                }
            }
            flushRedis();
        } catch (error) {
            /* istanbul ignore next */
            const { message } = errorResponse(error);
            Response = {
                success: false,
                message
            }
        }
        return Response;
    }

    async findOne(id: string): Promise<{ data?: IParamTemplate, success: boolean, message: string }> {
        let Response: { data?: IParamTemplate, success: boolean, message: string };
        try {
            const redisKey = `getTemplate_${id}`;
            const redis = await getRedis(redisKey);

            if (redis) {
                Response = redis;
            } else {
                const idValue = decodeHash(id) || null;
                const result = await this.getOne({
                    where: { id: idValue }
                });

                const data = { ...result }

                if (result && id) {
                    Response = { data: { ...data, id: encodeHash(data.id) }, success: true, message: GlobalMessage.GET_SUCCESS };
                    setRedis(redisKey, Response, false);
                } else {
                    Response = { success: false, message: GlobalMessage.NO_FOUND };
                }
            }
        } catch (error) {
            /* istanbul ignore next */
            const { message } = errorResponse(error);
            /* istanbul ignore next */
            Response = { success: false, message };
        }

        return Response;
    }

    async findAll(params: IParamListTemplate): Promise<{ data?: IParamTemplate[], success: boolean, message: string, total: number }> {
        let Response: { data?: IParamTemplate[], success: boolean, message: string, total: number };
        const { limitData, filter, orderData } = generateParamQuery(params, ["name"]);
        try {
            const redisKey = `getTemplates_${JSON.stringify({
                ...limitData,
                ...filter,
                ...orderData
            })}`;
            const redis = await getRedis(redisKey);

            if (redis) {
                Response = redis;
            } else {
                let { result, total } = await this.getListAndCount({
                    ...limitData,
                    where: filter,
                    order: { ...orderData }
                });

                const data: IParamTemplate[] = result.map((item: Template) => {
                    return Object.assign({
                        ...item,
                        id: encodeHash(item.id)
                    })
                });

                Response = { data, success: true, message: GlobalMessage.GET_SUCCESS, total };
                
                if (result.length > 0) {
                    setRedis(redisKey, Response, false);
                }
                
            }
        } catch (error) {
            const { message } = errorResponse(error);
            Response = {
                success: false,
                message,
                total: 0
            };
        }

        return Response;
    }

    async remove(
        id: string,
    ): Promise<{ data?: IParamTemplate, success: boolean, message: string }> {
        let Response: { data?: IParamTemplate, success: boolean, message: string };
        try {
            const idValue = decodeHash(id) || null;
            const exist = await this.getOne({ id: idValue });
            if (exist) {
                await this.removeFromDB(exist);

                Response = {
                    data: { ...exist, id: encodeHash(exist.id) },
                    success: true,
                    message: GlobalMessage.REMOVE_SUCCESS
                }

                flushRedis();

            } else {
                Response = { success: false, message: GlobalMessage.NO_FOUND };
            }
        } catch (error) {
            /* istanbul ignore next */
            const { message } = errorResponse(error);
            Response = {
                success: false,
                message
            }
        }
        return Response;
    }
}
