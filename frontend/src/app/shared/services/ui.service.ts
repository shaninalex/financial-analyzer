import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable({
    providedIn: 'root'
})
export class UiService {
    public darkTheme: boolean = true; // true = "dark", false = "light"
    public loading: BehaviorSubject<boolean> = new BehaviorSubject<boolean>(false);
    
    constructor() {}

    switchTheme(): void {
        this.darkTheme = !this.darkTheme;
        if (this.darkTheme) {
            document.documentElement.setAttribute("data-bs-theme", "dark");
        } else {
            document.documentElement.setAttribute("data-bs-theme", "light");
        }
    }
}