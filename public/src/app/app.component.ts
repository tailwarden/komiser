import { Component, AfterViewInit } from '@angular/core';
import { AWSService } from '@app/services';

import * as Datamap from 'datamaps';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements AfterViewInit {
  public costAndUsageChartData:Array<any> = [];
  public costAndUsageChartLabels:Array<any> = [];
  public costAndUsageChartOptions:any = {
    responsive: true,
    scales: {
                    xAxes: [{
                        display: !0,
                        gridLines: {
                            display: !1
                        }
                    }],
                    yAxes: [{
                        ticks: {
                            max: 40,
                            min: 0,
                            stepSize: .5
                        },
                        display: !1,
                        gridLines: {
                            display: !1
                        }
                    }]
                },
                legend: {
                    display: !1
                }
  };
  public costAndUsageChartType:string = 'line';

  public currentBill: number;
  public billUnit : string;
  public currentDate: string;
  
  public currentVPC: number = 0;
  public currentACL: number = 0;
  public currentSecurityGroup: number = 0;
  public currentNatGateway: number = 0;
  public currentInternetGateway: number = 0;
  public currentElasticIP: number = 0;
  public currentKeyPair: number = 0;
  public currentAutoscalingGroup: number = 0;
  public currentRouteTable: number = 0;
  public currentDynamoDBReadCapacity: number = 0;
  public currentDynamoDBWriteCapacity: number = 0;
  public currentDynamoDBTable: number = 0;
  public currentEBSVolumes: number = 0;
  public currentEBSSize: number = 0;
  public currentSnapshot: number = 0;
  public currentSnapshotSize: number = 0;
  public currentStoppedInstances: number = 0;
  public currentRunningInstances: number = 0;
  public currentTerminatedInstances: number = 0;
  public currentCSharpLambdaFunctions: number = 0;
  public currentJavaLambdaFunctions: number = 0;
  public currentGolangLambdaFunctions: number = 0;
  public currentPythonLambdaFunctions: number = 0;
  public currentNodeJSLambdaFunctions: number = 0;
  public currentS3Buckets: number = 0;
  public currentSQSQueues: number = 0;
  public currentSNSTopics: number = 0;
  public currentHostedZones: number = 0;
  public currentIAMRoles: number = 0;
  public currentIAMPolicies: number = 0;
  public currentIAMGroups: number = 0;
  public currentIAMUsers: number = 0;
  public currentOKStateAlarms: number = 0;
  public currentAlarmStateAlarms: number = 0;
  public currentInsufficientDataStateAlarms: number = 0;
  public currentCloudFrontDistributions: number = 0;
  public currentECSClusters: number = 0;
  public currentECSTasks: number = 0;
  public currentECSServices: number = 0;

  public errors: string[] = [];

  public ec2FamilliesChartLabels:string[] = [];
  public ec2FamilliesChartData:number[] = [];
  public ec2FamilliesChartLegend:boolean = true;
  public ec2FamilliesOptions: any = {
    responsive: true
  }
 
  public ec2FamilliesChartType:string = 'polarArea';

  public doughnutChartOptions: any = {
    cutoutPercentage: 80
  }

  public barChartOptions:any = {
    responsive: true,
    scales: {
                xAxes: [{
                    display: !1
                }],
                yAxes: [{
                    display: !1
                }]
            },
            legend: {
                display: !1
            }
  };
  public elbFamiliesChartLabels:string[] = [];
  public elbFamiliesChartType:string = 'bar';
  public elbFamiliesChartData:any[] = [];

  public ebsFamiliesChartLabels:string[] = [];
  public ebsFamiliesChartData:number[] = [];
  public ebsFamiliesChartType:string = 'doughnut';
  public ebsFamiliesChartOptions: any = {
    cutoutPercentage: 80,
    legend: {
       display: 0
    }
  }

  private regions = {
    us_east_1 : {
      latitude : 39.020812,
      longitude : -77.433357 
    },
    us_east_2 : {
      latitude : 40.4172871,
      longitude : -82.907123
    },
    us_west_1 : {
      latitude : 36.778261,
      longitude : -119.4179324
    },
    us_west_2 : {
      latitude : 43.8041334,
      longitude : -120.5542012
    },
    ca_central_1 : {
      latitude : 51.253775,
      longitude : -85.323214
    },
    eu_central_1 : {
      latitude : 50.1109221,
      longitude : 8.6821267
    },
    eu_west_1 : {
      latitude : 53.4058314,
      longitude : -6.0624418
    },
    eu_west_2 : {
      latitude : 51.5073509,
      longitude : -0.1277583
    },
    eu_west_3 : {
      latitude : 48.856614,
      longitude : 2.3522219
    },
    ap_northeast_1 : {
      latitude : 35.6894875,
      longitude : 139.6917064
    },
    ap_northeast_2 : {
      latitude : 37.566535,
      longitude : 126.9779692
    },
    ap_northeast_3 : {
      latitude : 34.6937378,
      longitude : 135.5021651
    },
    ap_southeast_1 : {
      latitude : 1.3553794,
      longitude : 103.8677444
    },
    ap_southeast_2 : {
      latitude : -33.8688197,
      longitude : 151.2092955
    },
    ap_south_1 : {
      latitude : 19.0759837,
      longitude : 72.8776559
    },
    sa_east_1 : {
      latitude : -23.5505199,
      longitude : -46.6333094
    }
  }

  constructor(private awsService: AWSService){
    this.getCurrentVPCs()
    this.getCurrentACLs()
    this.getCurrentSecurityGroups()
    this.getCurrentNatGateways()
    this.getCurrentInternetGateways()
    this.getCurrentElasticIPs()
    this.getCurrentKeyPairs()
    this.getCurrentAutoscalingGroups()
    this.getCurrentRouteTables()
    this.getCurrentDynamoDBTables()
    this.getCurrentEBSVolumes()
    this.getCurrentSnapshots()
    this.getCostAndUsage()
    this.getCurrentLambdaFunctions()
    this.getCurrentElasticLoadBalancers()
    this.getCurrentSQSQueues()
    this.getCurrentSNSTopics()
    this.getCurrentHostedZones()
    this.getCurrentCloudwatchAlarms()
    this.getCurrentIAMRoles()
    this.getCurrentIAMGroups()
    this.getCurrentIAMPolicies()
    this.getCurrentIAMUsers()
    this.getCurrentCloudFrontDistributions()
    this.getCurrentS3Buckets()
  }

  private getCurrentVPCs(): void {
    this.awsService.getCurrentVPCs().subscribe(current => {
      this.currentVPC = (current ? current : 0)
    }, msg => {
      this.errors.push(msg)
    })
  }

  private getCurrentACLs(): void {
    this.awsService.getCurrentACLs().subscribe(current => {
      this.currentACL = (current ? current : 0)
    }, msg => {
      this.errors.push(msg)
    })
  }

  private getCurrentSecurityGroups(): void {
    this.awsService.getCurrentSecurityGroups().subscribe(current => {
      this.currentSecurityGroup = (current ? current : 0)
    }, msg => {
      this.errors.push(msg)
    })
  }

  private getCurrentNatGateways(): void {
    this.awsService.getCurrentNatGateways().subscribe(current => {
      this.currentNatGateway = (current ? current : 0)
    }, msg => {
      this.errors.push(msg)
    })
  }

  private getCurrentInternetGateways(): void {
    this.awsService.getCurrentInternetGateways().subscribe(current => {
      this.currentInternetGateway = (current ? current : 0)
    }, msg => {
      this.errors.push(msg)
    })
  }

  private getCurrentElasticIPs(): void {
    this.awsService.getCurrentElasticIPs().subscribe(current => {
      this.currentElasticIP = (current ? current : 0)
    }, msg => {
      this.errors.push(msg)
    })
  }

  private getCurrentKeyPairs(): void {
    this.awsService.getCurrentKeyPairs().subscribe(current => {
      this.currentKeyPair = (current ? current : 0)
    }, msg => {
      this.errors.push(msg)
    })
  }

  private getCurrentAutoscalingGroups(): void {
    this.awsService.getCurrentAutoscalingGroups().subscribe(current => {
      this.currentAutoscalingGroup = (current ? current : 0)
    }, msg => {
      this.errors.push(msg)
    })
  }

  private getCurrentRouteTables(): void {
    this.awsService.getCurrentRouteTables().subscribe(current => {
      this.currentRouteTable = (current ? current : 0)
    }, msg => {
      this.errors.push(msg)
    })
  }

  private getCurrentDynamoDBTables(): void {
    this.awsService.getCurrentDynamoDBTables().subscribe(current => {
      this.currentDynamoDBTable = (current.total ? current.total : 0)
      this.currentDynamoDBReadCapacity = (current.throughput.readCapacity ? current.throughput.readCapacity : 0)
      this.currentDynamoDBWriteCapacity = (current.throughput.writeCapacity ? current.throughput.writeCapacity : 0)
    }, msg => {
      this.errors.push(msg)
    })
  }

  private getCurrentEBSVolumes(): void {
    this.currentEBSVolumes = 0;
    this.awsService.getCurrentEBSVolumes().subscribe(current => {
      this.currentEBSSize = (current.total ? current.total : 0)
      for(var ebs in current.family){
        this.currentEBSVolumes += current.family[ebs]
        this.ebsFamiliesChartData.push(current.family[ebs])
        this.ebsFamiliesChartLabels.push(ebs)
      }
    }, msg => {
      this.errors.push(msg)
    })
  }

  private getCurrentSnapshots(): void {
    this.awsService.getCurrentSnapshots().subscribe(current => {
      this.currentSnapshot = (current.total ? current.total : 0)
      this.currentSnapshotSize = (current.size ? current.size : 0)
    }, msg => {
      this.errors.push(msg)
    })
  }

  private getCurrentS3Buckets(): void {
    this.awsService.getCurrentS3Buckets().subscribe(current => {
      this.currentS3Buckets = (current ? current : 0)
    }, msg => {
      this.errors.push(msg)
    })
  }

  private getCurrentSQSQueues(): void {
    this.awsService.getCurrentSQSQueues().subscribe(current => {
      this.currentSQSQueues = (current ? current : 0)
    }, msg => {
      this.errors.push(msg)
    })
  }

  private getCurrentSNSTopics(): void {
    this.awsService.getCurrentSNSTopics().subscribe(current => {
      this.currentSNSTopics = (current ? current : 0)
    }, msg => {
      this.errors.push(msg)
    })
  }

  private getCurrentHostedZones(): void {
    this.awsService.getCurrentHostedZones().subscribe(current => {
      this.currentHostedZones = (current ? current : 0)
    }, msg => {
      this.errors.push(msg)
    })
  }

  private getCurrentIAMRoles(): void {
    this.awsService.getCurrentIAMRoles().subscribe(current => {
      this.currentIAMRoles = (current ? current : 0)
    }, msg => {
      this.errors.push(msg)
    })
  }

  private getCurrentIAMPolicies(): void {
    this.awsService.getCurrentIAMPolicies().subscribe(current => {
      this.currentIAMPolicies = (current ? current : 0)
    }, msg => {
      this.errors.push(msg)
    })
  }

  private getCurrentIAMGroups(): void {
    this.awsService.getCurrentIAMGroups().subscribe(current => {
      this.currentIAMGroups = (current ? current : 0)
    }, msg => {
      this.errors.push(msg)
    })
  }

  private getCurrentIAMUsers(): void {
    this.awsService.getCurrentIAMUsers().subscribe(current => {
      this.currentIAMUsers = (current ? current : 0)
    }, msg => {
      this.errors.push(msg)
    })
  }

  private getCurrentCloudwatchAlarms(): void {
    this.awsService.getCurrentCloudwatchAlarms().subscribe(current => {
      this.currentOKStateAlarms = (current.OK ? current.OK : 0)
      this.currentAlarmStateAlarms = (current.ALARM ? current.ALARM : 0)
      this.currentInsufficientDataStateAlarms = (current.INSUFFICIENT_DATA ? current.INSUFFICIENT_DATA : 0)
    }, msg => {
      this.errors.push(msg)
    })
  }

  private getCurrentCloudFrontDistributions(): void {
    this.awsService.getCurrentCloudFrontDistributions().subscribe(current => {
      this.currentCloudFrontDistributions = (current ? current : 0)
    }, msg => {
      this.errors.push(msg)
    })
  }

  private getCurrentECS(): void {
    this.awsService.getCurrentECS().subscribe(current => {
      this.currentECSClusters = (current.clusters ? current.clusters : 0)
      this.currentECSTasks = (current.tasks ? current.tasks : 0)
      this.currentECSServices = (current.services ? current.services : 0)
    }, msg => {
      this.errors.push(msg)
    })
  }

  private getCurrentLambdaFunctions(): void {
    this.currentGolangLambdaFunctions = 0
    this.currentNodeJSLambdaFunctions = 0
    this.currentJavaLambdaFunctions = 0
    this.currentPythonLambdaFunctions = 0
    this.currentCSharpLambdaFunctions = 0
    this.awsService.getCurrentLambdaFunctions().subscribe(res => {
      for(var runtime in res){
        if(runtime.startsWith('go')){
          this.currentGolangLambdaFunctions+=res[runtime]
        }
        if(runtime.startsWith('nodejs')){
          this.currentNodeJSLambdaFunctions+=res[runtime]
        }
        if(runtime.startsWith('java')){
          this.currentJavaLambdaFunctions+=res[runtime]
        }
        if(runtime.startsWith('python')){
          this.currentPythonLambdaFunctions+=res[runtime]
        }
        if(runtime.startsWith('dotnet')){
          this.currentCSharpLambdaFunctions+=res[runtime]
        }
      }
    }, msg => {
      this.errors.push(msg)
    })
  }

  private getCurrentElasticLoadBalancers(): void {
    let data = []
    let labels = []
    this.awsService.getCurrentElasticLoadBalancers().subscribe(res => {
      for(var i in res){
        data.push(res[i])
        labels.push(i)
      }
      this.elbFamiliesChartData = [
          {
            label: "Total",
            borderWidth: 1,
            data: data
          }
      ]
      this.elbFamiliesChartLabels = labels
    }, msg => {
      this.errors.push(msg)
    })
  }

  private getCurrentEC2Instances(): void {
    var map = new Datamap({
      scope: 'world',
      responsive: true,
      element: document.getElementById('map'),
      geographyConfig: {
        popupOnHover: false,
        highlightOnHover: false
      },
      fills: {
        defaultFill: '#ABDDA4',
        instance: 'red'
      }
    });
    let data = []
    this.awsService.getCurrentEC2Instances().subscribe(current => {
      this.currentStoppedInstances = (current.state.stopped ? current.state.stopped : 0)
      this.currentTerminatedInstances = (current.state.terminated ? current.state.terminated : 0)
      this.currentRunningInstances = (current.state.running ? current.state.running : 0)
      for(var i in current.family){
        this.ec2FamilliesChartLabels.push(i)
        this.ec2FamilliesChartData.push(current.family[i])
      }
      for(var region in current.region){
        var params = this.regions[region.split("-").join("_")]
        data.push({
          radius: current.region[region],
          latitude: params.latitude,
          longitude: params.longitude,
          region: region,
          fillKey: 'instance',
          instances: current.region[region]
        })
      }
      map.bubbles(data, {
        popupTemplate: function(geo, data){
          return `<div class="hoverinfo">${data.region}: ${data.instances} EC2`
        }
      })
    }, msg => {
      this.errors.push(msg)
    })
  }


  private getCostAndUsage(): void {
    var values = []
    var labels = []
    this.awsService.getBilling().subscribe(res => {
      this.currentBill = res[res.length - 1].Amount.toFixed(2)
      this.billUnit = res[res.length - 1].Unit;
      this.currentDate = `${(new Date().toLocaleString("en-us", { month: "long" }))} ${(new Date()).getFullYear()}`
      res.forEach(item => {
        values.push(item.Amount.toFixed(3))
        labels.push(new Date(item.End).toLocaleString("en-us", { month: "long" }))
      })
      this.costAndUsageChartData = [
        {
          label: "Bill",
          fill: 0,
          lineTension: 0,
          backgroundColor: "transparent",
          borderColor: "#6ccef0",
          pointBorderColor: "#59c2e6",
          pointHoverBackgroundColor: "#59c2e6",
          borderCapStyle: "butt",
          borderDash: [],
          borderDashOffset: 0,
          borderJoinStyle: "miter",
          borderWidth: 3,
          pointBackgroundColor: "#59c2e6",
          pointBorderWidth: 0,
          pointHoverRadius: 4,
          pointHoverBorderColor: "#fff",
          pointHoverBorderWidth: 0,
          pointRadius: 4,
          pointHitRadius: 0,
          data: values,
          spanGaps: 1
        }
      ]
      this.costAndUsageChartLabels = labels
    }, msg => {
      this.errors.push(msg)
    }) 
  }

  ngAfterViewInit(){
    this.getCurrentEC2Instances();
  }
}
