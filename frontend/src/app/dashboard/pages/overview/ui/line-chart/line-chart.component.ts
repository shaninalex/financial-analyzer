import { AfterViewInit, Component, ElementRef, Input, ViewChild } from '@angular/core';
import Chart from 'chart.js/auto';


@Component({
  selector: 'app-line-chart',
  template: `<canvas #chart class="chart-item"></canvas>`,
  styleUrls: ['line-chart.component.css']
})
export class LineChartComponent implements AfterViewInit {
  @ViewChild('chart') private chartRef: ElementRef;
  @Input('data') data: {label: string, values: Array<[string, number]>};

  public chart: any;

  ngAfterViewInit() {
    this.chart = new Chart(this.chartRef.nativeElement, {
      type: "line",
      data: {
        labels: this.data.values.map(d => d[0]),
        datasets: [
          {
            label: this.data.label,
            data: this.data.values.map(d => d[1]),
            pointRadius: 0,
            showLine: true,
            borderColor: "blue"
          }
        ]
      },
      options: {
        aspectRatio: 2.5,
        interaction: {
          intersect: false
        },
      }
    });
  }
}
