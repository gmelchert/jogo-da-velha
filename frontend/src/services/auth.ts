import { api } from "@/utils"

import {
    ILoginBody,
    ILoginResponse,
    IMeResponse,
    IRegisterBody,
    IRegisterResponse,
} from "@/@types";

export class AuthService {
    public static async login(loginBody: ILoginBody) {
        return api<ILoginResponse>('login').post<ILoginBody>(loginBody);
    }

    public static async register(registerBody: IRegisterBody) {
        return api<IRegisterResponse>('register').post<IRegisterBody>(registerBody);
    }

    public static async getAuthenticatedUser() {
        return api<IMeResponse>('api/v1/me').get();
    }
}