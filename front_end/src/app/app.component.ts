import { Component, OnInit } from '@angular/core';

// Import status.service.ts file created
import { StatusService } from './shared/status.service';
// Import app-config.service file created
import { AppConfigService } from './app-config.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'go-angular';
  status = 'DOWN';

  // App component object has StatusService object and AppConfigService object
  constructor(private statusService: StatusService, private appConfigService: AppConfigService) { }

  // On app init
  ngOnInit() {

    // Get status with statusService object
    /* this.statusService
      .getStatus()
      .then((result: any) => {
        this.status = result.status;
      }); */

    // Get database info with appConfigService object
    var databaseObservable = this.appConfigService.getDatabase()
    console.log(databaseObservable);          // LOGGING OBSERVABLE, FIGURE OUT HOW TO GET DATA IN APP-CONFIG & RETURN CLEANER DATA
  }

}
