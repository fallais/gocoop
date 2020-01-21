import { Pipe, PipeTransform } from '@angular/core';
import { WeatherResponse } from '../models/weather';

@Pipe({
  name: 'weatherIcon'
})
export class WeatherIconPipe implements PipeTransform {

  transform(value: WeatherResponse, ...args: any[]): any {
    if (value.weather.length <= 0) {
      return ""
    }

    switch(value.weather[0].main){
      case "Clear":
        return "day-sunny";
      case "Rain":
      case "Drizzle":
        return "day-rain";
      case "Thunderstorm":
        return "wi-day-thunderstorm";
      case "Snow":
        return "day-snow";
      case "Clouds":
        return "day-cloudy";
      case "Atmosphere":
        return "day-haze";
      case "Extreme":
        return "";
      default:
        return "";
    }
  }

}
