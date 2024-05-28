import { Component, OnInit } from '@angular/core';

import { Observable, Subject } from 'rxjs';
import {
  debounceTime, distinctUntilChanged, switchMap
} from 'rxjs/operators';
import { PaymentDetailService } from '../shared/payment-detail.service'
import { PaymentDetail } from '../shared/payment-detail.model';

@Component({
  selector: 'app-payment-search',
  templateUrl: './payment-search.component.html',
  styleUrl: './payment-search.component.css'
})
export class PaymentSearchComponent {
  payment$!: Observable<PaymentDetail[]>
  private searchTerms = new Subject<string>();
  constructor(private paymentService: PaymentDetailService) { }
  search(term: string): void {
    this.searchTerms.next(term);
  }

  ngOnInit(): void {
    this.payment$ = this.searchTerms.pipe(
      debounceTime(300),

      distinctUntilChanged(),

      switchMap((term: string) => this.paymentService.searchPaymentDetail(term)),
    );
  }
}
