import { HttpClient, HttpParams } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, catchError, finalize, shareReplay, throwError } from 'rxjs';
import { MessagesService } from '../../shared/services/messages.service';
import { UiService } from '../../shared/services/ui.service';


@Injectable()
export class AuthService {

    private request_params = {
        withCredentials: true
    }

    constructor(
        private http: HttpClient,
        private messages: MessagesService,
        private ui: UiService
    ) {}

    private handleRequest<T>(observable: Observable<T>): Observable<T> {
        this.ui.loading.next(true);
        return observable.pipe(
            finalize(() => this.ui.loading.next(false)),
            shareReplay(),
            catchError(err => {
                this.messages.message.next(err);
                return throwError(() => err);
            })
        )
    }

    getError(error_id: string | null = null): Observable<any> {
        let params = new HttpParams();
        if (error_id) params = params.append("id", error_id);
        return this.handleRequest(this.http.get<any>(`/api/v2/auth/error`, { params: params, ...this.request_params }));
    }

    getLoginFlow(flow: string | null = null): Observable<any> {
        let params = new HttpParams();
        if (flow) params = params.append("id", flow);
        return this.handleRequest(this.http.get<any>(`/api/v2/auth/login`, { params: params, ...this.request_params }))
    }

    getRegistrationFlow(flow: string | null = null): Observable<any> {
        let params = new HttpParams();
        if (flow) params = params.append("id", flow);
        return this.handleRequest(this.http.get<any>(`/api/v2/auth/registration`, { params: params, ...this.request_params }));
    }

    getVerification(flow: string): Observable<any> {
        let params = new HttpParams().append("flow", flow);
        return this.handleRequest(this.http.get<any>(`/api/v2/auth/verification`, { params: params, ...this.request_params }))
    }

    getRecoveryFlow(flow: string | null = null): Observable<any> {
        let params = new HttpParams();
        if (flow) params = params.append("id", flow);
        return this.handleRequest(this.http.get<any>(`/api/v2/auth/recovery`, { params: params, withCredentials: true }))
    }
}
