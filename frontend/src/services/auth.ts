import { api } from "../utils"

import {
    ILoginBody,
    ILoginResponse,
    IRegisterBody,
    IRegisterResponse,
} from "../@types";

export const loginService = async (loginBody: ILoginBody) => {
    return api<ILoginResponse>('login').post<ILoginBody>(loginBody);
}

export const registerService = async (registerBody: IRegisterBody) => {
    return api<IRegisterResponse>('login').post<IRegisterBody>(registerBody);
}