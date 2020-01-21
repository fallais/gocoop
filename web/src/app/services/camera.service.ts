import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment';
import { Coop } from '../models/coop';

@Injectable({
  providedIn: 'root'
})
export class CameraService {
  private url: string = environment.baseURL + '/api/v1/cameras';

  constructor(
    private http: HttpClient
  ) {}

  list(): Observable<Map<string, string>> {
    return this.http.get<Map<string, string>>(this.url);
  }
}
