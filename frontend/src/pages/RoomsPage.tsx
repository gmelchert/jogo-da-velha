import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

import { notify } from "@/hooks";
import { useAuth } from "@/stores";
import { ROOM_STATUS } from "@/enums";
import { authValidator } from "@/utils";
import { RoomsService } from "@/services";
import { Button, Layout, Loading } from "@/components";

import { IRooms } from "@/@types";

export const RoomsPage = () => {
    const [rooms, setRooms] = useState<IRooms[]>([]);
    const [loading, setLoading] = useState<boolean>(false);
    const [page, setPage] = useState<number>(1);

    const navigate = useNavigate();
    const { logged, login, logout } = useAuth();

    const fetchRooms = async (page: number) => {
        setLoading(true);
        try {
            const { data } = await RoomsService.findRooms({
                status: ROOM_STATUS.OPEN,
                limit: 20,
                page,
            });
            setRooms(data.rows);
        } catch (error) {
            console.error("Error fetching rooms:", error);
        } finally {
            setLoading(false);
        }
    }

    const onCreateRoomClick = async () => {
        try {
            const { data } = await RoomsService.createRoom();
            navigate(`/game/${data.roomId}`);
        } catch (error) {
            notify("Erro ao criar sala.").error();
        }
    }

    useEffect(() => {
        fetchRooms(page);
    }, [page]);
    
    useEffect(() => {
        authValidator({ logged, login, logout });
    }, []);

    return (
        <Layout>
            <section className="relative">
                <h1
                    className="text-4xl text-white w-fit
                    mx-auto border-b-2 border-amber-500 pb-2"
                >
                    Salas disponíveis: <span className="text-amber-500">{rooms.length}</span>
                </h1>
                
                <Button
                    className="absolute top-1.5 right-4 w-32 py-2"
                    onClick={onCreateRoomClick}
                >Criar Sala</Button>
            </section>

            <section
                className="grid-cols-3"
            >

            </section>

            <Loading loading={loading} isBlocker={true} text="Carregando Salas..." />
        </Layout>
    )
}