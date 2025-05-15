import { useNavigate } from "react-router-dom";
import { Button } from "./Button";

import { IRooms } from "@/@types";
import { RoomsService } from "@/services";

interface RoomCardProps {
    room: IRooms;
}

export const RoomCard = ({ room }: RoomCardProps) => {
    const { owner, roomId } = room;

    const navigate = useNavigate();

    const onJoinRoomClick = async () => {
        try {
            RoomsService.joinRoom(roomId);
            navigate(`/game/${roomId}`);
        } catch (error) {
            console.error("Error joining room:", error);
        }
    }

    return (
        <article
            className="bg-violet-500 rounded shadow-md shadow-amber-400 p-4"
        >
            <h2 className="text-violet-100 text-center text-2xl font-bold flex justify-between">
                <span>Jogador:</span>
                <span className="text-amber-300">{owner.username?.toUpperCase()}</span>
            </h2>

            <hr className="border-violet-200 my-2" />

            <h3
                className="text-violet-100 text-center text-xl font-bold"
            >Estatísticas do Jogador</h3>

            <section className="grid grid-cols-4 gap-x-2 text-violet-100 text-center">
                <span>Jogos:</span>
                <span>Vitórias:</span>
                <span>Derrotas:</span>
                <span>Empates:</span>
                
                <span className="text-amber-300">{owner.gamesPlayed}</span>
                <span className="text-amber-300">{owner.wins}</span>
                <span className="text-amber-300">{owner.losses}</span>
                <span className="text-amber-300">{owner.draws}</span>
            </section>

            <Button className="w-full mt-4" onClick={onJoinRoomClick}>Entrar na Sala</Button>
        </article>
    )
}