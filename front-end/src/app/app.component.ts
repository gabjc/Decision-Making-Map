import { Component, OnInit } from '@angular/core';

// Import created services
import { StatusService } from './shared/status.service';
import { AppConfigService } from './app-config.service';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'go-angular';
  status = 'DOWN';

  // App component object has StatusService object and AppConfigService object
  constructor(private statusService: StatusService, private appConfigService: AppConfigService, private http: HttpClient) { }

  // TODO: mock services for unit testing parts of the frontend

  // On app init
  ngOnInit() {
    // Log simple get route from backend
    this.http.get('http://localhost:9000/user/get/admin').subscribe(res => {
      console.log('res', res)
    })
  }

}
