import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment';
import { Coop } from '../models/coop';

@Injectable({
  providedIn: 'root'
})
export class CoopService {
  private coopURL: string = environment.baseURL + '/api/v1/coop';

  constructor(
    private http: HttpClient
  ) {}

  get(): Observable<Coop> {
    return this.http.get<Coop>(this.coopURL);
  }

  update(input: any): Observable<string> {
    return this.http.post<string>(this.coopURL, input);
  }

  open(): Observable<undefined> {
    return this.http.post<undefined>(this.coopURL + '/open', null);
  }

  close(): Observable<undefined> {
    return this.http.post<undefined>(this.coopURL + '/close', null);
  }
}
