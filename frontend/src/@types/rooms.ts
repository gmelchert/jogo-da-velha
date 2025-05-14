import { IAPIPaginatedResponse, IAPIResponse } from ".";
import { IUser } from "./user";
import { IStats } from "./stats";

import { ROOM_STATUS } from "@/enums";

interface PartialUser extends IStats {
    user: Partial<IUser>;
}

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

export interface ICreateRoom {
    roomId: string;
}

export interface IFindRoomResponse extends IAPIPaginatedResponse<IRooms>{}

export interface ICreateRoomResponse extends IAPIResponse<IRooms>{}

/*{
    RoomID     string `json:"roomId"`
    OwnerID    uint   `json:"userId"`
    Owner      User   `json:"owner" gorm:"foreingKey:UserID"`
    OpponentID uint   `json:"opponentID"`
    Opponent   User   `json:"opponent" gorm:"foreingKey:OpponentID"`
    Status     string `json:"status"`
    ID        uint           // through Model 
    CreatedAt time.Time      // through Model 
    UpdatedAt time.Time      // through Model 
    DeletedAt gorm.DeletedAt // through Model
}*/