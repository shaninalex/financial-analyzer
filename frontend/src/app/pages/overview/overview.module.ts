import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { OverviewComponent } from './overview.component';
import { RouterModule, Routes } from '@angular/router';
import { UiModule } from 'src/app/ui/ui.module';
import { NgIconsModule } from '@ng-icons/core';
import { 
    iconoirArrowTr,
    iconoirSaveFloppyDisk,
    iconoirPlus,
} from '@ng-icons/iconoir';
import { ReactiveFormsModule } from '@angular/forms';

const routes: Routes = [
    { path: "", component: OverviewComponent }
]

@NgModule({
    declarations: [
        OverviewComponent
    ],
    imports: [
        CommonModule,
        UiModule,
        ReactiveFormsModule,
        NgIconsModule.withIcons({ 
            iconoirArrowTr,
            iconoirSaveFloppyDisk,
            iconoirPlus,
        }),
        RouterModule.forChild(routes)
    ]
})
export class OverviewModule { }

