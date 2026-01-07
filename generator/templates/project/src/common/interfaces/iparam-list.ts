export interface IParamList {
  page: number;
  limit: number;
  sort: string;
  orderBy: string;
  query: string;
  noCache?: boolean;
}
