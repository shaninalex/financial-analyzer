import { Component } from "@angular/core";

@Component({
    selector: "card",
    template: `
    <div class="card mb-4">
      <div class="card-body">
        <ng-content></ng-content>
      </div>
    </div>
    `
})
export class CardComponent {}