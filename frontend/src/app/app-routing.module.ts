import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

const routes: Routes = [
    {
        path: "",
        loadChildren: () => import('./pages/overview/overview.module').then(m => m.OverviewModule)
    },
    {
        path: "reports",
        loadChildren: () => import('./pages/reports/reports.module').then(m => m.ReportsModule)
    },
    {
        path: "subscription",
        loadChildren: () => import('./pages/subscription/subscription.module').then(m => m.SubscriptionModule)
    },
    {
        path: "profile",
        loadChildren: () => import('./pages/profile/profile.module').then(m => m.ProfileModule)
    }
];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})
export class AppRoutingModule { }
