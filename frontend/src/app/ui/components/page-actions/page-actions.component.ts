import { Component, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { Observable } from 'rxjs';
import { AppState } from 'src/app/store';
import { toggleMainSidebar } from 'src/app/store/ui/actions';
import { UIState } from 'src/app/store/ui/reducer';
import { selectMainSidebarVisible } from 'src/app/store/ui/selectors';

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
