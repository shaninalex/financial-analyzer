import { createSelector } from '@ngrx/store';
import { IReportState } from './reducer';


export const selectUI = (state: any) => state.dashboard.report;

export const selectPriceChartData = createSelector(
    selectUI,
    (state: IReportState) => state.priceChartData
);
