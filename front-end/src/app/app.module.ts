import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { HttpClientModule } from '@angular/common/http';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

import { JwtModule } from '@auth0/angular-jwt';

import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';

//JWT token behavior—local storage (decision_making_map is the key name, the actual string of gibberish will be the value paired to it)
export const LOCALSTORAGE_TOKEN_KEY = 'decision_making_map';
//TokenGetter for key
export function tokenGetter() {
  return localStorage.getItem(LOCALSTORAGE_TOKEN_KEY);
}


@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    FormsModule,
    ReactiveFormsModule,
    BrowserModule,
    BrowserAnimationsModule,
	  HttpClientModule,
    //Import module routes
    AppRoutingModule,
    //Material for popup notifs
    MatSnackBarModule,
    //JWT import
    JwtModule.forRoot({
      config: {
        tokenGetter: tokenGetter,
        allowedDomains: ['localhost:4200', 'localhost:9000']
      }
    })
    
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }