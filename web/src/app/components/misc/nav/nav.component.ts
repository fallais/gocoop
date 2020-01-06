import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../../../services/auth.service';

@Component({
  selector: 'app-nav',
  templateUrl: './nav.component.html',
  styleUrls: ['./nav.component.css']
})
export class NavComponent implements OnInit {
  isPrivate: boolean;
  public isMenuCollapsed = true;

  constructor(
    private router: Router, 
    private authService: AuthService
    ) { }

  ngOnInit() {
  }

  isAuthenticated(){
    return this.authService.isAuthenticated();
  }

  onLogout(){
    this.authService.logout()
    this.router.navigate(['login']);
  }

}
