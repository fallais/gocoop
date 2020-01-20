import { Component, OnInit } from '@angular/core';
import { CoopService } from '../../../services/coop.service';
import { NotificationsService } from 'angular2-notifications';
import { Coop } from '../../../models/coop';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {
  coopStatus: string = ""
  coop: Coop;

  constructor(
    private coopService: CoopService,
    private notificationService: NotificationsService
  ) { }

  ngOnInit() {
    this.getStatus();
    this.get();
  }

  get(): void {
    this.coopService.get().subscribe(
      (resp: Coop) => {
        this.coop = resp;
      },
      err => {
        // Notify
        this.notificationService.error('Error while getting the coop configuration', err.error.error_description, {
          timeOut: 5000,
          showProgressBar: true,
          pauseOnHover: true,
          clickToClose: false,
          clickIconToClose: true
        });
      });
  }

  getStatus(): void {
    this.coopService.getStatus().subscribe(
      (resp: string) => {
        this.coopStatus = resp;
      },
      err => {
        // Notify
        this.notificationService.error('Error while getting the status', err.error.error_description, {
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
        this.getStatus();
      },
      err => {
       console.log(err)

        // Notify
        this.notificationService.error('Error while updating the status', err.error.error_description, {
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
        // Notify
        this.notificationService.success('Door is opening', '', {
          timeOut: 5000,
          showProgressBar: true,
          pauseOnHover: true,
          clickToClose: false,
          clickIconToClose: true
        });
      },
      err => {
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
        // Notify
        this.notificationService.success('Door is closing', '', {
          timeOut: 5000,
          showProgressBar: true,
          pauseOnHover: true,
          clickToClose: false,
          clickIconToClose: true
        });
      },
      err => {
        // Notify
        this.notificationService.error('Error while closing the coop', err.error.error_description, {
          timeOut: 5000,
          showProgressBar: true,
          pauseOnHover: true,
          clickToClose: false,
          clickIconToClose: true
        });
      });
  }
}