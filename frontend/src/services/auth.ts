import { api } from "@/utils"
import { API_V1_ROOT } from "@/constants";

import {
    ILoginBody,
    ILoginResponse,
    IMeResponse,
    IRegisterBody,
    IRegisterResponse,
} from "@/@types";

export class AuthService {
    public static async login(loginBody: ILoginBody) {
        return api<ILoginResponse>('/login').post<ILoginBody>(loginBody);
    }

    public static async register(registerBody: IRegisterBody) {
        return api<IRegisterResponse>('/register').post<IRegisterBody>(registerBody);
    }

    public static async getAuthenticatedUser() {
        return api<IMeResponse>(`/${API_V1_ROOT}/me`).get();
    }
}