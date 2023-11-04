import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SidebarComponent } from './sidebar/sidebar.component';
import { HeaderComponent } from './header/header.component';
import { NgIconsModule } from '@ng-icons/core';
import {
    iconoirHome,
    iconoirStar,
    iconoirPageSearch,
    iconoirWallet,
    iconoirUser,
    iconoirSidebarCollapse,
    iconoirSidebarExpand,
    iconoirSearch,
    iconoirSunLight,
    iconoirHalfMoon,
    iconoirBell,
    iconoirGraphUp,
} from '@ng-icons/iconoir';
import { BreadcrumbsComponent } from './components/breadcrumbs/breadcrumbs.component';
import { PageActionsComponent } from './components/page-actions/page-actions.component';
import { GlobalSearchComponent } from './components/global-search/global-search.component';
import { HeaderActionsComponent } from './components/header-actions/header-actions.component';
import { NotificationsSidebarComponent } from './notifications-sidebar/notifications-sidebar.component';
import { NotificationItemComponent } from './notifications-sidebar/notification-item/notification-item.component';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule } from '@angular/router';


@NgModule({
    declarations: [
        SidebarComponent,
        HeaderComponent,
        BreadcrumbsComponent,
        PageActionsComponent,
        GlobalSearchComponent,
        HeaderActionsComponent,
        NotificationsSidebarComponent,
        NotificationItemComponent,
    ],
    exports: [
        SidebarComponent,
        HeaderComponent,
        NotificationsSidebarComponent,
        PageActionsComponent,
        HeaderActionsComponent
    ],
    imports: [
        CommonModule,
        RouterModule,
        NgIconsModule.withIcons({
            iconoirHome,
            iconoirStar,
            iconoirPageSearch,
            iconoirWallet,
            iconoirUser,
            iconoirSidebarCollapse,
            iconoirSidebarExpand,
            iconoirSearch,
            iconoirSunLight,
            iconoirHalfMoon,
            iconoirBell,
            iconoirGraphUp,
        }),
    ]
})
export class UiModule { }
