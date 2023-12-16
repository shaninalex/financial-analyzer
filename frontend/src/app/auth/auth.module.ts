import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RegisterComponent } from './pages/register.component';
import { LoginComponent } from './pages/login.component';
import { AuthComponent } from './auth.component';
import { AuthRoutingModule } from './auth-routing.module';
import { HttpClientModule } from '@angular/common/http';
import { VerificationComponent } from './pages/verification.component';
import { RecoveryComponent } from './pages/recovery.component';
import { GeneratedFormComponent } from './components/generated-form/generated-form.component';
import { ReactiveFormsModule } from '@angular/forms';
import { AuthService } from './services/auth.service';


@NgModule({
    declarations: [
        RegisterComponent,
        LoginComponent,
        AuthComponent,
        VerificationComponent,
        RecoveryComponent,
        GeneratedFormComponent,
    ],
    imports: [
        CommonModule,
        AuthRoutingModule,
        HttpClientModule,
        ReactiveFormsModule,
    ],
    providers: [
        AuthService
    ]
})
export class AuthModule { }
