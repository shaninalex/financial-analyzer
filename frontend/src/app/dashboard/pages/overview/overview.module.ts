import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { OverviewComponent } from './overview.component';
import { RouterModule, Routes } from '@angular/router';
import { UiModule } from '../../ui/ui.module';
import { ChartModule } from 'angular-highcharts';

import { ReactiveFormsModule } from '@angular/forms';
import { PriceChartComponent } from './components/price-chart/price-chart.components';
import { FinancialsChartComponent } from './components/financials-chart/financials-chart.component';

const routes: Routes = [
    { path: "", component: OverviewComponent }
]

@NgModule({
    declarations: [
        OverviewComponent,
        PriceChartComponent,
        FinancialsChartComponent
    ],
    imports: [
        CommonModule,
        UiModule,
        ReactiveFormsModule,
        RouterModule.forChild(routes),
        ChartModule
    ]
})
export class OverviewModule { }

