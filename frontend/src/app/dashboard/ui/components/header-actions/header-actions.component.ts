import { Component } from '@angular/core';
import { Store } from '@ngrx/store';
import { Observable } from 'rxjs';
import { AppState } from 'src/app/dashboard/store';
import { toggleDarkTheme, toggleNotificationsSidebar } from 'src/app/dashboard/store/ui/actions';
import { selectDarkTheme, selectNotificationSidebarVisible } from 'src/app/dashboard/store/ui/selectors';

@Component({
  selector: 'app-header-actions',
  templateUrl: './header-actions.component.html',
  styleUrls: ['./header-actions.component.css']
})
export class HeaderActionsComponent {
    notifySidebar: Observable<boolean>;
    darkTheme: Observable<boolean>;

    constructor(private store: Store<AppState>) {
        this.notifySidebar = this.store.select(selectNotificationSidebarVisible);
        this.darkTheme = this.store.select(selectDarkTheme);
    }

    notificationsSidebar(): void {
        this.store.dispatch(toggleNotificationsSidebar());
    }

    toggleTheme(): void {
        this.store.dispatch(toggleDarkTheme());
    }
}
