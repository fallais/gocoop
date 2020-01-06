import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { HttpErrorResponse } from '@angular/common/http';
import { HttpEvent, HttpInterceptor, HttpHandler, HttpRequest } from '@angular/common/http';
import { Observable } from 'rxjs';
import { tap } from 'rxjs/operators';

@Injectable()
export class JWTInterceptor implements HttpInterceptor {
    constructor(
        private router: Router
    ) { }

    intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
        return next.handle(req).pipe(
          tap(
            event => {},
            err => {
              if (err instanceof HttpErrorResponse && err.status == 401) {
                this.router.navigateByUrl('/login');
              } else if (err instanceof HttpErrorResponse && err.status == 403) {
                this.router.navigate(['/login']);
              }
            })
        );
    }
}