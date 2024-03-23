import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { WebsocketService } from '../../services/websocket.service';
import { ActionTypeReport, ISocketAction } from '../../typedefs/global';
import { Store } from '@ngrx/store';
import { IReportState } from '../../store/report/reducer';
import { setFinancialsChartData, setPriceChartData } from '../../store/report/actions';


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
  dividend: any;
  keyratios: any;

  is_price_chart_loaded: boolean;
  is_financials_chart_loaded: boolean;

  constructor(private socket: WebsocketService, private store: Store<IReportState>) {
    this.socket.messages.subscribe({
      next: payload => {
        switch (payload?.type) {
          case "summary":
            this.summary = payload.data;
            break;
          case "financials":
            this.store.dispatch(setFinancialsChartData({ data: payload.data.reverse() }));
            this.is_financials_chart_loaded = true;
            break;
          case "dividend":
            this.dividend = payload.data;
            break;
          case "price":
            this.store.dispatch(setPriceChartData({ data: payload.data }));
            this.is_price_chart_loaded = true;
            break;
          case "keyratios":
            this.keyratios = payload.data;
            break;
        }
      }
    });
  }


  onSubmit(): void {
    if (this.tickerForm.valid) {
      const search_payload: ISocketAction = {
        action: ActionTypeReport,
        payload: {
          ticker: this.tickerForm.value.ticker.toUpperCase(),
        }
      };
      this.socket.send(search_payload);
      this.is_price_chart_loaded = false;
      this.is_financials_chart_loaded = false;
    }
  }
}
