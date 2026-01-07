export interface IParamListTemplate {
    page: number;
    limit: number;
    sort: string;
    orderBy: string;
    query: string;
    noCache?: boolean;
}