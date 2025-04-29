let socket: WebSocket | null = null;

export function connectToChannel(channelID: string, token: string): WebSocket {
    const ws = new WebSocket(`ws://localhost:8080/ws/${channelID}?token=${token}`);
    socket = ws;
    return ws;
}

export const getSocket = (): WebSocket | null => socket;