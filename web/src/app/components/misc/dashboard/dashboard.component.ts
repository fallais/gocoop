import { Component, OnInit } from '@angular/core';
import { Subscription, timer, Observable } from 'rxjs';
import { CoopService } from '../../../services/coop.service';
import { NotificationsService } from 'angular2-notifications';
import { WeatherService } from '../../../services/weather.service';
import { CameraService } from '../../../services/camera.service';
import { Coop } from '../../../models/coop';
import { WeatherResponse } from 'src/app/models/weather';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {
  private timer$: Observable<number> = timer(0, 10000);
  private subscription: Subscription = new Subscription();

  coop: Coop;
  weather: WeatherResponse;
  cameras: Map<string, string>;
  hasData: boolean = false;

  constructor(
    private coopService: CoopService,
    private weatherService: WeatherService,
    private cameraService: CameraService,
    private notificationService: NotificationsService
  ) { }

  ngOnInit() {
    this.get();
    this.getWeather();
    this.listCameras();
  }

  listCameras(): void {
    this.cameraService.list().subscribe(resp => {
      this.cameras = resp;
    });
  }

  getWeather(): void {
    this.coopService.get().subscribe(resp => {
      this.weatherService.get(resp.latitude, resp.longitude, "a4e6ca400a6006140999a787fdc13883").subscribe(
        (resp: WeatherResponse) => {
          this.weather = resp;
        },
        err => {
          console.log(err)
        }
      )
    });      
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