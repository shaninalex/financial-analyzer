import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RegisterComponent } from './pages/register/register.component';
import { LoginComponent } from './pages/login/login.component';
import { AuthComponent } from './auth.component';
import { AuthRoutingModule } from './auth-routing.module';
import { AuthService } from './services/auth.service';
import { HttpClientModule } from '@angular/common/http';


@NgModule({
    declarations: [
        RegisterComponent,
        LoginComponent,
        AuthComponent,
    ],
    imports: [
        CommonModule,
        AuthRoutingModule,
        HttpClientModule
    ],
    providers: [
        AuthService
    ]
})
export class AuthModule { }
