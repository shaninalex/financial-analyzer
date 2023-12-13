import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { LoginForm } from '../../typedefs/auth';
import { AuthService } from 'src/app/services/auth.service';


@Component({
    selector: 'app-login',
    templateUrl: './login.component.html',
})
export class LoginComponent implements OnInit {
    form$: Observable<LoginForm>;

    constructor(private auth: AuthService) { }

    ngOnInit(): void {
        this.form$ = this.auth.getLoginFlow();
    }
}
