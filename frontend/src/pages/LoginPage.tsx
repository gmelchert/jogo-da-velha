import { FormEvent, useRef } from "react";
import { useNavigate } from "react-router-dom";

import { AuthService } from "@/services";
import { notify } from "@/hooks";
import { useAuth } from "@/stores";

import { Input, Layout } from "@/components";

export const LoginPage = () => {
    const { login } = useAuth();

    const usernameRef = useRef<HTMLInputElement>(null);
    const passwordRef = useRef<HTMLInputElement>(null);
    const navigate = useNavigate();

    async function handleSubmit(e: FormEvent) {
        e.preventDefault();

        const username = usernameRef.current!.value;
        const password = passwordRef.current!.value;

        try {
            const loginResponse = await AuthService.login({ username, password });
            login(loginResponse);
            localStorage.setItem("token", loginResponse.token);
            navigate("/");
        } catch(err) {
            notify("Nome de usuário ou senha incorretos.").error();
        }
    }

    return (
        <Layout>
            <form
                onSubmit={handleSubmit}
                className="max-w-xl m-auto mt-6 space-y-4
                border-2 rounded-lg px-20 py-12 shadow-lg
                shadow-indigo-900 backdrop-blur-lg"
            >
                <h1 className="text-3xl font-bold text-white text-center border-b pb-2">Login</h1>
                <Input
                    type="text"
                    placeholder="Nome de Usuário"
                    forwardedRef={usernameRef}
                />
                <Input
                    type="password"
                    placeholder="Senha"
                    forwardedRef={passwordRef}
                />
                <button
                    type="submit"
                    className="bg-amber-500 text-white p-2 w-full
                    cursor-pointer rounded hover:bg-amber-600
                    transition-colors ease-linear"
                >
                    Entrar
                </button>
            </form>
        </Layout>
    )
}