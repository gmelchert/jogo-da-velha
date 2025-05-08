import { IAPIResponse } from './'

export interface ILoginBody {
    username: string;
    password: string;
}

export interface ILoginResponse extends IAPIResponse<{
    token: string;
    username: string;
    id: number;
    message: string;
}>{}

export interface IRegisterBody {
    username: string;
    password: string;
    email: string;
}

export interface IRegisterResponse extends IAPIResponse<{
    token: string;
    username: string;
    message: string;
    id: number;
}>{}

export interface IMeResponse extends IAPIResponse<{
    username: string;
    email: string;
    createAt: string;
    updatedAt: string;
    deletedAt: string;
    message: string;
    id: number;
}>{}