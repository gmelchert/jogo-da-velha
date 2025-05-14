import { v4 as uuidv4 } from 'uuid';

import { api, queryParamBuilder } from "@/utils";
// import { ROOM_STATUS } from '@/enums';
import { API_V1_ROOT } from '@/constants';

import {
    ICreateRoom,
    IFindRoomQuery,
    IFindRoomResponse,
    ICreateRoomResponse,
} from "@/@types";

export class RoomsService {
    public static async findRooms(query: IFindRoomQuery) {
        const queryUrl = queryParamBuilder(query);
        return api<IFindRoomResponse>(`/${API_V1_ROOT}/rooms?${queryUrl}`).get();
    }

    public static async createRoom() {
        const roomId = uuidv4();
        return api<ICreateRoomResponse>(`/${API_V1_ROOT}/rooms`).post<ICreateRoom>({ roomId });
    }
}