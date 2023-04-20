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

  // On app init
  ngOnInit() {

  }

}
