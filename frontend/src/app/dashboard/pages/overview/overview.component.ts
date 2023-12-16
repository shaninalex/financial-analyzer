import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Observable } from 'rxjs';
import { WebsocketService } from '../../services/websocket.service';
import { ITickerAction } from '../../typedefs/global';

@Component({
    selector: 'app-overview',
    templateUrl: './overview.component.html',
})
export class OverviewComponent {
    result_data: boolean = false;
    current_date: Date = new Date();
    tickerForm: FormGroup = new FormGroup({
        "ticker": new FormControl("IBM", [Validators.required])
    })
    messageHub$: Observable<any>;
    overview: any;
    cashflow: any;
    earnings: any;

    constructor(private socket: WebsocketService) {
        this.socket.messages.subscribe({
            next: data => {
                if (data) {
                    const res = JSON.parse(data.data);
                    switch(res.type) {
                        case "alph_overview":
                            this.overview = res;
                            break;
                        case "alph_cashflow":
                            this.cashflow = res;
                            break;
                        case "alph_earnings":
                            this.earnings = res;
                            break;
                    }
                }
            }
        });
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
