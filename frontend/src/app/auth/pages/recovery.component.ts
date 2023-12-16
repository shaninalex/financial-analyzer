import { Component } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Observable } from 'rxjs';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-recovery',
  template: `<app-generated-form [form$]="form$"></app-generated-form>`,
})
export class RecoveryComponent {
  form$: Observable<any>;

  constructor(private auth: AuthService, private route: ActivatedRoute) {
      this.route.queryParams.subscribe(data => this.form$ = this.auth.getRecoveryFlow(data["flow"]));
  }
}
