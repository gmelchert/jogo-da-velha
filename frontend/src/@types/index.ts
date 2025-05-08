export * from './authApi';
export * from './user';
export * from './rooms';
export * from './stats';

export interface IAPIResponse<D> {
    message: string;
    data: D;
}