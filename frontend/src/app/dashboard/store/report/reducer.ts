import { createReducer, on } from '@ngrx/store';
import * as ReportActions from './actions';
import { IFinancials } from '../../typedefs/global';

export interface IReportState {
    priceChartData: Array<[string, number]>
    financialsChartData: Array<IFinancials>
}

export const InitialReportState: IReportState = {
    priceChartData: [],
    financialsChartData: []
}

export const reportReducer = createReducer(
    InitialReportState,
    on(ReportActions.setPriceChartData, (state, action) => ({...state, priceChartData: action.data})),
    on(ReportActions.setFinancialsChartData, (state, action) => ({...state, financialsChartData: action.data})),
)