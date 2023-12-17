import { createSelector } from '@ngrx/store';
import { IReportState } from './reducer';


const report = (state: any) => state.dashboard.report;

export const selectPriceChartData = createSelector(
    report,
    (state: IReportState) => state.priceChartData
);

export const selectFinancialsChartData = createSelector(
    report,
    (state: IReportState) => state.financialsChartData
);
