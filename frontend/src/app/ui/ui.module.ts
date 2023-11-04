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
    iconoirSidebarExpand
} from '@ng-icons/iconoir';


@NgModule({
    declarations: [
        SidebarComponent,
        HeaderComponent
    ],
    exports: [
        SidebarComponent,
        HeaderComponent
    ],
    imports: [
        CommonModule,
        NgIconsModule.withIcons({
            iconoirHome,
            iconoirStar,
            iconoirPageSearch,
            iconoirWallet,
            iconoirUser,
            iconoirSidebarCollapse,
            iconoirSidebarExpand
        }),
    ]
})
export class UiModule { }
