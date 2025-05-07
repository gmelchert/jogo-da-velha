import { useLocation } from "react-router-dom";
import { useAuth } from "@/stores"

export const Header = () => {
    const { logged, user, logout } = useAuth();
    const location = useLocation();

    const listItemClassNameBuilder = (path: string) => 
        path === location.pathname
            ? 'border-b text-amber-500'
            : 'border-b border-indigo-950 hover:border-amber-400 transition-all ease-linear hover:text-amber-400';

    return (
        <header
            className="h-20 bg-indigo-950 sticky top-0 text-white
            flex items-center justify-between px-8"
        >
            <h1 className="text-2xl font-bold font-mono cursor-pointer">
                <a href="/">Jogo da Velha!</a>
            </h1>

            <section className="flex items-center">
                <ul className="flex gap-4 border-r-2 pr-4">
                    <li>
                        <a
                            href="/"
                            className={listItemClassNameBuilder('/')}
                        >Home</a>
                    </li>
                    <li>
                        <a
                            href="/rooms"
                            className={listItemClassNameBuilder('/rooms')}
                        >Salas</a>
                    </li>
                    <li>
                        <a
                            href="/stats"
                            className={listItemClassNameBuilder('/stats')}
                        >Estatísticas</a>
                    </li>
                    {logged && <li>
                        <a
                            href={`/history/${user.id}`}
                            className={listItemClassNameBuilder(`/history/${user.id}`)}
                        >Histórico</a>
                    </li>}
                </ul>

                <div className="ml-4 flex gap-4 items-center">
                    { logged ? (
                        <button
                            onClick={logout}
                            className="bg-amber-500 text-white rounded-md
                            py-1 w-20 hover:bg-amber-600 text-center
                            transition-colors ease-linear cursor-pointer"
                        >
                            Logout
                        </button>
                    ) : (<>
                        <a
                            href="/register"
                            className="border border-amber-500 text-amber-500 rounded-md
                            py-1 w-20 hover:bg-amber-500 hover:text-white
                            transition-colors ease-linear text-center"
                        >
                            Registrar
                        </a>
                        <a
                            href="/login"
                            className="bg-amber-500 text-white rounded-md
                            py-1 w-20 hover:bg-amber-600 text-center
                            transition-colors ease-linear"
                        >
                            Login
                        </a>    
                    </>)}
                    
                </div>
            </section>
        </header>
    )
}