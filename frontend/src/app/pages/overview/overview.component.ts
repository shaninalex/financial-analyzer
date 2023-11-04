import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { WebsocketService } from 'src/app/services/websocket.service';
import { ITickerAction } from 'src/app/typedefs';

@Component({
    selector: 'app-overview',
    templateUrl: './overview.component.html',
    styleUrls: ['./overview.component.css']
})
export class OverviewComponent {
    report_ready: boolean = false;

    tickerForm: FormGroup = new FormGroup({
        "ticker": new FormControl("", [Validators.required])
    })

    constructor(private socket: WebsocketService) {

    }

    onSubmit(): void {
        if (this.tickerForm.valid) {
            const search_payload: ITickerAction = {
                ticker: this.tickerForm.value.ticker.toUpperCase(),
                action: "search"
            };
            this.socket.send(search_payload);
        }
    }
}
