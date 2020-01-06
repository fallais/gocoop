import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class CoopService {
  private coopURL: string = environment.baseURL + '/api/v1/coop';

  constructor(
    private http: HttpClient
  ) {}

  get(): Observable<any> {
    return this.http.get<any>(this.coopURL);
  }

  getStatus(): Observable<string> {
    return this.http.get<string>(this.coopURL + '/status');
  }

  updateStatus(status: any): Observable<string> {
    return this.http.post<string>(this.coopURL + '/status', status);
  }

  open(): Observable<undefined> {
    return this.http.post<undefined>(this.coopURL + '/open', null);
  }

  close(): Observable<undefined> {
    return this.http.post<undefined>(this.coopURL + '/close', null);
  }
}
