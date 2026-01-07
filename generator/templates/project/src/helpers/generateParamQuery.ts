import { IParamList } from "@common/interfaces/iparam-list";
import { Like } from "typeorm";

export function generateParamQuery(params: IParamList, columns: string[]) {
    const { page, limit, sort, orderBy, query } = params;
    let filter = [];

    if (query) {
        const qString = `%${query}%`;
        const filterCol: any[] = columns.map((column: string) => {
            const filterCol = Object.assign(JSON.parse(`{"${column}": null}`));
            Object.keys(filterCol).map((key: any) => filterCol[key] = Like(qString));
            return filterCol;
        });
        filter = filterCol;
    }

    let limitData = {};
    if (limit) {
        limitData = { take: limit, skip: page > 1 ? (page - 1) * limit : 0 };
    }

    let orderData = null;
    if (orderBy) {
        const orderStr = `{"${orderBy}": "${sort ? sort.toUpperCase() : "ASC"}"}`;
        orderData = JSON.parse(orderStr);
    }

    return { filter, limitData, orderData }
}