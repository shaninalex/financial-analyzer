import { Component } from '@angular/core';
import { AuthService } from '../../services/auth.service';
import { RegistrationForm } from '../../typedefs/auth';
import { Observable } from 'rxjs';

@Component({
    selector: 'app-register',
    templateUrl: './register.component.html',
    styleUrls: ['./register.component.css']
})
export class RegisterComponent {
    form$: Observable<RegistrationForm>;

    constructor(private auth: AuthService) { }

    ngOnInit(): void {
        this.form$ = this.auth.formGetRegistration();
    }
}
