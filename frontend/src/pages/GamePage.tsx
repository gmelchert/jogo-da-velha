import { useEffect, useState } from "react";
import { v4 as uuidv4 } from 'uuid';

import { connectToChannel } from "../websocket";
import { notify } from "../hooks";
import { Board } from "../components";

interface Message {
    type: string;
    data: any;
}

export const GamePage = () => {
    const [ws, setWs] = useState<WebSocket | null>(null);
    const [board, setBoard] = useState<number[]>(Array(9).fill(0));
    const [turn, setTurn] = useState(1);
    const [player, setPlayer] = useState(1);

    const token = localStorage.getItem("token")!
    const channelID = uuidv4();

    useEffect(() => {
        const socket = connectToChannel(channelID, token);

        socket.onmessage = event => {
            const msg: Message = JSON.parse(event.data);

            if (msg.type === "start") {
                notify('Jogo começou!').success();
                setTurn(msg.data.turn);
            }

            if (msg.type === "play") {
                setBoard(prev => {
                    const newBoard = [...prev];
                    newBoard[msg.data.position] = msg.data.player;
                    return newBoard;
                })
                setTurn(prev => (prev === 1 ? 2 : 1));
            }

            if (msg.type === "win") {
                notify(`Jogador ${msg.data} venceu!`).success();
            }

            if (msg.type === "draw") {
                notify("Empate!").success();
            }

            if (msg.type === "invalid") {
                notify("Jogada inválida").warn();
            }
        }

        setWs(socket);

        return () => socket.close();
    }, []);

    const sendPlay = (pos: number) => {
        if (turn !== player) return;

        ws?.send(JSON.stringify({ type: "play", data: { position: pos } }));
    }

    return (
        <div className="text-center">
            <h1 className="text-xl font-bold mt-4">Jogo da Velha</h1>
            <p>Você é o jogador {player === 1 ? "X" : "O"}</p>
            <Board board={board} onPlay={sendPlay} />
        </div>
    )
}