import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SidebarComponent } from './sidebar/sidebar.component';
import { HeaderComponent } from './header/header.component';
import { RouterModule } from '@angular/router';
import { CardComponent } from './components/card/card.component';


@NgModule({
  declarations: [
    SidebarComponent,
    HeaderComponent,
    CardComponent,
  ],
  exports: [
    SidebarComponent,
    HeaderComponent,
    CardComponent,
  ],
  imports: [
    CommonModule,
    RouterModule,
  ]
})
export class UiModule { }
