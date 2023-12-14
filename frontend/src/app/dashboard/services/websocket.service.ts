import { Injectable } from "@angular/core";
import { BehaviorSubject } from "rxjs";
import { environment } from "src/environments/environment.development";

@Injectable({
    providedIn: "root"
})
export class WebsocketService {
    private socket = new WebSocket(environment.WEBSOCKET_ADDRESS);
    public messages: BehaviorSubject<any> = new BehaviorSubject("");

    constructor() {
        this.socket.onopen = (event: Event) => {
            console.log(event);
        }

        this.socket.onmessage = (event: Event) => {
            this.messages.next(event);
        }
    }

    public send(message: any) {
        this.socket.send(JSON.stringify(message));
    }
}
