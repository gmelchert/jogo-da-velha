import { IAPIPaginatedResponse, IAPIResponse } from ".";
import { IUser } from "./user";

import { ROOM_STATUS } from "@/enums";

type PartialUser = Partial<IUser>;

export interface IRooms {
    roomId: string;
    ownerID: number;
    owner: PartialUser;
    opponentID: number;
    opponent: PartialUser;
    status: ROOM_STATUS;
    id: number;
    createdAt: Date;
    updatedAt: Date;
    deletedAt: Date;
}

export interface IFindRoomQuery {
    roomId?: string;
    ownerId?: number;
    opponentID?: number;
    page?: number;
    limit?: number;
    status?: ROOM_STATUS;
}

export interface ICreateRoom { roomId: string; }

export interface IFindRoomResponse extends IAPIPaginatedResponse<IRooms>{}

export interface ICreateRoomResponse extends IAPIResponse<IRooms>{}

export interface IJoinRoomResponse { message: string; }