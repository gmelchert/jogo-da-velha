interface WelcomeMessageProps {
    username: string;
}

export const WelcomeMessage = ({ username }: WelcomeMessageProps) =>
    username ? <>Bem vindo, <span className="text-amber-500">{username.toUpperCase()}</span>!</> : <>Bem vindo!</>
