import { Component } from '@angular/core';
import { WebsocketService } from './services/websocket.service';
import { ITickerAction } from './typedefs';

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.css']
})
export class AppComponent {
    constructor(
        // private socket: WebsocketService
    ) {}

    test_send(): void {
        // const search_payload: ITickerAction = {
        //     ticker: "AAPL",
        //     action: "search"
        // };
        // this.socket.send(search_payload);
    }
}
