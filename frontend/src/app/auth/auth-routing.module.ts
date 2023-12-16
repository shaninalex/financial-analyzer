import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './pages/login.component';
import { RegisterComponent } from './pages/register.component';
import { AuthComponent } from './auth.component';
import { VerificationComponent } from './pages/verification.component';
import { RecoveryComponent } from './pages/recovery.component';

const routes: Routes = [{
  path: "",
  component: AuthComponent,
  children: [
    { path: "login", component: LoginComponent },
    { path: "registration", component: RegisterComponent },
    { path: "verification", component: VerificationComponent },
    { path: "recovery", component: RecoveryComponent },
    { path: "", redirectTo: "login", pathMatch: 'full' }
  ]
}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class AuthRoutingModule { }
