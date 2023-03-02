import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { HttpClientModule } from '@angular/common/http'; 

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { MatInputModule } from '@angular/material/input';
import {  MatFormFieldModule } from '@angular/material/form-field';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';

import { Routes, RouterModule } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { signinComponent } from './signin/signin.component';
import { HomeComponent } from './home/home.component';

const routes: Routes = [
  { path: '', component: HomeComponent },
  {path: 'login', component: LoginComponent},
  {path: 'signin', component: signinComponent}
];

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    signinComponent,
    HomeComponent

  ],
  imports: [
    RouterModule.forRoot(routes),
    MatSlideToggleModule,
    MatInputModule,
    MatFormFieldModule,
    MatButtonModule,
    MatCardModule,
    BrowserModule,
    AppRoutingModule,
	  HttpClientModule,
    BrowserAnimationsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }