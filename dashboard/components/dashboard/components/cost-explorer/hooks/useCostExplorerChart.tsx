import { ChartData, ChartOptions } from 'chart.js';
import {
  ChangeEvent,
  Dispatch,
  SetStateAction,
  useEffect,
  useState
} from 'react';
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

  const colors = ['#80AAF2', '#F19B6E', '#FBC864', '#9BD6CC', '#B8D987'];

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
    values: ['provider', 'service', 'region', 'account'],
    displayValues: [
      'Cloud provider',
      'Cloud service',
      'Cloud region',
      'Cloud account'
    ]
  };

  const granularitySelect: GroupBySelectProps = {
    values: ['monthly', 'daily'],
    displayValues: ['Monthly view', 'Daily view']
  };

  const dateSelect = {
    values: [
      'lastMonth',
      'lastThreeMonths',
      'lastSixMonths',
      'lastTwelveMonths'
    ],
    displayValues: [
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
            return `${chart.dataset.label}: $${formatNumber(
              Number(chart.formattedValue)
            )}`;
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

  function handleFilterChange(e: ChangeEvent<HTMLSelectElement>, type: string) {
    if (type === 'group') {
      setQueryGroup(e.currentTarget.value as CostExplorerQueryGroupProps);
    }

    if (type === 'granularity') {
      setQueryGranularity(
        e.currentTarget.value as CostExplorerQueryGranularityProps
      );
    }

    if (type === 'date') {
      setQueryDate(e.currentTarget.value as CostExplorerQueryDateProps);
    }
  }

  useEffect(() => {
    if (data) {
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
