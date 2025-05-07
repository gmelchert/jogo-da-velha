import { useEffect } from "react";

import { Layout, WelcomeMessage } from "@/components";
import { authValidator } from "@/utils";
import { useAuth } from "@/stores";

export const HomePage = () => {
    const { user, logged, login, logout } = useAuth();

    useEffect(() => {
        authValidator({ logged, login, logout });
    }, []);

    return (
        <Layout>
            <h1
                className="mx-auto text-white text-5xl mt-4 border-b-2 w-fit pb-2
                border-amber-500"
            >
                <WelcomeMessage username={user.username} />
            </h1>
        </Layout>
    )
}