import { create } from "zustand";

import { IUser } from "@/@types";

interface AuthStoreI {
    user: IUser;
    logged: boolean;
    login: (user: IUser) => void;
    logout: () => void;
}

export const useAuth = create<AuthStoreI>(set => ({
    user: { username: "", id: 0 ,token: "" },
    logged: false,
    login: (user: IUser) => set(() => ({ user, logged: true })),
    logout: () => {
        localStorage.removeItem("token");
        set(() => ({ user: { username: "", id: 0, token: "" }, logged: false }))
    },
}))