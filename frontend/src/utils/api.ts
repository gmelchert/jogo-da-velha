import { useNavigate } from "react-router-dom";
import { API_URL } from "@/constants";

const errorMessageHandling = (msg: string) => {
    const navigate = useNavigate();
    if (msg === 'Token não fornecido' || msg === 'Token inválido') {
        navigate('/login');
        throw 'Usuário não autenticado. Por favor faça o login.';
    }
    throw msg;
}

export const api = <R>(endpoint: string) => {
    const url = API_URL + endpoint

    const token = localStorage.getItem("token");

    const headers: Headers = new Headers({
        "Content-Type": "application/json",
    })
    
    token && headers.append("Authorization", `Bearer ${token}`);
    // token && headers.append("credentials", "include");

    const init: RequestInit = {
        headers,
    }

    return {
        async get(): Promise<R> {
            init.method = "GET";
            const res = await fetch(url, init);

            if (!res.ok) {
                const msg = await res.text();
                return errorMessageHandling(msg);
            }

            return res.json();
        },
        async post<B>(body?: B): Promise<R> {
            init.method = "POST";
            if (body) init.body = JSON.stringify(body);

            const res = await fetch(url, init);
            if (!res.ok) {
                const msg = await res.text();
                return errorMessageHandling(msg);
            }

            return res.json();
        },
        async put<B>(body?: B): Promise<R> {
            init.method = "PUT";
            if (body) init.body = JSON.stringify(body);

            const res = await fetch(url, init);
            if (!res.ok) {
                const msg = await res.text();
                return errorMessageHandling(msg);
            }

            return res.json();
        },
        async delete(): Promise<R> {
            init.method = "DELETE";

            const res = await fetch(url, init);
            if (!res.ok) {
                const msg = await res.text();
                return errorMessageHandling(msg);
            }

            return res.json();
        },
    }
    
}