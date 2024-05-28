import { Component, OnInit } from '@angular/core';
import { PaymentDetail } from '../shared/payment-detail.model';
import { PaymentDetailService } from '../shared/payment-detail.service';
import { ToastrService } from 'ngx-toastr';
import { Subject } from 'rxjs';


@Component({
  selector: 'app-payment-details',
  templateUrl: './payment-details.component.html',
  styleUrls: ['./payment-details.component.css']
})

export class PaymentDetailsComponent implements OnInit {
  searchKeyword: string = '';
  searchTerms = new Subject<string>();
  constructor(public service: PaymentDetailService , private toastr: ToastrService) { }

  ngOnInit(): void {
    this.service.refreshList();
  }

  populateForm(selectedRecord: PaymentDetail) {
    this.service.formData = Object.assign({}, selectedRecord);
  }

  onDelete(id: number): void {
    if(confirm("Are you sure to delete this Record?"))
    this.service.deletePaymentDetail(id)
      .subscribe({
        next: (res: any) => {
          this.service.list = res as PaymentDetail[]
          this.toastr.error('Deleted successfully', 'Payment Detail Register')
        },
        error: (err: any) => { console.log(err) }
      })
  }
  onSearch(): void {
    this.searchTerms.next(this.searchKeyword);
  }

}
