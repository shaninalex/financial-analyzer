import { Action, createReducer, on } from '@ngrx/store';
import * as UiActions from './actions';

export interface UIState {
    mainSidebar: boolean
    notificationSidebar: boolean
}

export const InitialUiState: UIState = {
    mainSidebar: true,
    notificationSidebar: true,
}

export const uiReducer = createReducer(
    InitialUiState,
    on(UiActions.toggleMainSidebar, state => ({...state, mainSidebar: !state.mainSidebar})),
    on(UiActions.toggleNotificationsSidebar, state => ({...state, notificationSidebar: !state.notificationSidebar})),
)