import { createReducer, on } from '@ngrx/store';
import * as UiActions from './actions';

export interface UIState {
    mainSidebar: boolean
    notificationSidebar: boolean
    dark_theme: boolean
}

export const InitialUiState: UIState = {
    mainSidebar: true,
    notificationSidebar: true,
    dark_theme: true,
}

export const uiReducer = createReducer(
    InitialUiState,
    on(UiActions.toggleMainSidebar, state => ({...state, mainSidebar: !state.mainSidebar})),
    on(UiActions.toggleNotificationsSidebar, state => ({...state, notificationSidebar: !state.notificationSidebar})),
    on(UiActions.toggleDarkTheme, state => ({...state, dark_theme: !state.dark_theme})),
)