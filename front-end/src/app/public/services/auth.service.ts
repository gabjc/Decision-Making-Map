import { Injectable, EventEmitter } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { map, Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private apiUrl = 'http://localhost:9000';
  authenticated = false;
  authenticated$: EventEmitter<boolean> = new EventEmitter<boolean>();

  constructor(private http: HttpClient) { }

  login(email: string, password: string): Observable<any> {
    return this.http.post<any>(`${this.apiUrl}/user/login`, { email: email, hash: password }, { headers: { 'Content-Type': 'application/json' } }).pipe(
      map(response => {
        if (response && response.token) {
          localStorage.setItem('access_token', response.token);
          localStorage.setItem('authenticated', 'true');
          this.authenticated = true;
          this.authenticated$.emit(true);
        }
        return response;
      })
    );
  }

  register(name: string, email: string, password: string): Observable<any> {
    return this.http.post<any>(`${this.apiUrl}/user/register`, { name: name, email: email, hash: password }, { headers: { 'Content-Type': 'application/json' } });
  }

  isAuthenticated(): boolean {
    return this.authenticated;
  }


}
