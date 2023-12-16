import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ProfileComponent } from './profile.component';
import { RouterModule, Routes } from '@angular/router';
import { ReactiveFormsModule } from '@angular/forms';
import { TextFormComponent } from './components/text-form/text-form.component';


const routes: Routes = [
    { path: "", component: ProfileComponent }
]

@NgModule({
    declarations: [
        ProfileComponent,
        TextFormComponent
    ],
    imports: [
        CommonModule,
        RouterModule.forChild(routes),
        ReactiveFormsModule
    ]
})
export class ProfileModule { }
