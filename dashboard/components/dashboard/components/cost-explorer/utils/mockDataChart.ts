const dummyStats = {
  regions: 17,
  resources: 257,
  users: 0,
  cost: [
    { date: '2023-01-05T21:09:56.176672794Z', amount: 74 },
    { date: '2022-08-05T21:09:56.176673175Z', amount: 120 },
    { date: '2022-09-05T21:09:56.176674426Z', amount: 140 },
    { date: '2022-10-05T21:09:56.176674926Z', amount: 200 },
    { date: '2022-11-05T21:09:56.176675116Z', amount: 60 },
    { date: '2022-12-05T21:09:56.176675608Z', amount: 120 },
    { date: '2023-01-05T21:09:56.176675756Z', amount: 20 }
  ]
};

const dummyData = [
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

const mockDataChart = [dummyStats, dummyData];

export default mockDataChart;
