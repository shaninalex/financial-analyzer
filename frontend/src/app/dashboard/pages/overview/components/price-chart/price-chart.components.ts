import { Component } from "@angular/core";
import { Store } from "@ngrx/store";
import { Chart } from "angular-highcharts";
import { IReportState } from "src/app/dashboard/store/report/reducer";
import { selectPriceChartData } from "src/app/dashboard/store/report/selectors";

@Component({
    selector: "price-chart",
    host: { "class": "col-md-6" },
    template: `@if(priceChart) {
        <div class="card mb-4 overflow-y-scroll">
            <div class="card-body">
                <div [chart]="priceChart"></div>
            </div>
        </div>
    }`
})
export class PriceChartComponent {
    priceChart: Chart;

    constructor(private store: Store<IReportState>) {
        this.store.select(selectPriceChartData).subscribe({
            next: data => {
                if (data.length) {
                    this.priceChart = new Chart({
                        chart: {
                            type: 'line'
                        },
                        title: {
                            text: 'Price'
                        },
                        xAxis: {
                            categories: data.map(d => d[0])
                        },
                        yAxis: {
                            title: {
                                text: "$"
                            }
                        },
                        credits: {
                            enabled: false
                        },
                        series: [{
                            type: "line",
                            name: 'Line 1',
                            data: data.map(i => i[1])
                        }]
                    });
                }
            }
        });
    }
}