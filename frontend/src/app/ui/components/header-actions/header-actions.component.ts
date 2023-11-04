import { Component } from '@angular/core';
import { Store } from '@ngrx/store';
import { Observable } from 'rxjs';
import { AppState } from 'src/app/store';
import { toggleNotificationsSidebar } from 'src/app/store/ui/actions';

@Component({
  selector: 'app-header-actions',
  templateUrl: './header-actions.component.html',
  styleUrls: ['./header-actions.component.css']
})
export class HeaderActionsComponent {
    notifySidebar: Observable<boolean>;

    constructor(private store: Store<AppState>) {}

    notificationsSidebar(): void {
        this.store.dispatch(toggleNotificationsSidebar());
    }
}
