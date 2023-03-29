import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot, UrlTree } from '@angular/router';
import { JwtHelperService } from '@auth0/angular-jwt';
import { Observable } from 'rxjs';
import { AuthService } from '../public/services/auth.service';

@Injectable({
  providedIn: 'root'
})

export class AuthGuard{

  constructor(
    private authService : AuthService,
    private router: Router,
    private jwtService: JwtHelperService
  ) {}
}
