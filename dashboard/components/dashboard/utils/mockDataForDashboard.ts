import { GlobalData } from '../../layout/context/GlobalAppContext';
import { DashboardCloudMapRegions } from '../components/cloud-map/hooks/useCloudMap';
import { DashboardCostExplorerData } from '../components/cost-explorer/hooks/useCostExplorer';
import { ResourcesManagerChartProps } from '../components/resources-manager/hooks/useResourcesManagerChart';

const stats: GlobalData = {
  regions: 17,
  resources: 255,
  accounts: 4,
  costs: 60
};

const regions: DashboardCloudMapRegions = [
  {
    name: 'Ohio',
    label: 'us-east-2',
    latitude: '40.367474',
    longitude: '-82.996216',
    resources: 3
  },
  {
    name: 'N.Virginia',
    label: 'us-east-1',
    latitude: '37.926868',
    longitude: '-78.024902',
    resources: 119
  },
  {
    name: 'N.California',
    label: 'us-west-1',
    latitude: '36.778261',
    longitude: '-119.4179324',
    resources: 3
  },
  {
    name: 'Oregon',
    label: 'us-west-2',
    latitude: '45.523062',
    longitude: '-122.676482',
    resources: 3
  },
  {
    name: 'Cape Town',
    label: 'af-south-1',
    latitude: '-33.924869',
    longitude: '18.424055',
    resources: 0
  },
  {
    name: 'Hong Kong',
    label: 'ap-east-1',
    latitude: '22.302711',
    longitude: '114.177216',
    resources: 0
  },
  {
    name: 'Jakarta',
    label: 'ap-southeast-3',
    latitude: '-6.2087634',
    longitude: '106.816666',
    resources: 0
  },
  {
    name: 'Mumbai',
    label: 'ap-south-1',
    latitude: '19.076090',
    longitude: '72.877426',
    resources: 3
  },
  {
    name: 'Osaka',
    label: 'ap-northeast-3',
    latitude: '34.6937378',
    longitude: '135.5021651',
    resources: 3
  },
  {
    name: 'Seoul',
    label: 'ap-northeast-2',
    latitude: '37.566535',
    longitude: '126.9779692',
    resources: 3
  },
  {
    name: 'Singapore',
    label: 'ap-southeast-1',
    latitude: '1.290270',
    longitude: '103.851959',
    resources: 3
  },
  {
    name: 'Sydney',
    label: 'ap-southeast-2',
    latitude: '-33.8667',
    longitude: '151.206990',
    resources: 3
  },
  {
    name: 'Tokyo',
    label: 'ap-northeast-1',
    latitude: '35.652832',
    longitude: '139.839478',
    resources: 3
  },
  {
    name: 'Canada',
    label: 'ca-central-1',
    latitude: '56.130367',
    longitude: '-106.346771',
    resources: 3
  },
  {
    name: 'Frankfurt',
    label: 'eu-central-1',
    latitude: '50.1109221',
    longitude: '8.6821267',
    resources: 91
  },
  {
    name: 'Ireland',
    label: 'eu-west-1',
    latitude: '53.350140',
    longitude: '-6.266155',
    resources: 3
  },
  {
    name: 'London',
    label: 'eu-west-2',
    latitude: '51.5073509',
    longitude: '-0.1277583',
    resources: 3
  },
  {
    name: 'Milan',
    label: 'eu-south-1',
    latitude: '45.4654219',
    longitude: '9.1859243',
    resources: 0
  },
  {
    name: 'Paris',
    label: 'eu-west-3',
    latitude: '48.864716',
    longitude: '2.352222',
    resources: 3
  },
  {
    name: 'Stockholm',
    label: 'eu-north-1',
    latitude: '59.334591',
    longitude: '18.063240',
    resources: 3
  },
  {
    name: 'Bahrain',
    label: 'me-south-1',
    latitude: '26.066700',
    longitude: '50.557700',
    resources: 0
  }
];

const resources: ResourcesManagerChartProps[] = [
  { label: 'AWS', total: 451 },
  { label: 'Azure', total: 153 },
  { label: 'GCP', total: 259 },
  { label: 'OVH', total: 100 }
];

const costs: DashboardCostExplorerData[] = [
  {
    date: '2022-03',
    datapoints: [
      {
        name: 'Amazon Elastic Load Balancer',
        amount: 85.1
      },
      {
        name: 'AWS Lambda',
        amount: 45.1
      },
      {
        name: 'Amazon API Gateway',
        amount: 25.1
      },
      {
        name: 'Amazon SQS',
        amount: 75.1
      },
      {
        name: 'Others',
        amount: 15.1
      }
    ]
  },
  {
    date: '2022-04',
    datapoints: [
      {
        name: 'Amazon Elastic Load Balancer',
        amount: 85.1
      },
      {
        name: 'AWS Lambda',
        amount: 35.1
      },
      {
        name: 'Amazon API Gateway',
        amount: 7.1
      },
      {
        name: 'Amazon SQS',
        amount: 25.1
      },
      {
        name: 'Others',
        amount: 65.1
      }
    ]
  },
  {
    date: '2022-05',
    datapoints: [
      {
        name: 'Amazon Elastic Load Balancer',
        amount: 105.1
      },
      {
        name: 'AWS Lambda',
        amount: 40.1
      },
      {
        name: 'Amazon API Gateway',
        amount: 17.6
      },
      {
        name: 'Amazon SQS',
        amount: 25.1
      },
      {
        name: 'Others',
        amount: 65.1
      }
    ]
  },
  {
    date: '2022-06',
    datapoints: [
      {
        name: 'Amazon Elastic Load Balancer',
        amount: 99.4
      },
      {
        name: 'AWS Lambda',
        amount: 40.1
      },
      {
        name: 'Amazon API Gateway',
        amount: 37.6
      },
      {
        name: 'Amazon SQS',
        amount: 25.1
      },
      {
        name: 'Others',
        amount: 65.1
      }
    ]
  },
  {
    date: '2022-07',
    datapoints: [
      {
        name: 'Amazon Elastic Load Balancer',
        amount: 75.1
      },
      {
        name: 'AWS Lambda',
        amount: 54.1
      },
      {
        name: 'Amazon API Gateway',
        amount: 37.6
      },
      {
        name: 'Amazon SQS',
        amount: 5.1
      },
      {
        name: 'Others',
        amount: 109.1
      }
    ]
  },
  {
    date: '2022-08',
    datapoints: [
      {
        name: 'Amazon Elastic Load Balancer',
        amount: 175.1
      },
      {
        name: 'AWS Lambda',
        amount: 84.1
      },
      {
        name: 'Amazon API Gateway',
        amount: 47.6
      },
      {
        name: 'Amazon SQS',
        amount: 15.1
      },
      {
        name: 'Others',
        amount: 39.1
      }
    ]
  }
];

const mockDataForDashboard = { stats, regions, resources, costs };

export default mockDataForDashboard;
