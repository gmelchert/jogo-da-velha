import { FormEvent, useState } from "react";
import { useNavigate } from "react-router-dom";
import { login } from "../services";
import { notify } from "../hooks";

export const LoginPage = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const navigate = useNavigate();

    async function handleSubmit(e: FormEvent) {
        e.preventDefault();
        try {
            const { token } = await login({ username, password });
            localStorage.setItem("token", token);
            navigate("/game");
        } catch {
            notify("Nome de usuário ou senha incorretos.").error();
        }
    }

    return (
        <form onSubmit={handleSubmit} className="max-w-sm m-auto mt-10 space-y-4">
            <h1 className="text-xl font-bold">Login</h1>
            <input
                type="text"
                placeholder="Nome de Usuário"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                className="border p-2 w-full"
            />
            <input
                type="password"
                placeholder="Senha"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                className="border p-2 w-full"
            />
            <button type="submit" className="bg-blue-500 text-white p-2 w-full">
                Entrar
            </button>
        </form>
    )
}