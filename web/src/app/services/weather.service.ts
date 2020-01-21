import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment';
import { WeatherResponse } from '../models/weather'

@Injectable({
  providedIn: 'root'
})
export class WeatherService {
  private url: string = 'https://api.openweathermap.org/data/2.5/weather';

  constructor(
    private http: HttpClient
  ) {}

  get(latitude: number, longitude: number, apiKey: string): Observable<WeatherResponse> {
    let params = new HttpParams().set("lat", latitude.toString()).set("lon", longitude.toString()).set("APPID", apiKey).set('units', 'metric');

    return this.http.get<WeatherResponse>(this.url, {params: params});
  }
}
