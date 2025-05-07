import { ButtonHTMLAttributes, DetailedHTMLProps } from "react"

export const Button = (props: DetailedHTMLProps<ButtonHTMLAttributes<HTMLButtonElement>,HTMLButtonElement>) => {
    const baseClass = `bg-amber-500 text-white rounded-md
    py-1 px-2 hover:bg-amber-600 text-center
    transition-colors ease-linear cursor-pointer`;

    const className = props.className ? `${baseClass} ${props.className}` : baseClass;

    return (
        <button
            {...props}
            className={className}
        >{props.children}</button>
    )
}
