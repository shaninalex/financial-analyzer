import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SubscriptionComponent } from './subscription.component';
import { RouterModule, Routes } from '@angular/router';

const routes: Routes = [
    { path: "", component: SubscriptionComponent }
]

@NgModule({
    declarations: [
        SubscriptionComponent
    ],
    imports: [
        CommonModule,
        RouterModule.forChild(routes)
    ]
})
export class SubscriptionModule { }
