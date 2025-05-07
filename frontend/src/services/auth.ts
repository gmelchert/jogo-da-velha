import { api } from "@/utils"

import {
    ILoginBody,
    ILoginResponse,
    IMeResponse,
    IRegisterBody,
    IRegisterResponse,
} from "@/@types";

export const loginService = async (loginBody: ILoginBody) =>
    api<ILoginResponse>('login').post<ILoginBody>(loginBody);

export const registerService = async (registerBody: IRegisterBody) =>
    api<IRegisterResponse>('login').post<IRegisterBody>(registerBody);

export const getAuthenticatedUserService = async () =>
    api<IMeResponse>('me').get();