import { Component, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { UIState } from '../../store/ui/reducer';
import { selectDarkTheme } from '../../store/ui/selectors';
import { Observable } from 'rxjs';
import { toggleDarkTheme } from '../../store/ui/actions';
import { Traits } from 'src/app/store/typedefs';
import { ProfileService } from '../../services/profile.service';
import { selectTraits } from 'src/app/store/selectors';
import { IAppState } from 'src/app/store/store';


@Component({
    selector: 'app-header',
    templateUrl: './header.component.html',
})
export class HeaderComponent {
    theme$: Observable<boolean>;
    identity$: Observable<Traits | undefined>;

    constructor(
        private store: Store<UIState>,
        private storeIdentity: Store<IAppState>,
        private profile: ProfileService,
    ) {
        this.identity$ = this.storeIdentity.select(selectTraits);
        this.theme$ = this.store.select(selectDarkTheme);
    }

    toggleTheme(): void {
        this.store.dispatch(toggleDarkTheme());
    }

    logout(): void {
        this.profile.getLogoutLink().subscribe({
            next: data => {
                window.location.href = data.logout_url;
            }
        })
    }
}
