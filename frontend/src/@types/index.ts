export * from './auth';
export * from './user';
export * from './rooms';
export * from './stats';

export interface IAPIResponse<D> {
    message: string;
    data: D;
}

export interface IAPIPaginatedResponse<D> {
    message: string;
    data: {
        rows: D[];
        page: number;
        limit: number;
        totalRows: number;
        totalPages: number;
    };
}