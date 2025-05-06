import { create } from "zustand";

interface UserI {
    name: string;
    id: string;
    token: string;
}

interface AuthStoreI {
    user: UserI;
    logged: boolean;
    login: (user: UserI) => void;
    logout: () => void;
}

export const useAuth = create<AuthStoreI>(set => ({
    user: { name: "", id: "" ,token: "" },
    logged: false,
    login: (user: UserI) => set(() => ({ user, logged: true })),
    logout: () => set(() => ({ user: { name: "", id: "", token: "" }, logged: false })),
}))