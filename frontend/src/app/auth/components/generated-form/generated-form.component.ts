import { Component, Input } from '@angular/core';
import { Observable } from 'rxjs';

@Component({
    selector: 'app-generated-form',
    templateUrl: 'generated-form.component.html'
})
export class GeneratedFormComponent {
    @Input() form$: Observable<any>;
}
