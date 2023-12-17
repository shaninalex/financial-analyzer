import { Component, ViewEncapsulation } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClient, HttpClientModule, HttpParams } from '@angular/common/http';
import { ActivatedRoute, RouterModule } from '@angular/router';
import { Observable } from 'rxjs';

@Component({
    selector: 'app-error',
    standalone: true,
    imports: [
        CommonModule,
        HttpClientModule,
        RouterModule
    ],
    templateUrl: './error.component.html',
    encapsulation: ViewEncapsulation.None
})
export class ErrorComponent {
    error$: Observable<any>;

    constructor(
        private http: HttpClient,
        private route: ActivatedRoute
    ) {
        this.route.queryParams.subscribe({
            next: data => {
                let params = new HttpParams();
                params = params.append("id", data["id"]);
                this.error$ = this.http.get<any>(`/api/v2/auth/error`, { params: params, withCredentials: true });
            }
        })
    }
}
