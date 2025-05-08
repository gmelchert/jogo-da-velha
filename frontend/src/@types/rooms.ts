import { ROOM_STATUS } from "@/enums";
import { IUser } from "./user";
import { IStats } from "./stats";

interface Owner extends IStats {
    user: Partial<IUser>;
}

export interface IRooms {
    id: string;
    owner: Owner;
    roomId: string;
    status: ROOM_STATUS;
}