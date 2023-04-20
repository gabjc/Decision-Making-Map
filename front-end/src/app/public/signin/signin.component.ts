import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, NgControl, Validators, FormBuilder } from '@angular/forms';
import { map, tap } from 'rxjs/operators';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/public/services/auth.service';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-signin',
  templateUrl: './signin.component.html',
  styleUrls: ['./signin.component.css']
})
export class signinComponent implements OnInit{

  registerForm!: FormGroup;

  constructor(
    private router: Router,
    private http: HttpClient,
    private authService: AuthService,
  ) { }

  ngOnInit() : void {
    this.registerForm = new FormGroup ({
      name: new FormControl(null, [
        Validators.required
      ]),
      email: new FormControl(null, [
        Validators.required,
        Validators.email
      ]),
      password: new FormControl(null, [
        Validators.required
      ])
    });
  }

  register(): void {
    if (this.registerForm.valid) {
      const name = this.registerForm.get('name')?.value;
      const email = this.registerForm.get('email')?.value;
      const password = this.registerForm.get('password')?.value;
      console.log(name, email, password);
    this.authService.register(name, email, password).subscribe(
      () => {
        this.router.navigate(['../login'])
      }
    );
    }
  }
}
