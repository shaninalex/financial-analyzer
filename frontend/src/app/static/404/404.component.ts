import { Component, ViewEncapsulation } from '@angular/core';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-404',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './404.component.html',
  encapsulation: ViewEncapsulation.None
})
export class NotFoundComponent {

}
