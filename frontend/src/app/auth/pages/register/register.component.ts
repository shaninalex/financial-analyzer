import { Component } from '@angular/core';
import { AuthService } from '../../services/auth.service';

@Component({
    selector: 'app-register',
    templateUrl: './register.component.html',
    styleUrls: ['./register.component.css']
})
export class RegisterComponent {
    constructor(private auth: AuthService) { }

    ngOnInit(): void {
        this.auth.formGetRegistration().subscribe({
            next: data => {
                console.log(data);
            }
        })
    }
}
