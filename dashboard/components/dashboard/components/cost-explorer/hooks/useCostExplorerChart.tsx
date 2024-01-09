import { ChartData, ChartOptions } from 'chart.js';
import { Dispatch, SetStateAction, useEffect, useState } from 'react';
import formatNumber from '../../../../../utils/formatNumber';
import { dateFormatter } from '../utils/dateHelper';
import {
  CostExplorerQueryDateProps,
  CostExplorerQueryGranularityProps,
  CostExplorerQueryGroupProps,
  DashboardCostExplorerData
} from './useCostExplorer';

type CostExplorerDatapointsProps = {
  name: string;
  amount: number;
};

type CostExplorerDatapoints = {
  date: string;
  datapoints: CostExplorerDatapointsProps[];
};

export type GroupBySelectProps = {
  values: string[];
  displayValues: string[];
};

type useCostExplorerChartProps = {
  data: DashboardCostExplorerData[] | undefined;
  setQueryGroup: Dispatch<SetStateAction<CostExplorerQueryGroupProps>>;
  queryGranularity: CostExplorerQueryGranularityProps;
  setQueryGranularity: Dispatch<
    SetStateAction<CostExplorerQueryGranularityProps>
  >;
  setQueryDate: Dispatch<SetStateAction<CostExplorerQueryDateProps>>;
};

function useCostExplorerChart({
  data,
  setQueryGroup,
  queryGranularity,
  setQueryGranularity,
  setQueryDate
}: useCostExplorerChartProps) {
  const [chartData, setChartData] = useState<ChartData<'bar'>>();

  const colors = [
    'rgba(0, 114, 178, 0.65)',
    'rgba(255, 140, 0, 0.65)',
    'rgba(34, 139, 34, 0.65)',
    'rgba(255, 215, 0, 0.65)',
    'rgba(153, 50, 204, 0.65)',
    'rgba(30, 144, 255, 0.65)',
    'rgba(255, 105, 180, 0.65)',
    'rgba(50, 205, 50, 0.65)',
    'rgba(255, 99, 71, 0.65)',
    'rgba(138, 43, 226, 0.65)',
    'rgba(139, 0, 0, 0.65)',
    'rgba(255, 160, 122, 0.65)',
    'rgba(65, 105, 225, 0.65)',
    'rgba(255, 182, 193, 0.65)',
    'rgba(0, 191, 255, 0.65)',
    'rgba(147, 112, 219, 0.65)',
    'rgba(143, 188, 139, 0.65)',
    'rgba(255, 127, 80, 0.65)',
    'rgba(0, 206, 209, 0.65)',
    'rgba(220, 20, 60, 0.65)',
    'rgba(0, 255, 127, 0.65)',
    'rgba(106, 90, 205, 0.65)',
    'rgba(0, 128, 128, 0.65)',
    'rgba(255, 165, 0, 0.65)',
    'rgba(75, 0, 130, 0.65)'
  ];

  /* To be un-commented when 'view' is supported 
  const groupBySelect: GroupBySelectProps = {
    values: ['provider', 'service', 'region', 'account', 'view'],
    displayValues: [
      'Cloud provider',
      'Cloud service',
      'Cloud region',
      'Cloud account',
      'Custom view'
    ]
  }; */

  const groupBySelect: GroupBySelectProps = {
    values: ['provider', 'service', 'region', 'account','Resource'],
    displayValues: [
      'Cloud provider',
      'Cloud service',
      'Cloud region',
      'Cloud account',
      'Resource'
    ]
  };

  const granularitySelect: GroupBySelectProps = {
    values: ['monthly', 'daily'],
    displayValues: ['Monthly view', 'Daily view']
  };

  const dateSelect = {
    values: [
      'thisMonth',
      'lastMonth',
      'lastThreeMonths',
      'lastSixMonths',
      'lastTwelveMonths'
    ],
    displayValues: [
      'This month',
      'Last month',
      'Last 3 months',
      'Last 6 months',
      'Last 12 months'
    ]
  };

  const options: ChartOptions<'bar'> = {
    maintainAspectRatio: false,
    interaction: {
      intersect: false,
      mode: 'index'
    },
    plugins: {
      legend: {
        position: 'bottom',
        labels: {
          color: '#635972',
          font: {
            family: 'Noto Sans'
          },
          usePointStyle: true,
          padding: 24
        }
      },
      tooltip: {
        position: 'nearest',
        xAlign: 'center',
        yAlign: 'bottom',
        usePointStyle: true,
        backgroundColor: 'rgba(0,0,0,.75)',
        boxPadding: 8,
        padding: 16,
        titleMarginBottom: 16,
        bodyFont: {
          family: 'Noto Sans'
        },
        bodyColor: '#E9E4EC',
        titleFont: {
          family: 'Noto Sans',
          size: 14
        },
        callbacks: {
          title(chart) {
            return `${chart[0].label}. Total: $${formatNumber(
              chart
                .map(item => item.raw as number)
                .reduce((item, a) => item + a, 0)
            )}`;
          },
          label(chart) {
            return `${chart.dataset.label}: ${chart.formattedValue}`;
          }
        }
      }
    },
    scales: {
      x: {
        display: true,
        grid: {
          lineWidth: 1.5,
          color: '#F6F2FB'
        },
        ticks: {
          color: '#635972',
          font: {
            family: 'Noto Sans'
          }
        }
      },
      y: {
        display: true,
        grid: {
          lineWidth: 1.5,
          color: '#F6F2FB'
        },
        ticks: {
          color: '#635972',
          font: {
            family: 'Noto Sans'
          },
          callback(value) {
            return `$${formatNumber(Number(value))}`;
          }
        }
      }
    }
  };

  function getDatasets(dataArray: CostExplorerDatapoints[]) {
    const dates = dataArray.map(item =>
      dateFormatter(item.date, queryGranularity)
    );

    const datapoints = dataArray.map(
      item => item.datapoints
    ) as CostExplorerDatapointsProps[][];

    const tempDataSets = new Map<string, any>();
    for (let i = 0; i < datapoints.length; i += 1) {
      for (let j = 0; j < datapoints[i].length; j += 1) {
        if (!tempDataSets.has(datapoints[i][j].name)) {
          tempDataSets.set(datapoints[i][j].name, []);
        }
      }
    }

    for (let i = 0; i < dataArray.length; i += 1) {
      const newDatapoints = dataArray[i].datapoints;
      const tempLabels: string[] = [];
      newDatapoints.forEach(datapoint => {
        const { name } = datapoint;
        tempLabels.push(name);
        const listOfAmounts = tempDataSets.get(name);
        listOfAmounts.push(Number(datapoint.amount.toFixed(0)));
        tempDataSets.set(name, listOfAmounts);
      });

      Array.from(tempDataSets.keys()).forEach(name => {
        let found = false;
        tempLabels.forEach(l => {
          if (name === l) {
            found = true;
          }
        });
        if (!found) {
          const listOfAmounts = tempDataSets.get(name);
          listOfAmounts.push(null);
          tempDataSets.set(name, listOfAmounts);
        }
      });
    }

    const datasets: {
      label: string;
      data: number[];
      backgroundColor: string;
      borderColor: string;
      borderRadius: number;
      barThickness: number;
    }[] = [];

    let index = 0;
    tempDataSets.forEach((key, value) => {
      datasets.push({
        label: value,
        data: key,
        backgroundColor: colors[index],
        borderColor: 'transparent',
        borderRadius: 3,
        barThickness: queryGranularity === 'monthly' ? 20 : 8
      });
      index += 1;
    });

    const newData: ChartData<'bar'> = {
      labels: dates,
      datasets
    };

    setChartData(newData);
  }

  function handleFilterChange(type: string, newValue: string) {
    if (type === 'group') {
      setQueryGroup(newValue as CostExplorerQueryGroupProps);
    }

    if (type === 'granularity') {
      setQueryGranularity(newValue as CostExplorerQueryGranularityProps);
    }

    if (type === 'date') {
      setQueryDate(newValue as CostExplorerQueryDateProps);
    }
  }

  useEffect(() => {
    if (data && data.length > 0) {
      getDatasets(data);
    }
  }, [data]);

  return {
    chartData,
    options,
    groupBySelect,
    granularitySelect,
    dateSelect,
    handleFilterChange
  };
}

export default useCostExplorerChart;
