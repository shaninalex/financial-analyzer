import { combineReducers } from "@ngrx/store";
import { uiReducer } from "./ui/reducer";
import { reportReducer } from "./report/reducer";


export const dashboardReducer = combineReducers({
    ui: uiReducer,
    report: reportReducer,
})
   