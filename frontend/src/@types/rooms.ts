import { ROOM_STATUS } from "@/enums";
import { IUser } from "./user";
import { IStats } from "./stats";

interface Owner extends IStats {
    user: Partial<IUser>;
}

export interface IRooms {
    id: string;
    status: ROOM_STATUS;
    owner: Owner;
}