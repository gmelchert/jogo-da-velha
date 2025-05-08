import { FormEvent, useRef } from "react";
import { useNavigate } from "react-router-dom";

import { AuthService } from "@/services";
import { notify } from "@/hooks";
import { useAuth } from "@/stores";

import { Input, Layout } from "@/components";

export const RegisterPage = () => {
    const { login } = useAuth();
    
    const usernameRef = useRef<HTMLInputElement>(null);
    const emailRef = useRef<HTMLInputElement>(null);
    const passwordRef = useRef<HTMLInputElement>(null);
    const confirmPasswordRef = useRef<HTMLInputElement>(null);
    const navigate = useNavigate();

    async function handleSubmit(e: FormEvent) {
        e.preventDefault();
        
        const username = usernameRef.current!.value;
        const email = emailRef.current!.value;
        const password = passwordRef.current!.value;
        const confirmPassword = confirmPasswordRef.current!.value;

        if (password !== confirmPassword) return notify('As senhas não são iguais.').error();

        try {
            const { data } = await AuthService.register({ username, password, email });
            login({
                id: data.id,
                token: data.token,
                username: data.username
            });
            
            localStorage.setItem("token", data.token);
            navigate("/");
        } catch {
            notify("Erro ao cadastrar.").error();
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
                <h1 className="text-3xl font-bold text-white text-center border-b pb-2">Registrar</h1>
                <Input
                    type="text"
                    placeholder="Nome de Usuário"
                    forwardedRef={usernameRef}
                />
                <Input
                    type="email"
                    placeholder="Email"
                    forwardedRef={emailRef}
                />
                <Input
                    type="password"
                    placeholder="Senha"
                    forwardedRef={passwordRef}
                />
                <Input
                    type="password"
                    placeholder="Confirme sua senha"
                    forwardedRef={confirmPasswordRef}
                />
                <button
                    type="submit"
                    className="bg-amber-500 text-white p-2 w-full
                    cursor-pointer rounded hover:bg-amber-600
                    transition-colors ease-linear"
                >
                    Cadastrar
                </button>
            </form>
        </Layout>
    )
}