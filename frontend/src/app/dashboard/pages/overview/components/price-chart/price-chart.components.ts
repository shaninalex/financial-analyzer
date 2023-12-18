import { Component } from "@angular/core";
import { Store } from "@ngrx/store";
import { IReportState } from "src/app/dashboard/store/report/reducer";
import { selectPriceChartData } from "src/app/dashboard/store/report/selectors";

@Component({
    selector: "price-chart",
    host: { "class": "col-md-6" },
    template: `@if(chartdata.values.length) {
        <div class="card mb-4 overflow-y-scroll">
            <div class="card-body">
                <app-line-chart [data]="chartdata" />
            </div>
        </div>
    }`
})
export class PriceChartComponent {
    chartdata: {label: string, values: Array<[string, number]>} = {label: "Price", values: []};

    constructor(private store: Store<IReportState>) {
        this.store.select(selectPriceChartData).subscribe({
            next: data => {
                if (data.length) this.chartdata = {...this.chartdata, values: data};
            }
        });
    }
}