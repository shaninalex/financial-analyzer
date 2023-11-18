import { Component } from '@angular/core';
import { Observable } from 'rxjs';
import { WebsocketService } from './services/websocket.service';
import { Store } from '@ngrx/store';
import { AppState } from './store';
import { selectMainSidebarVisible, selectNotificationSidebarVisible } from './store/ui/selectors';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent {
    mainSidebarCollapsed: Observable<boolean>;
    notificationSidebarVisible: Observable<boolean>;
    
    constructor(
        private socket: WebsocketService,
        private store: Store<AppState>
    ) {}

    ngOnInit(): void {
        this.mainSidebarCollapsed = this.store.select(selectMainSidebarVisible);
        this.notificationSidebarVisible = this.store.select(selectNotificationSidebarVisible);
    }
}
