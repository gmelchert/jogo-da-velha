import { v4 as uuidv4 } from 'uuid';

import { IRooms } from "@/@types";
import { ROOM_STATUS } from '@/enums';
// import { api } from "@/utils";

export class RoomsService {
    public static async getAllRooms() {
        const rooms: IRooms[] = [
            {
                id: uuidv4(),
                status: ROOM_STATUS.OPEN,
                owner: {
                    draws: 12,
                    gamesPlayed: 26,
                    losses: 4,
                    wins: 10,
                    userId: 1,
                    user: {
                        id: "1",
                        username: 'Fulano',
                    }
                }
            }
        ];
        return rooms;
        // return api<IRooms[]>('').get();
    }
}