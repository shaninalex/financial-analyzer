import { Component, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { UIState } from '../../store/ui/reducer';
import { selectDarkTheme } from '../../store/ui/selectors';
import { Observable } from 'rxjs';
import { toggleDarkTheme } from '../../store/ui/actions';


@Component({
    selector: 'app-header',
    templateUrl: './header.component.html',
})
export class HeaderComponent {
    theme$: Observable<boolean>;

    constructor(private store: Store<UIState>) {
        this.theme$ = this.store.select(selectDarkTheme);
    }

    toggleTheme(): void {
        this.store.dispatch(toggleDarkTheme());
    }
}
