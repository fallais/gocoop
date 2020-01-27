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
          latitude: [this.coop.latitude, Validators.required],
          longitude: [this.coop.longitude, Validators.required],
        });
      },
      err => {
        console.log(err)
      });
  }

  get formControls() { return this.coopForm.controls; }

  updateStatus(status): void {
    var s = {
      status: status
    }

    this.coopService.updateStatus(s).subscribe(
      (resp: string) => {
        // Notify
        this.notificationService.success('Successfully updated the status', '', {
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
    if (this.coopForm.invalid) {
      return;
    }

    console.log(this.coopForm.value)
  }
}
