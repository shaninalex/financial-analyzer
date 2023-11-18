import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { DashboardComponent } from './dashboard.component';

const routes: Routes = [
    {
        path: "", component: DashboardComponent, children: [
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
        ]
    }
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule]
})
export class DashboardRoutingModule { }
