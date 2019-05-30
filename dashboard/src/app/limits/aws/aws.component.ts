import { Component, OnInit,OnDestroy } from '@angular/core';
import { AwsService } from '../../aws.service';
import { StoreService } from '../../store.service';
import { Subject, Subscription } from 'rxjs';

@Component({
  selector: 'aws-limits',
  templateUrl: './aws.component.html',
  styleUrls: ['./aws.component.css']
})
export class AwsLimitsComponent implements OnInit,OnDestroy {

  public serviceLimits: Array<any> = [];

  public loadingServiceLimits: boolean = true;

  private _subscription: Subscription;

  constructor(private awsService: AwsService, private storeService: StoreService) {
    this.initState();

    this._subscription = this.storeService.profileChanged.subscribe(profile => {
      this.serviceLimits = [];
      
      this.loadingServiceLimits = true;

      this.initState();
    });
  }

  private initState(){
    this.awsService.getServiceLimits().subscribe(data => {
      this.serviceLimits = data;
      this.loadingServiceLimits = false;
    }, err => {
      this.serviceLimits = [];
      this.loadingServiceLimits = false;
    });
  }

  ngOnDestroy() {
    this._subscription.unsubscribe();
  }

  public getColor(status: string) {
    switch (status) {
      case 'ok':
        return 'card card-stats card-success';
      case 'warning':
        return 'card card-stats card-warning';
      case 'danger':
        return 'card card-stats card-danger';
      default:
        return 'card card-stats';
    }
  }

  public getServiceLogo(name: string){
    if (name.indexOf('Route 53') != -1) {
      return 'https://cdn.komiser.io/images/services/aws/white/route53.png';
    }
    else if (name.indexOf('EBS') != -1) {
      return 'https://cdn.komiser.io/images/services/aws/white/ebs.png';
    }
    else if (name.indexOf('RDS') != -1) {
      return 'https://cdn.komiser.io/images/services/aws/white/rds.png';
    }
    else if (name.indexOf('DynamoDB') != -1) {
      return 'https://cdn.komiser.io/images/services/aws/white/dynamodb.png';
    }
    else if (name.indexOf('IAM Group') != -1) {
      return 'https://cdn.komiser.io/images/services/aws/white/iam_groups.png';
    }
    else if (name.indexOf('VPC Internet Gateways') != -1) {
      return 'https://cdn.komiser.io/images/services/aws/white/igw.png';
    }
    else if (name.indexOf('IAM Roles') != -1) {
      return 'https://cdn.komiser.io/images/services/aws/white/iam_roles.png';
    }
    else if (name.indexOf('Elastic IP Address') != -1) {
      return 'https://cdn.komiser.io/images/services/aws/white/elastic_ip.png';
    }
    else if (name.indexOf('IAM Instance Profiles') != -1) {
      return 'https://cdn.komiser.io/images/services/aws/white/instance_profiles.png';
    }
    else if (name.indexOf('IAM Users') != -1) {
      return 'https://cdn.komiser.io/images/services/aws/white/iam_users.png';
    }
    else if (name.indexOf('ELB') != -1) {
      return 'https://cdn.komiser.io/images/services/aws/white/elb.png';
    }
    else if (name.indexOf('IAM Policies') != -1) {
      return 'https://cdn.komiser.io/images/services/aws/white/iam_policies.png';
    }
    else if (name.indexOf('CloudFormation') != -1) {
      return 'https://cdn.komiser.io/images/services/aws/white/cloudformation.png';
    }
    else if (name.indexOf('Auto Scaling Groups') != -1) {
      return 'https://cdn.komiser.io/images/services/aws/white/ec2.png';
    }
    else if (name.indexOf('SES') != -1) {
      return 'https://cdn.komiser.io/images/services/aws/white/ses.png';
    } else {
      return 'https://cdn.komiser.io/images/services/aws/white/aws.png';
    }
  }

  ngOnInit() {
  }

}
