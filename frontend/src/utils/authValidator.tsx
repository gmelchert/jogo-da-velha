import { AuthService } from "@/services";
import { IUser } from "@/@types";

interface IAuthValidatorProps {
    logged: boolean;
    login: (user: IUser) => void;
    logout: () => void;
}

// checks if there's a JWT in local storage and if it's still valid
// if so login user, if not logout
export const authValidator = async ({
    login,
    logout,
    logged,
}: IAuthValidatorProps): Promise<boolean> => {
    const token = localStorage.getItem("token");
    if (logged || !token) return false;

    try {
        const { data } = await AuthService.getAuthenticatedUser();
        login({
            id: data.id,
            token: token,
            username: data.username,
        })
        return true;
    } catch {
        (logged) && logout();
        return false;
    }
}