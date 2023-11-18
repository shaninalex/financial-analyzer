import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DashboardComponent } from './dashboard.component';
import { DashboardRoutingModule } from './dashboard-routing.module';
import { HttpClientModule } from '@angular/common/http';
import { UiModule } from './ui/ui.module';
import { StoreModule } from '@ngrx/store';
import { EffectsModule } from '@ngrx/effects';
import { uiReducer } from './store/ui/reducer';
import { UIEffects } from './store/ui/effects';


@NgModule({
    declarations: [
        DashboardComponent
    ],
    imports: [
        CommonModule,
        DashboardRoutingModule,
        HttpClientModule,
        UiModule,
        StoreModule.forFeature("dashboard", {
            ui: uiReducer
        }, {}),
        EffectsModule.forFeature([UIEffects]),
    ]
})
export class DashboardModule { }
