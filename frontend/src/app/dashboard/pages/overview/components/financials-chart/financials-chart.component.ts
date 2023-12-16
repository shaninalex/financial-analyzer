import { Component } from "@angular/core";
import { Store } from "@ngrx/store";
import { Chart } from "angular-highcharts";
import { IReportState } from "src/app/dashboard/store/report/reducer";
import { selectFinancialsChartData } from "src/app/dashboard/store/report/selectors";

@Component({
    selector: "financials-chart",
    host: { "class": "col-md-6" },
    template: `@if (financialsChart) {
        <div class="card mb-4 overflow-y-scroll">
            <div class="card-body">
                <div [chart]="financialsChart"></div>
            </div>
        </div>
    }`
})
export class FinancialsChartComponent {
    financialsChart: Chart;

    constructor(private store: Store<IReportState>) {
        this.store.select(selectFinancialsChartData).subscribe({
            next: data => {
                if (data.length) {
                    console.log(data);
                    this.financialsChart = new Chart({
                        chart: {
                            type: 'line'
                        },
                        title: {
                            text: 'Financials'
                        },
                        xAxis: {
                            categories: data.map(d => d.pay_date)
                        },
                        yAxis: {
                            title: {
                                text: data[0].currency
                            }
                        },
                        credits: {
                            enabled: false
                        },
                        series: [{
                            type: "line",
                            name: data[0].type,
                            data: data.map(d => parseFloat(d.amount))
                        }]
                    });
                }
            }
        });
    }
}