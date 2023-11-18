import { createAction } from "@ngrx/store";

export const toggleMainSidebar = createAction("[UI] Toggle main sidebar");
export const toggleNotificationsSidebar = createAction("[UI] Toggle notifications sidebar");
export const toggleDarkTheme = createAction("[UI] Toggle dark theme");
