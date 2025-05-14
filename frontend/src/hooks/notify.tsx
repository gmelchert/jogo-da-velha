import { Bounce, toast, ToastOptions } from "react-toastify";

export const notify = (
    message: string,
    config?: ToastOptions<unknown>,
) => {
    const defaultToastCOnfig: ToastOptions<unknown> = {
        position: "top-right",
        autoClose: 5000,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: false,
        progress: undefined,
        theme: "dark",
        transition: Bounce,
        ...config,
    }

    return {
        success: () => toast.success(message, defaultToastCOnfig),
        warn: () => toast.warn(message, defaultToastCOnfig),
        error: () => toast.error(message, defaultToastCOnfig),
        default: () => toast(message, defaultToastCOnfig),
    }
}