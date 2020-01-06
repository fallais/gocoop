import { Component, OnInit } from '@angular/core';
import { CoopService } from '../../../services/coop.service';
import { NotificationsService } from 'angular2-notifications';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {
  coopStatus: string = ""

  constructor(
    private coopService: CoopService,
    private notificationService: NotificationsService
  ) { }

  ngOnInit() {
    this.getStatus();
  }

  getStatus(): void {
    this.coopService.getStatus().subscribe(
      (resp: string) => {
        this.coopStatus = resp;
      },
      err => {
       console.log(err)

        // Notify
        this.notificationService.error('Error while listing the customers', err.error, {
          timeOut: 5000,
          showProgressBar: true,
          pauseOnHover: true,
          clickToClose: false,
          clickIconToClose: true
        });
      });
  }

  updateStatus(status): void {
    var s = {
      status: status
    }

    this.coopService.updateStatus(s).subscribe(
      (resp: string) => {
        this.coopStatus = resp;
      },
      err => {
       console.log(err)

        // Notify
        this.notificationService.error('Error while listing the customers', err.error, {
          timeOut: 5000,
          showProgressBar: true,
          pauseOnHover: true,
          clickToClose: false,
          clickIconToClose: true
        });
      });
  }

  open(): void {
    this.coopService.open().subscribe(
      (resp: string) => {
        
      },
      err => {
       console.log(err)

        // Notify
        this.notificationService.error('Error while opening the coop', err.error.error_description, {
          timeOut: 5000,
          showProgressBar: true,
          pauseOnHover: true,
          clickToClose: false,
          clickIconToClose: true
        });
      });
  }

  close(): void {
    this.coopService.close().subscribe(
      (resp: string) => {
        
      },
      err => {
       console.log(err)

        // Notify
        this.notificationService.error('Error while closing the coop', err.error, {
          timeOut: 5000,
          showProgressBar: true,
          pauseOnHover: true,
          clickToClose: false,
          clickIconToClose: true
        });
      });
  }
}