import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { map, tap } from 'rxjs/operators';
import { Router } from '@angular/router';
import { AuthService } from 'src/app/public/services/auth.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  loginForm!: FormGroup;

  constructor(
    private authService: AuthService,
    private router: Router
  ) { }

  ngOnInit() : void {
    this.loginForm = new FormGroup({
      email: new FormControl(null, [
        Validators.required,
        Validators.email
      ]),
      password: new FormControl(null, [
        Validators.required
      ])
    });
  }
  
  login() {
    if (this.loginForm.valid) {
      const email = this.loginForm.get('email')?.value;
      const password = this.loginForm.get('password')?.value;
      console.log(email, password);
      this.authService.login(email, password).subscribe(
        () => {
          //protected dashboard
          this.router.navigate(['src/app/protected/dashboard'])
        }
      );
    }
  }
}
