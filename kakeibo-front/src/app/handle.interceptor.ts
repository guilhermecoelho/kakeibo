import { Injectable } from '@angular/core';
import {
  HttpRequest,
  HttpHandler,
  HttpEvent,
  HttpInterceptor
} from '@angular/common/http';
import { Observable } from 'rxjs';
import { LocalStorageService } from './local-storage.service';
import { Router } from '@angular/router';

@Injectable()
export class HandleInterceptor implements HttpInterceptor {

  constructor(private localStorageService: LocalStorageService, private router: Router) {}

  intercept(request: HttpRequest<unknown>, next: HttpHandler): Observable<HttpEvent<unknown>> {

    if(request.headers.has('ignore')){
      return next.handle(request)
    }
    
    let token = this.localStorageService.Get("token");
    if(token === null){
      this.router.navigate(['/login']);
    }
    
    let modifiedRequest = request.clone({
      headers: request.headers.set(`Authorization`, token)
    })

    return next.handle(modifiedRequest);
  }
}
