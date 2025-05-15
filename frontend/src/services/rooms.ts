import { v4 as uuidv4 } from 'uuid';

import { api, queryParamBuilder } from "@/utils";
// import { ROOM_STATUS } from '@/enums';
import { API_V1_ROOT } from '@/constants';

import {
    ICreateRoom,
    IFindRoomQuery,
    IFindRoomResponse,
    ICreateRoomResponse,
    IJoinRoomResponse,
} from "@/@types";

export class RoomsService {
    private static readonly root = `/${API_V1_ROOT}/rooms`;

    public static async findRooms(query: IFindRoomQuery) {
        const queryUrl = queryParamBuilder(query);
        return api<IFindRoomResponse>(`${this.root}?${queryUrl}`).get();
    }

    public static async createRoom() {
        const roomId = uuidv4();
        return api<ICreateRoomResponse>(this.root).post<ICreateRoom>({ roomId });
    }

    public static async joinRoom(roomId: string) {
        return api<IJoinRoomResponse>(`${this.root}/${roomId}/join`).post();
    }
}