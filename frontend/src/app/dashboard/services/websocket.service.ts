import { Injectable } from "@angular/core";
import { BehaviorSubject } from "rxjs";

@Injectable({
    providedIn: "root"
})
export class WebsocketService {
    private socket = new WebSocket("ws://127.0.0.1:8080/ws");
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
