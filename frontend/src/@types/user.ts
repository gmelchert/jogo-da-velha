export interface IUser {
    username: string;
    id: number;
    token: string;
    ID: number;
    CreatedAt: Date;
    UpdatedAt: Date | null;
    DeletedAt: Date | null;
    email: string;
    password: string;
    draws: number;
    losses: number;
    wins: number;
    gamesPlayed: number;
}

export interface IAuthUser {
    username: string;
    id: number;
    token: string;
}