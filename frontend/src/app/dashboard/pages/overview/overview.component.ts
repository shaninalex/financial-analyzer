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
    summary: any;
    financials: any;
    dividend: any;
    price: any;
    keyratios: any;

    constructor(private socket: WebsocketService) {
        this.socket.messages.subscribe({
            next: payload => {
                console.log(payload);
                if (payload) {
                    switch(payload.type) {
                        case "summary":
                            this.summary = payload.data;
                            break;
                        case "financials":
                            this.financials = payload.data;
                            break;
                        case "dividend":
                            this.dividend = payload.data;
                            break;
                        case "price":
                            this.price = payload.data;
                            break;
                        case "keyratios":
                            this.keyratios = payload.data;
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
