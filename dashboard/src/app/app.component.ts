import { Component } from '@angular/core';
import { AwsService } from './aws.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {

  public accountName: string;
  
  constructor(private awsService: AwsService){
    this.awsService.getAccountName().subscribe(data => {
      this.accountName = data;
    }, err => {
      console.log(err)
      this.accountName = ""
    });
  }

}
