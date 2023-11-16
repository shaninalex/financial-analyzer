import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { BehaviorSubject, Observable, map, of } from 'rxjs';
import { WebsocketService } from 'src/app/services/websocket.service';
import { ITickerAction } from 'src/app/typedefs';

@Component({
    selector: 'app-overview',
    templateUrl: './overview.component.html',
    styleUrls: ['./overview.component.css']
})
export class OverviewComponent {
    result_data: boolean = false;
    current_date: Date = new Date();
    tickerForm: FormGroup = new FormGroup({
        "ticker": new FormControl("IBM", [Validators.required])
    })
    messageHub$: Observable<any>;
    overview$: Observable<any>;
    cashflow$: Observable<any>;
    earnings$: Observable<any>;

    constructor(private socket: WebsocketService) {
        this.messageHub$ = this.socket.messages;
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

// action
// type
// data