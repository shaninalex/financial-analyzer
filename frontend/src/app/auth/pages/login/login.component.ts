import { Component, OnInit } from '@angular/core';
import { AuthService } from '../../services/auth.service';
import { Observable } from 'rxjs';
import { LoginForm } from '../../typedefs/auth';

@Component({
    selector: 'app-login',
    templateUrl: './login.component.html',
})
export class LoginComponent implements OnInit {
    form$: Observable<LoginForm>;

    constructor(private auth: AuthService) { }

    ngOnInit(): void {
        this.form$ = this.auth.formGetLogin();
    }
}
