import { Component, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { Observable } from 'rxjs';
import { AppState } from 'src/app/dashboard/store';
import { toggleMainSidebar } from 'src/app/dashboard/store/ui/actions';
import { selectMainSidebarVisible } from 'src/app/dashboard/store/ui/selectors';

@Component({
    selector: 'app-page-actions',
    templateUrl: './page-actions.component.html',
    styleUrls: ['./page-actions.component.css']
})
export class PageActionsComponent implements OnInit {
    collapsed: Observable<boolean>;

    constructor(private store: Store<AppState>) {}

    ngOnInit(): void {
        this.collapsed = this.store.select(selectMainSidebarVisible);
    }

    toggleMainSidebar(): void {
        console.log("toggleMainSidebar");
        this.store.dispatch(toggleMainSidebar());
    }
}
