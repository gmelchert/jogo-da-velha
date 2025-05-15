import { useEffect, useState } from "react";
import { v4 as uuidv4 } from 'uuid';

import { connectToChannel } from "@/websocket";
import { notify } from "@/hooks";
import { Board, Layout } from "@/components";

interface Message {
    type: string;
    data: any;
}

export const GamePage = () => {
    

    return (
        <Layout>
            <h1 className="text-2xl text-white">Que comecem os jogos</h1>
        </Layout>
    )
}