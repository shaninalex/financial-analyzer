import { createSelector } from '@ngrx/store';
import { UIState } from './reducer';


export const selectUI = (state: any) => state.dashboard.ui;

export const selectMainSidebarVisible = createSelector(
    selectUI,
    (state: UIState) => state.mainSidebar
);

export const selectNotificationSidebarVisible = createSelector(
    selectUI,
    (state: UIState) => state.notificationSidebar
);

export const selectDarkTheme = createSelector(
    selectUI,
    (state: UIState) => state.dark_theme
);