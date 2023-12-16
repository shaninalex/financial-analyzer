import { createAction, props } from "@ngrx/store";

export const setPriceChartData = createAction(
    "[report] Set Price Chart Data",
    props<{ data: Array<[string, number]> }>()
);
