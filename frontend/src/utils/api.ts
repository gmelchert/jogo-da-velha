import { API_URL } from "../constants";

const checkIfEndpointStartsWithBar = (str: string) =>
    str.startsWith("/")
        ? str
        : `/${str}`;

export const api = <R>(endpoint: string) => {
    const url = API_URL + checkIfEndpointStartsWithBar(endpoint)

    const token = localStorage.getItem("token");

    const headers: Headers = new Headers({
        "Content-Type": "application/json",
    })
    
    token && headers.append("Authorization", `Bearer ${token}`);

    const init: RequestInit = {
        headers,
    }

    return {
        async get(): Promise<R> {
            init.method = "GET";
            const res = await fetch(url, init);
            if (!res.ok) throw new Error("Erro ao chamar API.");

            return res.json();
        },
        async post<B>(body?: B): Promise<R> {
            init.method = "POST";
            if (body) init.body = JSON.stringify(body);

            const res = await fetch(url, init);
            if (!res.ok) throw new Error("Erro ao chamar API.");

            return res.json();
        },
        async put<B>(body?: B): Promise<R> {
            init.method = "PUT";
            if (body) init.body = JSON.stringify(body);

            const res = await fetch(url, init);
            if (!res.ok) throw new Error("Erro ao chamar API.");

            return res.json();
        },
        async delete(): Promise<R> {
            init.method = "DELETE";

            const res = await fetch(url, init);
            if (!res.ok) throw new Error("Erro ao chamar API.");

            return res.json();
        },
    }
    
}