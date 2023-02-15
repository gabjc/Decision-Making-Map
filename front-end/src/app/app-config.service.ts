import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({providedIn: 'root'})
export class AppConfigService {
  
  // AppConfigService object has HttpClient object
  constructor(private http: HttpClient) {}

  // Func to retrieve database info
  getDatabase(): Observable<JSON> {

    // Create observer object
    const myObserver = {
      next: (x: JSON) => console.log('Observer got data: ' + x),
      error: (err: Error) => console.error('Observer got an error: ' + err),
      complete: () => console.log('Observer got a complete notification'),
    };

    // Perform GET with observer object
    return this.http.get<any>('http://localhost:9000/database/get')
    
  }
  
}


