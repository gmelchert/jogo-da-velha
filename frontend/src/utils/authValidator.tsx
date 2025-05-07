import { AuthService } from "@/services";
import { UserI } from "@/@types";

interface IAuthValidatorProps {
    logged: boolean;
    login: (user: UserI) => void;
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
    if (!token) return false;

    try {
        const user = await AuthService.getAuthenticatedUser();
        login({ token, ...user });
        return true;
    } catch {
        (logged) && logout();
        return false;
    }
}