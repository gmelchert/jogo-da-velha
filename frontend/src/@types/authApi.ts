export interface ILoginBody {
    username: string;
    password: string;
}

export interface ILoginResponse {
    token: string;
    username: string;
    id: number;
}

export interface IRegisterBody {
    username: string;
    password: string;
}

export interface IRegisterResponse {
    token: string;
    username: string;
    id: number;
}

export interface IMeResponse {
    username: string;
    id: number;
    password: string;
}