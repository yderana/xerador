export class BaseResponse<T> {
  result: T;
  statusCode: number;
  success: boolean;
  message: string;
  meta: Meta;
}

// tslint:disable-next-line:max-classes-per-file
export class Meta {
  page: number;
  limit: number;
  totalDataPage: number;
  totalData: number;
  totalPages: number;
}

// tslint:disable-next-line:max-classes-per-file
export class BaseResponseES<T, Agg> {
  result: T;
  statusCode: number;
  success: boolean;
  message: string;
  meta: Meta;
  aggregation?: Agg;
}
