import { Component } from '@angular/core';
import { Observable } from 'rxjs';
import { AuthService } from '../../services/auth.service';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-verification',
  templateUrl: './verification.component.html',
  styleUrls: ['./verification.component.css']
})
export class VerificationComponent {
    form$: Observable<any>;
    flow: string;

    constructor(private auth: AuthService, private route: ActivatedRoute) {
    }
    
    ngOnInit(): void {
        this.route.queryParams.subscribe(data => {
            this.flow = data["flow"];
            this.form$ = this.auth.formGetVerification(this.flow);
        });
    }
}