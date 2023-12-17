import { Component, OnInit } from '@angular/core';
import { ProfileService } from '../../services/profile.service';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
})
export class ProfileComponent implements OnInit {
  form$: Observable<any>;
  active_tab: string = "profile";
  nodes_profile: Array<any>;
  nodes_password: Array<any>;
  nodes_totp: Array<any>;
  form_settings: any = {};

  constructor(private profile: ProfileService) { }

  ngOnInit(): void {
    this.profile.getSettings().subscribe({
        next: (data: any) => {
            this.form_settings.action = data.ui.action;
            this.form_settings.method = "POST";
            this.form_settings.nodes = data.ui.nodes.filter((node:any) => node.group === "default");
            this.nodes_profile = data.ui.nodes.filter((node:any) => node.group === "profile");
            this.nodes_password = data.ui.nodes.filter((node:any) => node.group === "password");
            this.nodes_totp = data.ui.nodes.filter((node:any) => node.group === "totp" || node.group === "lookup_secret");
        }
    });
}

  changeTab(tab_name: string): void {
    this.active_tab = tab_name
  }
}
