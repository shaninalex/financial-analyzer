import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CanActiveteAccountPage } from './shared/guards/auth.guard';

const routes: Routes = [
    {
        path: "auth", loadChildren: () => import('./auth/auth.module').then(m => m.AuthModule)
    },
    {
        path: "app", 
        canLoad: [CanActiveteAccountPage],
        loadChildren: () => import('./dashboard/dashboard.module').then(m => m.DashboardModule)
    },
    {
        path: "404",
        loadComponent: () => import("./static/404/404.component").then(c => c.NotFoundComponent)
    },
    {
        path: "error",
        loadComponent: () => import("./static/error/error.component").then(c => c.ErrorComponent)
    },
    {
        path: "", redirectTo: "app", pathMatch: 'full'
    },
    {
        path: "**", redirectTo: "404"
    }
];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})
export class AppRoutingModule { }
