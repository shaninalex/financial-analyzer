import { createAction, props } from "@ngrx/store";
import { IFinancials } from "../../typedefs/global";

export const setPriceChartData = createAction(
    "[report] Set Price Chart Data",
    props<{ data: Array<[string, number]> }>()
);

export const setFinancialsChartData = createAction(
    "[report] Set Financials Chart Data",
    props<{ data: Array<IFinancials> }>()
);

