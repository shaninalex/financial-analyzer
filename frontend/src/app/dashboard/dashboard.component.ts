import { Component } from '@angular/core';
import { Observable } from 'rxjs';
import { WebsocketService } from './services/websocket.service';
import { Store } from '@ngrx/store';
import { selectMainSidebarVisible, selectNotificationSidebarVisible } from './store/ui/selectors';
import { UIState } from './store/ui/reducer';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['dashboard.component.css'],
})
export class DashboardComponent {
    mainSidebarCollapsed: Observable<boolean>;
    notificationSidebarVisible: Observable<boolean>;

    constructor(
        private socket: WebsocketService,
        private store: Store<UIState>
    ) {}

    ngOnInit(): void {
        this.mainSidebarCollapsed = this.store.select(selectMainSidebarVisible);
        this.notificationSidebarVisible = this.store.select(selectNotificationSidebarVisible);
    }
}
