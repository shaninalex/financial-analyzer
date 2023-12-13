import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { AuthService } from 'src/app/services/auth.service';
import { ActivatedRoute } from '@angular/router';


@Component({
    selector: 'app-login',
    template: `
        <app-generated-form [form$]="form$"></app-generated-form>
        <hr class="hr">
        <div class="d-flex justify-content-between">
            <a routerLink="/auth/registration">Register</a>
            <a class="ui floated right" routerLink="/auth/recovery">Recovery</a>
        </div>`,
})
export class LoginComponent {
    form$: Observable<any>;

    constructor(private auth: AuthService, private route: ActivatedRoute) {
        this.route.queryParams.subscribe(data => this.form$ = this.auth.getLoginFlow(data["flow"]))
    }
}
