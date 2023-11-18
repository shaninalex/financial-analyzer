import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './pages/login/login.component';
import { RegisterComponent } from './pages/register/register.component';
import { AuthComponent } from './auth.component';
import { VerificationComponent } from './pages/verification/verification.component';

const routes: Routes = [
    {
        path: "", component: AuthComponent, children: [
            { path: "", component: LoginComponent },
            { path: "register", component: RegisterComponent },
            { path: "verification", component: VerificationComponent }
        ]
    }
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule]
})
export class AuthRoutingModule { }
