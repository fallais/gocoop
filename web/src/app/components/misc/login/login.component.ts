import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { AuthService } from '../../../services/auth.service';
import { Router } from '@angular/router';
import { NotificationsService } from 'angular2-notifications';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  form: FormGroup;
  isSubmitted: boolean = false;

  constructor(
    private fb: FormBuilder,
    private authService: AuthService,
    private router: Router,
    private notificationService: NotificationsService
  ) {
    if (this.authService.isAuthenticated) { 
      this.router.navigate(['/']);
    }
   }

  ngOnInit() {
    this.form = this.fb.group({  
      username: ['', Validators.required],
      password: ['', Validators.required]
    });
  }

  get formControls() { return this.form.controls; }

  onSubmit() {
    this.isSubmitted = true;

    this.authService.login(this.form.value).subscribe(
      res => {
        this.authService.setToken(res['access_token'].toString());
        this.router.navigateByUrl('/dashboard');

        // Notify
        this.notificationService.success("Successfully authenticated", "", {
          timeOut: 5000,
          showProgressBar: true,
          pauseOnHover: true,
          clickToClose: false,
          clickIconToClose: true
        })
      },
      err => {
        // Notify
        this.notificationService.error("Error while authenticating", err.error.error_message, {
          timeOut: 5000,
          showProgressBar: true,
          pauseOnHover: true,
          clickToClose: false,
          clickIconToClose: true
        })
      });
  }

}
