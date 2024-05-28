import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http"
import { environment } from '../../environments/environment.development';
import { PaymentDetail } from './payment-detail.model';
import { NgForm } from '@angular/forms';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';
import { MessageService } from '../message.service'
import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'filterForUser'
})
@Injectable({
  providedIn: 'root'
})
export class PaymentDetailService {
    getPaymentDetails() {
        throw new Error('Method not implemented.');
    }

  url: string = environment.apiBaseurl + '/payment';
  list: PaymentDetail[] = [];
  formData: PaymentDetail = new PaymentDetail();
  formSubmitted: boolean = false;
  constructor(
    private http: HttpClient,
    private messageService: MessageService) { }

  refreshList() {
    this.http.get(this.url).subscribe({
      next: (res: any) => { this.list = res as PaymentDetail[]; },
      error: (err: any) => { console.log('An error occurred:', err); }
    });
  }
  postPaymentDetail() {
    return this.http.post(this.url, this.formData);
  }
  putPaymentDetail() {
    return this.http.put(this.url + '/' + this.formData.ID, this.formData);
  }
  deletePaymentDetail(id: number) {
    return this.http.delete(this.url + '/' + id);
  }
  resetForm(form: NgForm) {
    form.form.reset()
    this.formData = new PaymentDetail()
    this.formSubmitted = false
  }

  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {

      console.error(error);

      this.log(`${operation} failed: ${error.message}`);

      return of(result as T);
    };
  }
  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };


searchPaymentDetail(term: string): Observable<PaymentDetail[]> {
  if (!term.trim()) {
    // Trả về một mảng rỗng nếu không có từ khóa tìm kiếm
    return of([]);
  }
  // Gọi API tìm kiếm với từ khóa
  return this.http.get<PaymentDetail[]>(`${this.url}/SearchPaymentDetail?searchString=${term}`).pipe(
    catchError(this.handleError<PaymentDetail[]>('searchPayments', []))
  );
}


  private log(message: string) {
    this.messageService.add(`HeroService: ${message}`);
  }
}


