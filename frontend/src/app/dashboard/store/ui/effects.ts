import { Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { tap } from 'rxjs/operators';
import { toggleDarkTheme } from "./actions";


@Injectable()
export class UIEffects {
    switchTheme$ = createEffect(() => this.actions$.pipe(
        ofType(toggleDarkTheme),
        tap(() => {
            const theme = document.documentElement.getAttribute("data-theme");
            if (!theme || theme !== "dark") {
                document.documentElement.setAttribute("data-theme", "dark");
            } else {
                document.documentElement.removeAttribute("data-theme");
            }
        })
    ), {dispatch: false});

    constructor(
        private actions$: Actions,
    ) { }
}