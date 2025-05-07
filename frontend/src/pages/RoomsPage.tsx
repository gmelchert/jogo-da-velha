import { useState } from "react";

import { Button, Layout } from "@/components";
import { IRooms } from "@/@types";

export const RoomsPage = () => {
    const [rooms, setRooms] = useState<IRooms[]>([]);

    return (
        <Layout>
            <section className="relative">
                <h1
                    className="text-4xl text-white w-fit
                    mx-auto border-b-2 border-amber-500 pb-2"
                >
                    Salas disponíveis: <span className="text-amber-500">{rooms.length}</span>
                </h1>
                
                <Button className="absolute top-1.5 right-4 w-32 py-2">Criar Sala</Button>
            </section>

            <section
                className="grid-cols-3"
            >

            </section>
        </Layout>
    )
}