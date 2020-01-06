import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders  } from '@angular/common/http';
import { JwtHelperService } from '@auth0/angular-jwt';
import { environment } from '../../environments/environment';

@Injectable()
export class AuthService {
  private url: string = environment.baseURL + '/api/v1/login';
  private helper = new JwtHelperService();

  constructor(
    private http: HttpClient
  ) {}

  getToken(): string {
    return localStorage.getItem(environment.tokenName);
  }

  setToken(token: string): void {
    localStorage.setItem(environment.tokenName, token);
  }

  isAuthenticated(): boolean {
    if (localStorage.getItem(environment.tokenName) == null) {
      return false;
    }

    if (this.isTokenExpired(localStorage.getItem(environment.tokenName))) {
      return false;
    }

    return true;
  }

  getTokenExpirationDate(): Date {
    return this.helper.getTokenExpirationDate(this.getToken());
  }

  isTokenExpired(token?: string): boolean {
    return this.helper.isTokenExpired(token);
  }

  login(user) {
    return this.http.post(this.url, user);
  }

  logout() {
    localStorage.removeItem(environment.tokenName);
  }
}
