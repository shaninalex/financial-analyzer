import { createReducer, on } from '@ngrx/store';
import * as ReportActions from './actions';

export interface IReportState {
    priceChartData: Array<[string, number]>
}

export const InitialReportState: IReportState = {
    priceChartData: [],
}

export const reportReducer = createReducer(
    InitialReportState,
    on(ReportActions.setPriceChartData, (state, action) => ({...state, priceChartData: action.data})),
)