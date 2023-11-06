import { Component } from '@angular/core';
import { WebsocketService } from './services/websocket.service';
import { ITickerAction } from './typedefs';
import { Store } from '@ngrx/store';
import { AppState } from './store';
import { Observable } from 'rxjs';
import { selectMainSidebarVisible, selectNotificationSidebarVisible } from './store/ui/selectors';

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.css']
})
export class AppComponent {
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
