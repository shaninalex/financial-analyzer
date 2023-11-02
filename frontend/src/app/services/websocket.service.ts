import { Injectable } from "@angular/core";

@Injectable({
    providedIn: "root"
})
export class WebsocketService {
    socket = new WebSocket("ws://127.0.0.1:8080/ws");

    constructor() {
        this.socket.onopen = (event: Event) => {
            console.log(event);
        }

        this.socket.onmessage = (event: Event) => {
            console.log(event);
        }
    }

    public send(message: any) {
        this.socket.send(JSON.stringify(message));
    }
}
