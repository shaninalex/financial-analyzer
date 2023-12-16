import { Injectable } from "@angular/core";
import { BehaviorSubject } from "rxjs";
import { environment } from "src/environments/environment.development";
import { INotification, IResponseData } from "../typedefs/global";

@Injectable()
export class WebsocketService {
    private socket = new WebSocket(environment.WEBSOCKET_ADDRESS);
    public messages: BehaviorSubject<IResponseData | null> = new BehaviorSubject<IResponseData | null>(null);
    public notifications: BehaviorSubject<INotification | null> = new BehaviorSubject<INotification | null>(null);

    constructor() {
        this.socket.onopen = (event: Event) => {
            console.log(event);
        }

        this.socket.onmessage = (event: MessageEvent) => {
            const m = JSON.parse(event.data);
            switch(m.action) {
                case "notification":
                    this.notifications.next(m.payload);
                    break
                case "data_result":
                    this.messages.next(m.payload);
                    break
            }
        }

        this.socket.onerror = (event: Event) => {
            console.log(event);
        }
    }

    public send(message: any) {
        this.socket.send(JSON.stringify(message));
    }
}
