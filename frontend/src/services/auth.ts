import { api } from "../utils"

import {
    ILoginBody,
    ILoginResponse,
    IRegisterBody,
    IRegisterResponse,
} from "../@types";

export const login = async (loginBody: ILoginBody) => {
    return api<ILoginResponse>('login').post<ILoginBody>(loginBody);
}

export const register = async (registerBody: IRegisterBody) => {
    return api<IRegisterResponse>('login').post<IRegisterBody>(registerBody);
}