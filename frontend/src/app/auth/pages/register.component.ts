import { Component } from '@angular/core';
import { AuthService } from '../services/auth.service';
import { Observable } from 'rxjs';

@Component({
    selector: 'app-register',
    template: `
        <app-generated-form [form$]="form$"></app-generated-form>
        <hr class="hr">
        <div class="d-flex justify-content-between">
            <a routerLink="/auth/login">Login</a>
            <a class="ui floated right" routerLink="/auth/recovery">Recovery</a>
        </div>`,
})
export class RegisterComponent {
    form$: Observable<any>;

    constructor(private auth: AuthService) { }

    ngOnInit(): void {
        this.form$ = this.auth.getRegistrationFlow();
    }
}
