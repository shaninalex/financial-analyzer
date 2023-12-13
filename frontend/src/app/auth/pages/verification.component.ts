import { Component } from '@angular/core';
import { Observable } from 'rxjs';
import { AuthService } from 'src/app/services/auth.service';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-verification',
  template: `<app-generated-form [form$]="form$"></app-generated-form>`
})
export class VerificationComponent {
    form$: Observable<any>;

    constructor(
        private auth: AuthService,
        private route: ActivatedRoute
    ) {
        this.route.queryParams.subscribe(data => this.form$ = this.auth.getVerification(data["flow"]));
    }
}
