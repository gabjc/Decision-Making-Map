import { PublicRoutingModule } from './public-routing.module';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { signinComponent } from './signin/signin.component';
import { LoginComponent } from './login/login.component';
import { HomeComponent } from './home/home.component';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';




@NgModule({
  declarations: [
    LoginComponent,
    signinComponent,
    HomeComponent
  ],
  imports: [
    CommonModule,
    PublicRoutingModule,

    ReactiveFormsModule,
    FormsModule,
    
    //Angular Material imports
    MatSlideToggleModule,
    MatInputModule,
    MatFormFieldModule,
    MatButtonModule,
    MatCardModule
  ]
})
export class PublicModule { }
