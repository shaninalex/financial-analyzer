import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Observable } from "rxjs";


@Injectable()
export class ProfileService {
    constructor(
        private http: HttpClient
    ) { }

    getLogoutLink(): Observable<any> {
        return this.http.get<any>("/api/v2/auth/logout", { withCredentials: true })
    }
}
