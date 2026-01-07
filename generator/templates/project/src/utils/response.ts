import { GlobalMessage } from "@utils/globalMessage";
import { Meta } from "@utils/baseResponse";

export function Res(params?: {
  result?: any;
  total?: any;
  page?: number;
  limit?: number;
  statusCode?: number;
  success?: boolean;
  message?: string;
  meta?: Meta;
}) {
  const {
    statusCode,
    total,
    success,
    message,
    meta,
    result,
    page,
    limit
  } = params;
  const Response = generateResponse({
    statusCode: statusCode ? statusCode : 200,
    success,
    message,
    meta,
    result,
    total,
    page,
    limit
  });
  return Response;
}

export function generateResponse(params: any) {
  let { statusCode, success, message, meta, result } = params;
  const { total, page, limit } = params;
  result = result || undefined;
  statusCode = statusCode ? statusCode : 200;
  success = success ? success : statusCode === 200 ? true : false;
  message = message ? message : responseMessage(statusCode);
  meta = meta || undefined;

  if (total || total === 0) {
    const MetaData = new Meta();
    MetaData.page = page ? Number(page) : 1;
    MetaData.limit = limit ? Number(limit) : 0;
    MetaData.totalDataPage = result.length;
    MetaData.totalData = total;
    MetaData.totalPages = limit ? Math.ceil(total / limit) : 1;
    meta = MetaData;
  }

  const Response = { statusCode, success, message, meta, result };
  Object.keys(Response).forEach(key => Response[key] === undefined && delete Response[key]);
  return Response;
}

export function responseMessage(statusCode: number) {
  let message = null;
  switch (statusCode) {
    case 200:
      message = GlobalMessage.GET_SUCCESS;
      break;
    case 401:
      message = GlobalMessage.NOT_AUTHORIZED;
      break;
    case 400:
      message = GlobalMessage.GET_FAILED;
      break;
    case 404:
      message = GlobalMessage.NO_FOUND;
      break;
    case 500:
      message = GlobalMessage.INTERNAL_ERROR;
      break;
    default:
      break;
  }

  return message;
}
