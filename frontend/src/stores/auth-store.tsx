import { create } from "zustand";

import { IAuthUser } from "@/@types";

interface AuthStoreI {
    user: IAuthUser;
    logged: boolean;
    login: (user: IAuthUser) => void;
    logout: () => void;
}

export const useAuth = create<AuthStoreI>(set => ({
    user: { username: "", id: 0 ,token: "" },
    logged: false,
    login: (user: IAuthUser) => set(() => ({ user, logged: true })),
    logout: () => {
        localStorage.removeItem("token");
        set(() => ({ user: { username: "", id: 0, token: "" }, logged: false }))
    },
}))