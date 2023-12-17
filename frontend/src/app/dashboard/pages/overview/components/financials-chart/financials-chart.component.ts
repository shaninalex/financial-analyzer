import { Component } from "@angular/core";
import { Store } from "@ngrx/store";
import { IReportState } from "src/app/dashboard/store/report/reducer";
import { selectFinancialsChartData } from "src/app/dashboard/store/report/selectors";

@Component({
    selector: "financials-chart",
    host: { "class": "col-md-6" },
    template: `@if(chartdata.values.length) {
        <div class="card mb-4 overflow-y-scroll">
            <div class="card-body">
                <app-line-chart [data]="chartdata" />
            </div>
        </div>
    }`
})
export class FinancialsChartComponent {
    chartdata: {label: string, values: Array<[string, number]>} = {label: "Financials", values: []};

    constructor(private store: Store<IReportState>) {
        this.store.select(selectFinancialsChartData).subscribe({
            next: data => {
                if (data.length) this.chartdata = {...this.chartdata, values: data.map(d => [d.pay_date, parseFloat(d.amount)])};
            }
        });
    }
}