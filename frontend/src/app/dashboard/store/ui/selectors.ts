import { createSelector } from '@ngrx/store';
import { AppState } from '..';
import { UIState } from './reducer';


export const selectFeature = (state: AppState) => state.ui;

export const selectMainSidebarVisible = createSelector(
    selectFeature,
    (state: UIState) => state.mainSidebar
);

export const selectNotificationSidebarVisible = createSelector(
    selectFeature,
    (state: UIState) => state.notificationSidebar
);

export const selectDarkTheme = createSelector(
    selectFeature,
    (state: UIState) => state.dark_theme
);