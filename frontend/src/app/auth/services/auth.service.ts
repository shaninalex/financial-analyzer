import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { RegistrationForm } from "../typedefs/auth";
import { Observable, shareReplay } from "rxjs";

@Injectable()
export class AuthService {

    constructor(private http: HttpClient) {}

    formGetRegistration(): Observable<RegistrationForm> {
        return this.http.get<RegistrationForm>("http://127.0.0.1:8081/api/v2/auth/get-registration-form").pipe(
            shareReplay()
        );
    }
}