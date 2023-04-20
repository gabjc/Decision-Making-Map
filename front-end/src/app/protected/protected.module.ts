import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ProtectedRoutingModule } from './protected-routing.module';


import { DashboardComponent } from './dashboard/dashboard.component';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatButtonModule } from '@angular/material/button';



@NgModule({
  declarations: [
    DashboardComponent
  ],
  imports: [
    CommonModule,
    ProtectedRoutingModule,
    FormsModule,
    ReactiveFormsModule,
    MatButtonModule,
    MatFormFieldModule
  ]
})
export class ProtectedModule { }
