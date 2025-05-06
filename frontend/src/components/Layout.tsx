import { ReactNode } from "react"
import { Header } from "./Header"

interface LayoutProps {
    children: ReactNode
}

export const Layout = ({ children }: LayoutProps) => {
    return (
        <main className="flex flex-col h-screen bg-violet-950">
            <Header />

            <div>
                {children}
            </div>
        </main>
    )
}