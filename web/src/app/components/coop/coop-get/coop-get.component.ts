import { Component, OnInit } from '@angular/core';
import { CoopService } from '../../../services/coop.service';
import { NotificationsService } from 'angular2-notifications';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-coop-get',
  templateUrl: './coop-get.component.html',
  styleUrls: ['./coop-get.component.css']
})
export class CoopGetComponent implements OnInit {
  coop: any;
  coopForm: FormGroup;

  constructor(
    private coopService: CoopService,
    private notificationService: NotificationsService,
    private formBuilder: FormBuilder,
  ) { }

  ngOnInit() {
    this.coopService.get().subscribe(
      (coop: any) => {
        this.coop = coop;

        this.coopForm = this.formBuilder.group({
          opening_condition: this.formBuilder.group({
            mode: [this.coop.opening_condition.mode, [Validators.required]],
            value: [this.coop.opening_condition.value, [Validators.required]]
          }),
          closing_condition: this.formBuilder.group({
            mode: [this.coop.closing_condition.mode, [Validators.required]],
            value: [this.coop.closing_condition.value, [Validators.required]]
          }),
        });
      },
      err => {
        // Notify
        this.notificationService.error('Error while getting the coop', err.error.error_description, {
          timeOut: 5000,
          showProgressBar: true,
          pauseOnHover: true,
          clickToClose: false,
          clickIconToClose: true
        });
      });
  }

  get formControls() { return this.coopForm.controls; }

  setOpened(): void {
    var input = this.coop;
    input.status = "opened"
    
    this.coopService.update(input).subscribe(
      (resp: string) => {
        // Notify
        this.notificationService.success('Successfully updated the coop', '', {
          timeOut: 5000,
          showProgressBar: true,
          pauseOnHover: true,
          clickToClose: false,
          clickIconToClose: true
        });
      },
      err => {
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

  setClosed(): void {
    var input = this.coop;
    input.status = "closed"
    
    this.coopService.update(input).subscribe(
      (resp: string) => {
        // Notify
        this.notificationService.success('Successfully updated the coop', '', {
          timeOut: 5000,
          showProgressBar: true,
          pauseOnHover: true,
          clickToClose: false,
          clickIconToClose: true
        });
      },
      err => {
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

  enableAutomaticMode(): void {
    var input = this.coop;
    input.is_automatic = true
    
    this.coopService.update(input).subscribe(
      (resp: string) => {
        // Notify
        this.notificationService.success('Successfully updated the coop', '', {
          timeOut: 5000,
          showProgressBar: true,
          pauseOnHover: true,
          clickToClose: false,
          clickIconToClose: true
        });
      },
      err => {
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

  disableAutomaticMode(): void {
    var input = this.coop;
    input.is_automatic = false
    
    this.coopService.update(input).subscribe(
      (resp: string) => {
        // Notify
        this.notificationService.success('Successfully updated the coop', '', {
          timeOut: 5000,
          showProgressBar: true,
          pauseOnHover: true,
          clickToClose: false,
          clickIconToClose: true
        });
      },
      err => {
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

  onSubmit(): void {
    // Add the status
    this.coopForm.value.status = this.coop.status;

    this.coopService.update(this.coopForm.value).subscribe(
      (resp: string) => {
        // Notify
        this.notificationService.success('Successfully updated the coop', '', {
          timeOut: 5000,
          showProgressBar: true,
          pauseOnHover: true,
          clickToClose: false,
          clickIconToClose: true
        });
      },
      err => {
        // Notify
        this.notificationService.error('Error while updating the coop', err.error.error_description, {
          timeOut: 5000,
          showProgressBar: true,
          pauseOnHover: true,
          clickToClose: false,
          clickIconToClose: true
        });
      });

    console.log(this.coopForm.value)
  }
}
