export interface ILoginBody {
    username: string;
    password: string;
}

export interface ILoginResponse {
    token: string;
    name: string;
    id: string;
}

export interface IRegisterBody {
    username: string;
    password: string;
}

export interface IRegisterResponse {
    token: string;
}