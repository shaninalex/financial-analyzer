import { NgModule, isDevMode } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DashboardComponent } from './dashboard.component';
import { DashboardRoutingModule } from './dashboard-routing.module';
import { HttpClientModule } from '@angular/common/http';
import { UiModule } from './ui/ui.module';
import { StoreModule } from '@ngrx/store';
import { EffectsModule } from '@ngrx/effects';
import { UIEffects } from './store/ui/effects';
import { StoreDevtoolsModule } from '@ngrx/store-devtools';
import { dashboardReducer } from './store';
import { ProfileService } from './services/profile.service';
import { WebsocketService } from './services/websocket.service';


@NgModule({
    declarations: [
        DashboardComponent
    ],
    imports: [
        CommonModule,
        DashboardRoutingModule,
        HttpClientModule,
        UiModule,
        StoreModule.forFeature("dashboard", dashboardReducer),
        EffectsModule.forFeature([UIEffects]),
        StoreDevtoolsModule.instrument({ maxAge: 25, logOnly: !isDevMode() , connectInZone: true})
    ],
    providers: [
        ProfileService,
        WebsocketService
    ]
})
export class DashboardModule { }
