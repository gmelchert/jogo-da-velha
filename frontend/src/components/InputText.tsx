import { HTMLProps, RefObject } from "react";

interface InputText extends HTMLProps<HTMLInputElement> {
    forwardedRef?: RefObject<HTMLInputElement | null>;
}

export const Input = ({ forwardedRef, ...props}: InputText) => {
    const baseClass = `border border-indigo-950 border-b-amber-500 p-2 w-full 
    border-amber-500 bg-indigo-950 rounded-t placeholder-amber-700 text-amber-400 
    focus:border-amber-500 focus:outline-none focus:ring-0`;

    const className = props.className ? `${baseClass} ${props.className}` : baseClass;

    delete props.className;

    return (
        <input
            {...props}
            className={className}
            ref={forwardedRef}
        />
    )
}