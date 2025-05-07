import { create } from "zustand";

import { UserI } from "@/@types";

interface AuthStoreI {
    user: UserI;
    logged: boolean;
    login: (user: UserI) => void;
    logout: () => void;
}

export const useAuth = create<AuthStoreI>(set => ({
    user: { username: "", id: "" ,token: "" },
    logged: false,
    login: (user: UserI) => set(() => ({ user, logged: true })),
    logout: () => {
        localStorage.removeItem("token");
        set(() => ({ user: { username: "", id: "", token: "" }, logged: false }))
    },
}))