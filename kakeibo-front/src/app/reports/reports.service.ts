import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { EMPTY, Observable } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { Report } from '../models/report.model';
import { ReportRequest } from '../models/reportRequest.model';

@Injectable({
  providedIn: 'root'
})
export class ReportsService {

  constructor(private http: HttpClient) { }

  postReport(req: ReportRequest): Observable<Report> {
    return this.http.post<Report>('api/report/', req)
    .pipe( 
      catchError(() =>{
        return EMPTY
      })
    )
  }
}
