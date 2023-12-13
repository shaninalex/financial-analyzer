import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RegisterComponent } from './pages/register/register.component';
import { LoginComponent } from './pages/login/login.component';
import { AuthComponent } from './auth.component';
import { AuthRoutingModule } from './auth-routing.module';
import { HttpClientModule } from '@angular/common/http';
import { VerificationComponent } from './pages/verification/verification.component';
import { RecoveryComponent } from './pages/recovery/recovery.component';


@NgModule({
    declarations: [
        RegisterComponent,
        LoginComponent,
        AuthComponent,
        VerificationComponent,
        RecoveryComponent,
    ],
    imports: [
        CommonModule,
        AuthRoutingModule,
        HttpClientModule
    ],
})
export class AuthModule { }
